package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection interface {
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)

	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult

	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)

	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)

	BulkWrite(ctx context.Context, models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
}
