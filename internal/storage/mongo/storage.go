package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mmadfox/gpsgend/internal/device"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	db               *mongo.Database
	deviceCollection *mongo.Collection
	collectionName   string
	databaseName     string
}

func NewStorage(cli *mongo.Client, databaseName string, collectionName string) *Storage {
	db := cli.Database(databaseName)
	return &Storage{
		db:               db,
		deviceCollection: db.Collection(collectionName),
		collectionName:   collectionName,
		databaseName:     databaseName,
	}
}

func (s *Storage) CollectionName() string {
	return s.collectionName
}

func (s *Storage) DatabaseName() string {
	return s.databaseName
}

func (s *Storage) FindByID(ctx context.Context, deviceID uuid.UUID) (*device.Device, error) {
	filter := bson.M{"device_id": deviceID.String()}
	result := s.deviceCollection.FindOne(ctx, filter)
	deviceModel := new(deviceModel)
	if err := result.Decode(deviceModel); err != nil {
		return nil, err
	}
	return decodeDevice(deviceModel)
}

func (s *Storage) Insert(ctx context.Context, d *device.Device) error {
	deviceModel, err := encodeDevice(d)
	if err != nil {
		return err
	}
	if _, err := s.deviceCollection.InsertOne(ctx, deviceModel); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Update(ctx context.Context, d *device.Device) error {
	deviceModel, err := encodeDevice(d)
	if err != nil {
		return err
	}
	filter := bson.M{
		"device_id": deviceModel.ID,
		"version":   deviceModel.Version,
	}
	deviceModel.Version += 1
	deviceModel.ID = ""
	update := bson.M{
		"$set": deviceModel,
	}
	res, err := s.deviceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		return fmt.Errorf("failed to update device, mismatched document version")
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, deviceID uuid.UUID) error {
	filter := bson.M{
		"device_id": deviceID.String(),
	}
	_, err := s.deviceCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) WalkByStatus(ctx context.Context, ds []device.Status, fn func(*device.Device) error) error {
	filter := bson.M{
		"status.id": bson.M{
			"$in": ds,
		},
	}

	cursor, err := s.deviceCollection.Find(ctx, filter)
	if err != nil {
		return err
	}

	for {
		if cursor.TryNext(ctx) {
			var result deviceModel
			if err := cursor.Decode(&result); err != nil {
				return err
			}
			dev, err := decodeDevice(&result)
			if err != nil {
				return err
			}
			if err := fn(dev); err != nil {
				return err
			}
			continue
		}

		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}

		if cursor.ID() == 0 {
			break
		}
	}
	return nil
}

func (s *Storage) Search(ctx context.Context, f device.QueryFilter) (result device.SearchResult, err error) {
	filter := encodeQueryFilter(&f)
	if f.Limit <= 0 {
		f.Limit = 50
	}
	if f.Limit > 1000 {
		f.Limit = 1000
	}

	if f.Page < 0 {
		f.Page = 0
	}

	opts := options.Find().SetLimit(f.Limit)
	var page int64
	if f.Page > 0 {
		page = f.Page * f.Limit
		opts.SetSkip(page)
	}

	count, err := s.deviceCollection.CountDocuments(ctx, filter)
	if err != nil {
		return result, err
	}
	result.Meta.TotalDevices = count

	cursor, err := s.deviceCollection.Find(ctx, filter, opts)
	if err != nil {
		return result, err
	}

	result.Devices = make([]device.DeviceView, 0, f.Limit)
	result.Meta.Page = f.Page
	result.Meta.Limit = f.Limit

	for {
		if cursor.TryNext(ctx) {
			var deviceView device.DeviceView
			if err := cursor.Decode(&deviceView); err != nil {
				return result, err
			}
			result.Devices = append(result.Devices, deviceView)
			continue
		}

		if err := cursor.Err(); err != nil {
			return result, err
		}

		if cursor.ID() == 0 {
			break
		}
	}

	result.Meta.Found = int64(len(result.Devices))

	return result, nil
}

func (s *Storage) EnsureIndexes() error {
	indexes := []mongo.IndexModel{}
	ok := true

	indexes = append(indexes, mongo.IndexModel{
		Keys:    bson.M{"device_id": 1},
		Options: &options.IndexOptions{Unique: &ok},
	})

	indexes = append(indexes, mongo.IndexModel{
		Keys: bson.M{"status.id": 1},
	})

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := s.deviceCollection.Indexes().CreateMany(
		context.Background(),
		indexes,
		opts,
	)
	if err != nil {
		return err
	}
	return nil
}
