package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnsureIndexes(ctx context.Context, c *mongo.Collection) error {
	indexes := []mongo.IndexModel{}
	unique := true
	indexes = append(indexes, mongo.IndexModel{
		Keys:    bson.M{"tracker_id": 1},
		Options: &options.IndexOptions{Unique: &unique},
	})
	indexes = append(indexes, mongo.IndexModel{
		Keys: bson.M{"status.id": 1},
	})
	// TODO:
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	if _, err := c.Indexes().CreateMany(
		ctx,
		indexes,
		opts,
	); err != nil {
		return err
	}
	return nil
}
