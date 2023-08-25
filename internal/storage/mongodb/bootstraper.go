package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Bootstraper struct {
	collection Collection
}

func NewBootstraper(c Collection) *Bootstraper {
	return &Bootstraper{collection: c}
}

func (b *Bootstraper) LoadTrackers(ctx context.Context, g generator.Processes) error {
	docs := make([]mongo.WriteModel, 0, 16)
	updatedAt := time.Now().Unix()

	filter := bson.D{{Key: "status.id", Value: types.Paused}}
	opts := options.Find().SetProjection(bson.D{
		{Key: "tracker_id", Value: 1},
		{Key: "snapshot", Value: 1},
	})
	cursor, err := b.collection.Find(ctx, filter, opts)
	if err != nil {
		return fmt.Errorf("%w: mongo.bootstrapper.LoadTrackers()", err)
	}
	defer cursor.Close(ctx)

	for cursor.TryNext(ctx) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return err
		}
		snap, ok := doc["snapshot"].(primitive.Binary)
		if !ok || len(snap.Data) == 0 {
			continue
		}
		p, err := gpsgen.DecodeTracker(snap.Data)
		if err != nil {
			return fmt.Errorf("%w: gpsgen.DecodeTracker()", err)
		}

		g.Attach(p)

		update := mongo.NewUpdateOneModel().
			SetFilter(bson.D{
				{Key: fieldTrackerID, Value: p.ID()},
			}).
			SetUpdate(bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status.id", Value: types.Running},
					{Key: "status.name", Value: types.Running.String()},
					{Key: "updated_at", Value: updatedAt},
					{Key: "snapshot", Value: nil},
				}},
			}).
			SetUpsert(false)
		docs = append(docs, update)
	}
	if cursor.Err() != nil {
		return fmt.Errorf("%w mongo.Cursor()", cursor.Err())
	}

	if len(docs) > 0 {
		_, err = b.collection.BulkWrite(ctx, docs)
		if err != nil {
			return fmt.Errorf("%w: mongodb.bootstraper.LoadTrackers()", err)
		}
	}

	return nil
}

func (b *Bootstraper) UnloadTrackers(ctx context.Context, g generator.Processes) error {
	docs := make([]mongo.WriteModel, 0, 16)
	updatedAt := time.Now().Unix()

	var unloadErr error

	g.Each(func(n int, p *gpsgen.Device) bool {
		snapshot, err := p.MarshalBinary()
		if err != nil {
			unloadErr = err
			return false
		}

		doc := mongo.NewUpdateOneModel().
			SetFilter(bson.D{
				{Key: fieldTrackerID, Value: p.ID()},
				{Key: "status.id", Value: types.Running},
			}).
			SetUpdate(bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status.id", Value: types.Paused},
					{Key: "status.name", Value: types.Paused.String()},
					{Key: "updated_at", Value: updatedAt},
					{Key: "snapshot", Value: snapshot},
				}},
			}).
			SetUpsert(false)
		docs = append(docs, doc)
		return true
	})
	if unloadErr != nil {
		return fmt.Errorf("%w: mongodb.bootstraper.UnloadTrackers()", unloadErr)
	}

	if len(docs) > 0 {
		opts := options.BulkWrite()
		_, err := b.collection.BulkWrite(ctx, docs, opts)
		if err != nil {
			return fmt.Errorf("%w: mongodb.bootstraper.UnloadTrackers()", err)
		}
	}

	return nil
}
