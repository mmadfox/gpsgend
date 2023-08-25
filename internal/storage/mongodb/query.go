package mongodb

import (
	"context"

	"github.com/mmadfox/gpsgend/internal/generator"
)

type Query struct {
	collection Collection
}

func NewQuery(c Collection) *Query {
	return &Query{collection: c}
}

func (q *Query) SearchTrackers(ctx context.Context, f generator.Filter) (generator.SearchResult, error) {
	return generator.SearchResult{}, nil
}
