package mongodb

import (
	"context"
	"fmt"

	"github.com/mmadfox/gpsgend/internal/generator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Query struct {
	collection Collection
}

func NewQuery(c Collection) *Query {
	return &Query{collection: c}
}

func (q *Query) SearchTrackers(ctx context.Context, f generator.Filter) (generator.SearchResult, error) {
	findOpts := options.Find().SetSort(bson.M{"created_at": -1})
	filter := bson.D{}
	if f.Limit <= 1 {
		f.Limit = 1
	}
	if f.Limit > 1000 {
		f.Limit = 1000
	}

	f.Limit = f.Limit + 1
	findOpts.SetLimit(f.Limit)
	if f.Offset > 0 {
		filter = append(filter, bson.E{Key: "created_at", Value: bson.M{"$lte": f.Offset}})
	}
	if len(f.Term) > 0 {
		filter = append(filter, bson.E{Key: "$text", Value: bson.M{"$search": f.Term}})
	}
	if f.Status > 0 {
		filter = append(filter, bson.E{Key: "status.id", Value: f.Status})
	}
	if len(f.TrackerIDs) > 0 {
		filter = append(filter, bson.E{Key: "tracker_id", Value: bson.M{"$in": f.TrackerIDs}})
	}

	cursor, err := q.collection.Find(ctx, filter, findOpts)
	if err != nil {
		return generator.SearchResult{}, fmt.Errorf("%w: %v",
			generator.ErrStorageSearch, err)
	}
	defer cursor.Close(ctx)

	results := make([]generator.TrackerView, 0, f.Limit)

	for cursor.TryNext(ctx) {
		var model generator.TrackerView
		if err = cursor.Decode(&model); err != nil {
			return generator.SearchResult{}, err
		}
		results = append(results, model)
	}

	var next int64
	if len(results) > int(f.Limit)-1 {
		last := results[int(f.Limit)-1]
		results = results[:int(f.Limit)-1]
		next = last.CreatedAt
	}

	return generator.SearchResult{
		Trackers: results,
		Next:     next,
	}, nil
}
