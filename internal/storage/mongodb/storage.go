package mongodb

import (
	"context"
	"fmt"

	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	fieldTrackerID = "tracker_id"
)

type Storage struct {
	collection Collection
}

func New(c Collection) *Storage {
	return &Storage{collection: c}
}

func (s *Storage) Insert(ctx context.Context, t *generator.Tracker) error {
	if t == nil {
		return fmt.Errorf("%w: storage.Insert() *Tracker is nil",
			generator.ErrNoTracker)
	}

	doc := new(generator.TrackerSnapshot)
	t.TakeSnapshot(doc)

	if len(doc.ID) == 0 {
		return fmt.Errorf("%w: storage.Insert() *Tracker is invalid",
			generator.ErrNoTracker)
	}

	if _, err := s.collection.InsertOne(ctx, doc); err != nil {
		return fmt.Errorf("%w: %v", generator.ErrStorageInsert, err)
	}

	return nil
}

func (s *Storage) Find(ctx context.Context, trackerID types.ID) (*generator.Tracker, error) {
	if trackerID.IsEmpty() {
		return nil, fmt.Errorf("%w: storage.Find() trackerID is empty",
			types.ErrInvalidID)
	}

	filter := bson.M{fieldTrackerID: trackerID.String()}
	result := s.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, fmt.Errorf("%w: %v",
			generator.ErrStorageFind, result.Err())
	}

	doc := new(generator.TrackerSnapshot)
	if err := result.Decode(doc); err != nil {
		return nil, fmt.Errorf("%w: decode from mongodb: %v",
			generator.ErrBrokenTracker, err)
	}

	trk := new(generator.Tracker)
	if err := trk.FromSnapshot(doc); err != nil {
		return nil, fmt.Errorf("%w: decode from snapshot: %v",
			generator.ErrBrokenTracker, err)
	}

	return trk, nil
}

func (s *Storage) Delete(ctx context.Context, trackerID types.ID) error {
	if trackerID.IsEmpty() {
		return fmt.Errorf("%w: storage.Delete() trackerID is empty",
			types.ErrInvalidID)
	}

	filter := bson.M{fieldTrackerID: trackerID.String()}
	if _, err := s.collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("%w: %v", generator.ErrStorageDelete, err)
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, t *generator.Tracker) error {
	if t == nil {
		return fmt.Errorf("%w: storage.Update() *Tracker is nil",
			generator.ErrNoTracker)
	}

	doc := new(generator.TrackerSnapshot)
	t.TakeSnapshot(doc)

	filter := bson.M{
		fieldTrackerID: doc.ID,
		"version":      doc.Version,
	}
	doc.Version += 1
	doc.ID = "" // reset for update

	update := bson.M{"$set": doc}
	res, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("%w: %v", generator.ErrStorageUpdate, err)
	}
	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		return fmt.Errorf("%w: storage.Update(), expected %d",
			generator.ErrInvalidTrackerVersion, doc.Version)
	}

	return nil
}

var _ generator.Storage = &Storage{}
