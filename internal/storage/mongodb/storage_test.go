package mongodb_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mmadfox/gpsgend/internal/generator"
	storagemongodb "github.com/mmadfox/gpsgend/internal/storage/mongodb"
	"github.com/mmadfox/gpsgend/internal/types"
	mockmongodb "github.com/mmadfox/gpsgend/tests/mocks/storage/mongodb"
	stubgenerator "github.com/mmadfox/gpsgend/tests/stub/generator"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestStorage_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	tracker := stubgenerator.Tracker()

	type fields struct {
		collection func() *mockmongodb.MockCollection
	}
	type args struct {
		t *generator.Tracker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "should return error when tracker is nil",
			args: args{
				t: nil,
			},
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					return mock
				},
			},
			wantErr: generator.ErrNoTracker,
		},
		{
			name: "should return error when tracker is invalid",
			args: args{
				t: new(generator.Tracker),
			},
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					return mock
				},
			},
			wantErr: generator.ErrNoTracker,
		},
		{
			name: "should return error when collection.InsertOne failure",
			args: args{
				t: tracker,
			},
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					mock.EXPECT().InsertOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("someError"))
					return mock
				},
			},
			wantErr: generator.ErrStorageInsert,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storagemongodb.New(tt.fields.collection())
			if err := s.Insert(context.Background(), tt.args.t); !errors.Is(err, tt.wantErr) {
				t.Errorf("Storage.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	tracker := stubgenerator.Tracker()

	type fields struct {
		collection func() *mockmongodb.MockCollection
	}
	type args struct {
		ID types.ID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackeID is empty",
			args: args{
				ID: types.ID{},
			},
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					return mock
				},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when collection.FindOne failure",
			args: args{
				ID: tracker.ID(),
			},
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					res := mongo.NewSingleResultFromDocument(nil, errors.New("some error"), nil)
					mock.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(res)
					return mock
				},
			},
			wantErr: generator.ErrStorageFind,
		},
		{
			name: "should return error when restore invalid tracker ",
			args: args{
				ID: tracker.ID(),
			},
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					doc := new(generator.TrackerSnapshot)
					data, err := bson.Marshal(doc)
					require.NoError(t, err)
					res := mongo.NewSingleResultFromDocument(data, nil, nil)
					mock.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(res)
					return mock
				},
			},
			wantErr: generator.ErrBrokenTracker,
		},
		{
			name: "should not return error when all params are valid",
			args: args{
				ID: tracker.ID(),
			},
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					snap := new(generator.TrackerSnapshot)
					tracker.TakeSnapshot(snap)
					res := mongo.NewSingleResultFromDocument(snap, nil, nil)
					mock.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(res)
					return mock
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storagemongodb.New(tt.fields.collection())
			got, err := s.Find(context.Background(), tt.args.ID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Storage.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != nil {
				return
			}
			require.NotNil(t, got)
		})
	}
}

func TestStorage_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		trackerID types.ID
	}
	type fields struct {
		collection func() *mockmongodb.MockCollection
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackerID is emtpy",
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					return mock
				},
			},
			args: args{
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when collection.DeleteOne failure",
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					mock.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil, errors.New("someError"))
					return mock
				},
			},
			args: args{
				trackerID: types.NewID(),
			},
			wantErr: generator.ErrStorageDelete,
		},
		{
			name: "should not return error when all params are valid",
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					mock.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil, nil)
					return mock
				},
			},
			args: args{
				trackerID: types.NewID(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storagemongodb.New(tt.fields.collection())
			if err := s.Delete(context.Background(), tt.args.trackerID); !errors.Is(err, tt.wantErr) {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	tracker := stubgenerator.Tracker()

	type fields struct {
		collection func() *mockmongodb.MockCollection
	}
	type args struct {
		t *generator.Tracker
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "should return error when tracker is nil",
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					return mock
				},
			},
			wantErr: generator.ErrNoTracker,
		},
		{
			name: "should return error when storage.Update failure",
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					mock.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("someError"))
					return mock
				},
			},
			args: args{
				t: tracker,
			},
			wantErr: generator.ErrStorageUpdate,
		},
		{
			name: "should return error when mismatches tracker version",
			fields: fields{
				collection: func() *mockmongodb.MockCollection {
					mock := mockmongodb.NewMockCollection(ctrl)
					res := mongo.UpdateResult{
						MatchedCount:  0,
						ModifiedCount: 0,
					}
					mock.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(&res, nil)
					return mock
				},
			},
			args: args{
				t: tracker,
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storagemongodb.New(tt.fields.collection())
			if err := s.Update(context.Background(), tt.args.t); !errors.Is(err, tt.wantErr) {
				t.Errorf("Storage.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
