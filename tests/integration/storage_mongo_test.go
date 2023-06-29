package integration

import (
	"context"
	"time"

	"github.com/mmadfox/gpsgend/internal/device"
	storagemongo "github.com/mmadfox/gpsgend/internal/storage/mongo"
	stubdevice "github.com/mmadfox/gpsgend/tests/stubs/device"
	"github.com/mmadfox/testcontainers/infra"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type storageMongoSuite struct {
	suite.Suite
	infra         *infra.Sets
	collection    *mongo.Collection
	deviceStorage *storagemongo.Storage
	devices       []*device.Device
	ctx           context.Context
}

func (s *storageMongoSuite) SetupSuite() {
	s.deviceStorage = storagemongo.NewStorage(
		s.infra.MongoDB().Client(),
		"test",
		"devices",
	)
	s.collection = s.infra.MongoDB().
		Client().
		Database(s.deviceStorage.DatabaseName()).
		Collection(s.deviceStorage.CollectionName())
}

func (s *storageMongoSuite) SetupTest() {
	s.ctx = context.Background()
	s.devices = make([]*device.Device, 0)
	for i := 0; i < 10; i++ {
		newDevice := stubdevice.Device()
		err := s.deviceStorage.Insert(s.ctx, newDevice)
		require.NoError(s.T(), err)
		s.devices = append(s.devices, newDevice)
	}
}

func (s *storageMongoSuite) TearDownTest() {
	s.collection.DeleteMany(s.ctx, bson.M{})
}

func (s *storageMongoSuite) TestStorageMongoEnsureIndex() {
	wantIndexes := 3
	for i := 0; i < 5; i++ {
		err := s.deviceStorage.EnsureIndexes()
		require.NoError(s.T(), err)
	}
	indexView := s.collection.Indexes()
	opts := options.ListIndexes().SetMaxTime(2 * time.Second)
	cursor, err := indexView.List(context.TODO(), opts)
	require.NoError(s.T(), err)
	var result []bson.M
	err = cursor.All(s.ctx, &result)
	require.NoError(s.T(), err)
	require.Len(s.T(), result, wantIndexes)
}

func (s *storageMongoSuite) TestStorageMongoFindByID() {
	total := 0
	for i := 0; i < len(s.devices); i++ {
		want := s.devices[i]
		got, err := s.deviceStorage.FindByID(s.ctx, want.ID())
		require.NoError(s.T(), err)
		require.NotNil(s.T(), got)
		require.Equal(s.T(), want.ID(), got.ID())
		total++
	}
	require.Equal(s.T(), len(s.devices), total)
}

func (s *storageMongoSuite) TestStorageMongoUpdate() {
	for i := 0; i < len(s.devices); i++ {
		want := s.devices[i]
		ver := device.Version(want)
		err := s.deviceStorage.Update(s.ctx, want)
		require.NoError(s.T(), err)

		got, err := s.deviceStorage.FindByID(s.ctx, want.ID())
		require.NoError(s.T(), err)
		require.NotNil(s.T(), got)
		require.Equal(s.T(), ver+1, device.Version(got))
	}
}

func (s *storageMongoSuite) TestStorageMongoDelete() {
	cnt, err := s.collection.CountDocuments(s.ctx, bson.M{})
	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(len(s.devices)), cnt)
	for i := 0; i < len(s.devices); i++ {
		err := s.deviceStorage.Delete(s.ctx, s.devices[i].ID())
		require.NoError(s.T(), err)
	}
	cnt, err = s.collection.CountDocuments(s.ctx, bson.M{})
	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(0), cnt)
}

func (s *storageMongoSuite) TestStorageMongoWalkByStatus() {
	total := 0
	status := []device.Status{device.Stopped}
	err := s.deviceStorage.WalkByStatus(s.ctx, status, func(dev *device.Device) error {
		total++
		require.NotNil(s.T(), dev)
		return nil
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), len(s.devices), total)
}

func (s *storageMongoSuite) TestStorageMongoSearchWithEmptyFilter() {
	result, err := s.deviceStorage.Search(s.ctx, device.QueryFilter{})
	require.NoError(s.T(), err)
	require.Equal(s.T(), len(s.devices), int(result.Meta.TotalDevices))
	require.Equal(s.T(), len(s.devices), len(result.Devices))
}

func (s *storageMongoSuite) TestStorageMongoSearchWithPaging() {
	total := len(s.devices)
	visits := make(map[string]int, total)
	for i := 0; i < total; i++ {
		numPage := i
		result, err := s.deviceStorage.Search(s.ctx, device.QueryFilter{
			Page:  int64(numPage),
			Limit: 1,
		})
		require.NoError(s.T(), err)
		require.Len(s.T(), result.Devices, 1)
		visits[result.Devices[0].ID]++
	}
	require.Len(s.T(), visits, total)
}
