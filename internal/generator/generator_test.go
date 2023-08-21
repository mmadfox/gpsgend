package generator_test

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/go-gpsgen/properties"
	"github.com/mmadfox/go-gpsgen/proto"
	stdtypes "github.com/mmadfox/go-gpsgen/types"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
	mockgenerator "github.com/mmadfox/gpsgend/tests/mocks/generator"
	"github.com/stretchr/testify/require"
)

type mocks struct {
	storage     func() *mockgenerator.MockStorage
	processes   func() *mockgenerator.MockProcesses
	bootstraper func() *mockgenerator.MockBootstraper
}

func offlineType(t *testing.T, min, max int) types.Offline {
	typ, err := types.ParseOffline(min, max)
	require.NoError(t, err, "tracker.offline")
	return typ
}

func elevationType(t *testing.T, min, max float64, amplitude int, mode stdtypes.SensorMode) types.Elevation {
	typ, err := types.ParseElevation(min, max, amplitude, mode)
	require.NoError(t, err, "tracker.elevation")
	return typ
}

func speedType(t *testing.T, min, max float64, amplitude int) types.Speed {
	typ, err := types.ParseSpeed(min, max, amplitude)
	require.NoError(t, err, "tracker.speed")
	return typ
}

func batteryType(t *testing.T, min, max float64, chargeTime time.Duration) types.Battery {
	typ, err := types.ParseBattery(min, max, chargeTime)
	require.NoError(t, err, "tracker.battery")
	return typ
}

func descrType(t *testing.T, val string) types.Description {
	typ, err := types.ParseDescription(val)
	require.NoError(t, err, "tracker.description")
	return typ
}

func customIDType(t *testing.T, val string) types.CustomID {
	typ, err := types.ParseCustomID(val)
	require.NoError(t, err, "tracker.userID")
	return typ
}

func colorType(t *testing.T, val string) types.Color {
	typ, err := types.ParseColor(val)
	require.NoError(t, err, "tracker.color")
	return typ
}

func modelType(t *testing.T, val string) types.Model {
	typ, err := types.ParseModel(val)
	require.NoError(t, err, "tracker.model")
	return typ
}

func modelPtrType(t *testing.T, val string) *types.Model {
	model := modelType(t, val)
	return &model
}

func colorPtrType(t *testing.T, val string) *types.Color {
	color := colorType(t, val)
	return &color
}

func customIDPtrType(t *testing.T, val string) *types.CustomID {
	cid := customIDType(t, val)
	return &cid
}

func descrPtrType(t *testing.T, val string) *types.Description {
	descr := descrType(t, val)
	return &descr
}

func s2id(id string) types.ID {
	uid, err := types.ParseID(id)
	if err != nil {
		panic(err)
	}
	return uid
}

func newGeneratorFromMocks(m mocks) *generator.Generator {
	var (
		s *mockgenerator.MockStorage
		p *mockgenerator.MockProcesses
		b *mockgenerator.MockBootstraper
	)
	if m.storage != nil {
		s = m.storage()
	}

	if m.processes != nil {
		p = m.processes()
	}

	if m.bootstraper != nil {
		b = m.bootstraper()
	}

	return generator.New(s, p, b)
}

func TestGenerator_NewTracker(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx  context.Context
		opts generator.NewTrackerOptions
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		assert  func(*generator.Tracker)
		wantErr bool
	}{
		{
			name: "should return error when options is empty",
			args: args{
				ctx:  context.Background(),
				opts: generator.NewTrackerOptions{},
			},
			fields: mocks{
				storage:   func() *mockgenerator.MockStorage { return nil },
				processes: func() *mockgenerator.MockProcesses { return nil },
			},
			wantErr: true,
		},
		{
			name: "should return error when storage.Insert failure",
			args: args{
				ctx: context.Background(),
				opts: generator.NewTrackerOptions{
					Offline:   offlineType(t, 1, 10),
					Battery:   batteryType(t, 1.0, 100.0, time.Hour),
					Speed:     speedType(t, 1.0, 3.0, 8),
					Elevation: elevationType(t, 0.0, 100.0, 8, 0),
				},
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().Insert(gomock.Any(), gomock.Any()).
						Return(errors.New("err")).
						Times(1)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses { return nil },
			},
			wantErr: true,
		},
		{
			name: "should return new tracker when all params are valid",
			args: args{
				ctx: context.Background(),
				opts: generator.NewTrackerOptions{
					Model:       modelPtrType(t, "Tracker-N2x91"),
					Color:       colorPtrType(t, "#ff0000"),
					UserID:      customIDPtrType(t, uuid.NewString()),
					Descr:       descrPtrType(t, "some descr"),
					Offline:     offlineType(t, 1, 10),
					Battery:     batteryType(t, 1.0, 100.0, time.Hour),
					Speed:       speedType(t, 1.0, 3.0, 8),
					Elevation:   elevationType(t, 0.0, 100.0, 8, 0),
					Props:       &properties.Properties{"foo": 1},
					SkipOffline: true,
				},
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().Insert(gomock.Any(), gomock.Any()).
						DoAndReturn(func(ctx context.Context, t *generator.Tracker) (*generator.Tracker, error) { return t, nil })
					return mock
				},
				processes: func() *mockgenerator.MockProcesses { return nil },
			},
			assert: func(trk *generator.Tracker) {
				require.NotEmpty(t, trk.ID().String())
				require.NotEmpty(t, trk.Model().String())
				require.NotEmpty(t, trk.UserID().String())
				require.NotEmpty(t, trk.Color().String())
				require.NotEmpty(t, trk.Description().String())
				require.Equal(t, 1.0, trk.Battery().Min())
				require.Equal(t, 100.0, trk.Battery().Max())
				require.Equal(t, time.Hour, trk.Battery().ChargeTime())
				require.Equal(t, 1.0, trk.Speed().Min())
				require.Equal(t, 3.0, trk.Speed().Max())
				require.Equal(t, 8, trk.Speed().Amplitude())
				require.Equal(t, 0.0, trk.Elevation().Min())
				require.Equal(t, 100.0, trk.Elevation().Max())
				require.Equal(t, 8, trk.Elevation().Amplitude())
				require.Equal(t, 1, trk.Offline().Min())
				require.Equal(t, 10, trk.Offline().Max())
				require.NotEmpty(t, trk.Properties())
				require.True(t, trk.SkipOffline())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.NewTracker(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.NewTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if tt.assert != nil {
				tt.assert(got)
			}
		})
	}
}

func TestGenerator_RemoveTracker(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx     context.Context
		trackID types.ID
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr bool
	}{
		{
			name: "should return error when storage.Delete failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("fail"))
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Detach(gomock.Any()).Times(1)
					return mock
				},
			},
			wantErr: true,
		},
		{
			name: "should remove track when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Detach(gomock.Any()).Times(1)
					return mock
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.RemoveTracker(tt.args.ctx, tt.args.trackID); (err != nil) != tt.wantErr {
				t.Errorf("Generator.RemoveTracker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_UpdateTracker(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx       context.Context
		trackerID types.ID
		opts      generator.UpdateTrackerOptions
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr bool
	}{
		{
			name: "should return error when options are empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				opts:      generator.UpdateTrackerOptions{},
			},
			wantErr: true,
		},
		{
			name: "should return error when options invalid",
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				opts: generator.UpdateTrackerOptions{
					Model: &types.Model{},
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when tracker not found",
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				opts: generator.UpdateTrackerOptions{
					Model: modelPtrType(t, "model"),
					Color: colorPtrType(t, "#ff0000"),
				},
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when tracker is paused",
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				opts: generator.UpdateTrackerOptions{
					Model: modelPtrType(t, "model"),
					Color: colorPtrType(t, "#ff0000"),
				},
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					generator.Debug_InjectInvalidDatatForTracker(trk, "status.paused")
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when storage.Update failure",
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				opts: generator.UpdateTrackerOptions{
					Model: modelPtrType(t, "model"),
					Color: colorPtrType(t, "#ff0000"),
				},
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("error"))
					return mock
				},
			},
			wantErr: true,
		},
		{
			name: "should not return error when all params are valid",
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				opts: generator.UpdateTrackerOptions{
					Model:  modelPtrType(t, "model"),
					Color:  colorPtrType(t, "#ff0000"),
					Descr:  descrPtrType(t, "descr"),
					UserID: customIDPtrType(t, types.NewID().String()),
				},
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					tracker := gpsgen.NewAnimalTracker()
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(gomock.Any()).Times(1).Return(tracker, true)
					return mock
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.UpdateTracker(tt.args.ctx, tt.args.trackerID, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("Generator.UpdateTracker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_FindTracker(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr bool
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				trackerID: types.ID{},
			},
			wantErr: true,
		},
		{
			name: "should return tracker when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					trk := new(generator.Tracker)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Return(trk, nil)
					return mock
				},
			},
			args: args{
				trackerID: types.NewID(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			_, err := g.FindTracker(tt.args.ctx, tt.args.trackerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.FindTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGenerator_StartTracker(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx       context.Context
		trackerID types.ID
	}
	tests := []struct {
		name   string
		fields mocks
		args   args
		want   error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			want: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker is running",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerIsAlreadyRunning,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is paused",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerOff,
		},
		{
			name: "should return error when tracker is running",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerIsAlreadyRunning,
		},
		{
			name: "should return error when tracker without routes",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerHasNoRoutes,
		},
		{
			name: "should return error when tracker params are invalid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(gpsgen.RandomRouteForMoscow())
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: stdtypes.ErrMinAmplitude,
		},
		{
			name: "should return error when tracker storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Stopped)
					trk.AddRoute(gpsgen.RandomRouteForMoscow())
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should return error when tracker storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Stopped)
					trk.AddRoute(gpsgen.RandomRouteForMoscow())
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().HasTracker(gomock.Any()).Times(1).Return(false)
					mock.EXPECT().Attach(gomock.Any()).Times(1)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.StartTracker(tt.args.ctx, tt.args.trackerID); !errors.Is(err, tt.want) {
				t.Errorf("Generator.StartTracker() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func makeValidTracker(t *testing.T, status types.DeviceStatus) *generator.Tracker {
	builder := generator.NewTrackerBuilder()
	builder.Speed(speedType(t, 1, 5, 8))
	builder.Battery(batteryType(t, 0, 100, time.Hour))
	builder.Offline(offlineType(t, 1, 300))
	builder.Elevation(elevationType(t, 1, 300, 8, stdtypes.WithSensorStartMode))
	builder.Status(status)
	builder.Color(colorType(t, "#000000"))
	builder.Model(modelType(t, "Tracker-ty19"))
	builder.Props(properties.Properties{"foo": "foo"})
	tracker, err := builder.Build()
	require.NoError(t, err)
	tracker.AddRoute(gpsgen.RandomRouteForMoscow())
	tracker.AddRoute(gpsgen.RandomRouteForMoscow())
	tracker.AddRoute(gpsgen.RandomRouteForMoscow())
	tracker.AddRoute(gpsgen.RandomRouteForMoscow())
	s1, err := gpsgen.NewSensor("s1", 1, 10, 8, stdtypes.WithSensorEndMode)
	require.NoError(t, err)
	tracker.AddSensor(s1)
	return tracker
}

func TestGenerator_StopTracker(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx       context.Context
		trackerID types.ID
	}
	tests := []struct {
		name   string
		fields mocks
		args   args
		want   error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			want: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is paused",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerOff,
		},
		{
			name: "should return error when tracker is stopped",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Stopped)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrTrackerIsAlreadyStopped,
		},
		{
			name: "should return error when tracker storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Detach(gomock.Any()).Times(1)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: generator.ErrInvalidTrackerVersion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.StopTracker(tt.args.ctx, tt.args.trackerID); !errors.Is(err, tt.want) {
				t.Errorf("Generator.StopTracker() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestGenerator_TrackerState(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}

	tracker := gpsgen.NewAnimalTracker()
	tracker.AddRoute(gpsgen.RandomRouteForNewYork())
	tracker.DestinationTo(100)

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    *proto.Device
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not running",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(gomock.Any()).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should return device state when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(gomock.Any()).Return(tracker, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			want: tracker.State(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.TrackerState(tt.args.ctx, tt.args.trackerID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.TrackerState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.TrackerState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_AddRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx       context.Context
		trackerID types.ID
		newRoutes []*gpsgen.Route
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is paused",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
			},
			wantErr: generator.ErrTrackerOff,
		},
		{
			name: "should return error when routes list are emty",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				newRoutes: []*gpsgen.Route{},
			},
			wantErr: generator.ErrNoRoutes,
		},
		{
			name: "should return error when storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				newRoutes: []*gpsgen.Route{
					gpsgen.RandomRouteForNewYork(),
				},
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(gomock.All()).Return(gpsgen.NewTracker(), true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				newRoutes: []*gpsgen.Route{
					gpsgen.RandomRouteForNewYork(),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.AddRoutes(tt.args.ctx, tt.args.trackerID, tt.args.newRoutes); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.AddRoutes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_RemoveRoute(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx       context.Context
		trackerID types.ID
		routeID   types.ID
	}

	expectedTrackerID := types.NewID()
	expectedRoute := gpsgen.RandomRouteForNewYork()

	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when routeID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is paused",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrTrackerOff,
		},
		{
			name: "should return error when storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Times(1).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					proc := gpsgen.NewAnimalTracker()
					proc.AddRoute(expectedRoute)
					mock.EXPECT().Lookup(gomock.Any()).Times(1).Return(proc, true)
					mock.EXPECT().Detach(proc.ID()).Times(1)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.RemoveRoute(tt.args.ctx, tt.args.trackerID, tt.args.routeID); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.RemoveRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_Routes(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx       context.Context
		trackerID types.ID
		routeID   types.ID
	}

	expectedTrackerID := types.NewID()
	expectedRoute := gpsgen.RandomRouteForNewYork()

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    []*gpsgen.Route
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when routeID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			want:    []*gpsgen.Route{expectedRoute},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.Routes(tt.args.ctx, tt.args.trackerID, tt.args.routeID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.Routes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.Routes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_RouteAt(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	expectedRoute := gpsgen.RandomRouteForParis()

	type args struct {
		ctx        context.Context
		trackerID  types.ID
		routeIndex int
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    *gpsgen.Route
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:        context.Background(),
				trackerID:  expectedTrackerID,
				routeIndex: -1,
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when route not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:        context.Background(),
				trackerID:  expectedTrackerID,
				routeIndex: -1,
			},
			wantErr: generator.ErrRouteNotFound,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:        context.Background(),
				trackerID:  expectedTrackerID,
				routeIndex: 1,
			},
			wantErr: nil,
			want:    expectedRoute,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.RouteAt(tt.args.ctx, tt.args.trackerID, tt.args.routeIndex)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.RouteAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.RouteAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_RouteByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	expectedRoute := gpsgen.RandomRouteForParis()

	type args struct {
		ctx       context.Context
		trackerID types.ID
		routeID   types.ID
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    *gpsgen.Route
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when routeID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when route not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   types.NewID(),
			},
			wantErr: generator.ErrRouteNotFound,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: nil,
			want:    expectedRoute,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.RouteByID(tt.args.ctx, tt.args.trackerID, tt.args.routeID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.RouteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.RouteByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_ResetRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	expectedRoute := gpsgen.RandomRouteForParis()

	type args struct {
		ctx       context.Context
		trackerID types.ID
		routeID   types.ID
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is paused",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrTrackerOff,
		},
		{
			name: "should return error when storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddRoute(expectedRoute)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					proc := gpsgen.NewBicycleTracker()
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(proc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(expectedRoute.ID()),
			},
			wantErr: nil,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.ResetRoutes(tt.args.ctx, tt.args.trackerID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.ResetRoutes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.ResetRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_ResetNavigator(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					proc := gpsgen.NewAnimalTracker()
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(proc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.ResetNavigator(tt.args.ctx, tt.args.trackerID); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.ResetNavigator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_ToNextRoute(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	opts := gpsgen.NewDeviceOptions()
	opts.Navigator.SkipOffline = true
	expectedProc, err := gpsgen.NewDevice(opts)
	require.NoError(t, err)
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    func() types.Navigator
		want1   bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(expectedProc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: nil,
			want1:   true,
			want: func() types.Navigator {
				return types.NavigatorFromProc(expectedProc)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, got1, err := g.ToNextRoute(tt.args.ctx, tt.args.trackerID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.ToNextRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Generator.ToNextRoute() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Generator.ToNextRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerator_ToPrevRoute(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	opts := gpsgen.NewDeviceOptions()
	opts.Navigator.SkipOffline = true
	expectedProc, err := gpsgen.NewDevice(opts)
	require.NoError(t, err)
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())
	expectedProc.MoveToRoute(1)

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    func() types.Navigator
		want1   bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(expectedProc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: nil,
			want1:   true,
			want: func() types.Navigator {
				return types.NavigatorFromProc(expectedProc)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, got1, err := g.ToPrevRoute(tt.args.ctx, tt.args.trackerID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.ToPrevRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Generator.ToPrevRoute() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Generator.ToPrevRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerator_MoveToRoute(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	opts := gpsgen.NewDeviceOptions()
	opts.Navigator.SkipOffline = true
	expectedProc, err := gpsgen.NewDevice(opts)
	require.NoError(t, err)
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())

	type args struct {
		ctx        context.Context
		trackerID  types.ID
		routeIndex int
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    func() types.Navigator
		want1   bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(expectedProc, true)
					return mock
				},
			},
			args: args{
				ctx:        context.Background(),
				trackerID:  expectedTrackerID,
				routeIndex: 1,
			},
			wantErr: nil,
			want1:   true,
			want: func() types.Navigator {
				return types.NavigatorFromProc(expectedProc)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, got1, err := g.MoveToRoute(tt.args.ctx, tt.args.trackerID, tt.args.routeIndex)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.MoveToRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Generator.MoveToRoute() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Generator.MoveToRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerator_MoveToRouteByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	opts := gpsgen.NewDeviceOptions()
	opts.Navigator.SkipOffline = true
	expectedProc, err := gpsgen.NewDevice(opts)
	require.NoError(t, err)
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())
	r2 := gpsgen.RandomRouteForMoscow()
	expectedProc.AddRoute(r2)

	type args struct {
		ctx       context.Context
		trackerID types.ID
		routeID   types.ID
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    func() types.Navigator
		want1   bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when routeID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(r2.ID()),
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(expectedProc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   s2id(r2.ID()),
			},
			wantErr: nil,
			want1:   true,
			want: func() types.Navigator {
				return types.NavigatorFromProc(expectedProc)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, got1, err := g.MoveToRouteByID(tt.args.ctx, tt.args.trackerID, tt.args.routeID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.MoveToRouteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Generator.MoveToRouteByID() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Generator.MoveToRouteByIDs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerator_MoveToTrack(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	opts := gpsgen.NewDeviceOptions()
	opts.Navigator.SkipOffline = true
	expectedProc, err := gpsgen.NewDevice(opts)
	require.NoError(t, err)
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())
	r2 := gpsgen.RandomRouteForMoscow()
	expectedProc.AddRoute(r2)

	type args struct {
		ctx        context.Context
		trackerID  types.ID
		routeIndex int
		trackIndex int
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    func() types.Navigator
		want1   bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(expectedProc, true)
					return mock
				},
			},
			args: args{
				ctx:        context.Background(),
				trackerID:  expectedTrackerID,
				routeIndex: 1,
				trackIndex: 1,
			},
			wantErr: nil,
			want1:   true,
			want: func() types.Navigator {
				return types.NavigatorFromProc(expectedProc)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, got1, err := g.MoveToTrack(tt.args.ctx, tt.args.trackerID, tt.args.routeIndex, tt.args.trackIndex)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.MoveToTrack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Generator.MoveToTrack() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Generator.MoveToTrack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerator_MoveToTrackByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	opts := gpsgen.NewDeviceOptions()
	opts.Navigator.SkipOffline = true
	expectedProc, err := gpsgen.NewDevice(opts)
	require.NoError(t, err)
	expectedProc.AddRoute(gpsgen.RandomRouteForMoscow())
	r2 := gpsgen.RandomRouteForMoscow()
	expectedProc.AddRoute(r2)

	expectedTrackerID := types.NewID()
	expectedRouteID := s2id(r2.ID())
	expectedTrackID := s2id(r2.TrackAt(1).ID())

	type args struct {
		ctx       context.Context
		trackerID types.ID
		routeID   types.ID
		trackID   types.ID
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    func() types.Navigator
		want1   bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when routeID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when trackID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   expectedRouteID,
				trackID:   types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   expectedRouteID,
				trackID:   expectedTrackID,
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(expectedProc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				routeID:   expectedRouteID,
				trackID:   expectedTrackID,
			},
			wantErr: nil,
			want1:   true,
			want: func() types.Navigator {
				return types.NavigatorFromProc(expectedProc)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, got1, err := g.MoveToTrackByID(tt.args.ctx, tt.args.trackerID, tt.args.routeID, tt.args.trackID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.MoveToTrackByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Generator.MoveToTrackByID() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Generator.MoveToTrackByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerator_MoveToSegment(t *testing.T) {
	ctrl := gomock.NewController(t)

	opts := gpsgen.NewDeviceOptions()
	opts.Navigator.SkipOffline = true
	expectedProc, err := gpsgen.NewDevice(opts)
	require.NoError(t, err)
	r1 := navigator.RouteFromTracks(track300m1segment)
	expectedProc.AddRoute(r1)
	r2 := navigator.RouteFromTracks(track300m1segment, track300m1segment)
	expectedProc.AddRoute(r2)

	expectedTrackerID := types.NewID()

	type args struct {
		ctx          context.Context
		trackerID    types.ID
		routeIndex   int
		trackIndex   int
		segmentIndex int
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    func() types.Navigator
		want1   bool
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					return nil
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(expectedProc, true)
					return mock
				},
			},
			args: args{
				ctx:          context.Background(),
				trackerID:    expectedTrackerID,
				routeIndex:   1,
				trackIndex:   1,
				segmentIndex: 0,
			},
			wantErr: nil,
			want1:   true,
			want: func() types.Navigator {
				return types.NavigatorFromProc(expectedProc)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, got1, err := g.MoveToSegment(tt.args.ctx, tt.args.trackerID, tt.args.routeIndex, tt.args.trackIndex, tt.args.segmentIndex)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.MoveToSegment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Generator.MoveToSegment() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Generator.MoveToSegment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerator_AddSensor(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()

	type args struct {
		ctx       context.Context
		trackerID types.ID
		sensor    *gpsgen.Sensor
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is paused",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrTrackerOff,
		},
		{
			name: "should return error when storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Times(1).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Times(1).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					proc := gpsgen.NewAnimalTracker()
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(proc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.AddSensor(tt.args.ctx, tt.args.trackerID, tt.args.sensor); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.AddSensor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_RemoveSensor(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	expectedSensor, err := gpsgen.NewSensor("s1", 1, 10, 8, stdtypes.WithSensorRandomMode)
	require.NoError(t, err)

	type args struct {
		ctx       context.Context
		trackerID types.ID
		sensorID  types.ID
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
		want    bool
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when sensorID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				sensorID:  types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				sensorID:  s2id(expectedSensor.ID()),
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is paused",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				sensorID:  s2id(expectedSensor.ID()),
			},
			wantErr: generator.ErrTrackerOff,
		},
		{
			name: "should return error when sensor not found",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Return(trk, nil)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				sensorID:  s2id(expectedSensor.ID()),
			},
			wantErr: generator.ErrSensorNotFound,
		},
		{
			name: "should return error when storage.Update failure",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					trk.AddSensor(expectedSensor)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Times(1).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				sensorID:  s2id(expectedSensor.ID()),
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should not return error when all params are valid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					trk.AddSensor(expectedSensor)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Times(1).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					proc := gpsgen.NewAnimalTracker()
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(proc, true)
					return mock
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
				sensorID:  s2id(expectedSensor.ID()),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.RemoveSensor(tt.args.ctx, tt.args.trackerID, tt.args.sensorID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.RemoveSensor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generator.RemoveSensor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_Sensors(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}

	expectedTrackerID := types.NewID()
	expectedSensors := make([]*gpsgen.Sensor, 0)
	for i := 0; i < 10; i++ {
		sensor, err := gpsgen.NewSensor("s", 0, 1, 8, 0)
		require.NoError(t, err)
		expectedSensors = append(expectedSensors, sensor)
	}

	tests := []struct {
		name    string
		fields  mocks
		args    args
		want    []*gpsgen.Sensor
		wantErr error
	}{
		{
			name: "should return error when trackerID is empty",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should not return error when all params are valid",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := new(generator.Tracker)
					trk.AddSensor(expectedSensors...)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					return mock
				},
			},
			want:    expectedSensors,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			got, err := g.Sensors(tt.args.ctx, tt.args.trackerID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.Sensors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.Sensors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_ShutdownTracker(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackerID is invalid",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker is not running",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Stopped)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					return mock
				},
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should return error when process not found - invalid invariant",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Lookup(expectedTrackerID.String()).Times(1).Return(nil, false)
					return mock
				},
			},
			wantErr: generator.ErrTrackerNotRunning,
		},
		{
			name: "should return error when storage.Update failure",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					proc := gpsgen.NewAnimalTracker()
					mock.EXPECT().Lookup(gomock.Any()).Times(1).Return(proc, true)
					mock.EXPECT().Detach(gomock.All())
					return mock
				},
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should not return error when all params are valid",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					proc := gpsgen.NewAnimalTracker()
					mock.EXPECT().Lookup(gomock.Any()).Times(1).Return(proc, true)
					mock.EXPECT().Detach(gomock.All())
					return mock
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.ShutdownTracker(tt.args.ctx, tt.args.trackerID); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.ShutdownTracker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_ResumeTracker(t *testing.T) {
	ctrl := gomock.NewController(t)

	expectedTrackerID := types.NewID()
	expectedProc := gpsgen.NewAnimalTracker()

	type args struct {
		ctx       context.Context
		trackerID types.ID
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
	}{
		{
			name: "should return error when trackerID is invalid",
			args: args{
				ctx:       context.Background(),
				trackerID: types.ID{},
			},
			wantErr: types.ErrInvalidID,
		},
		{
			name: "should return error when tracker not found",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(nil, generator.ErrTrackerNotFound)
					return mock
				},
			},
			wantErr: generator.ErrTrackerNotFound,
		},
		{
			name: "should return error when tracker not paused",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Running)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					return mock
				},
			},
			wantErr: generator.ErrTrackerNotPaused,
		},
		{
			name: "should return error when storage.Update failure",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					proc := gpsgen.NewAnimalTracker()
					trk.ShutdownProcess(proc)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(generator.ErrInvalidTrackerVersion)
					return mock
				},
			},
			wantErr: generator.ErrInvalidTrackerVersion,
		},
		{
			name: "should not return error when all params are valid",
			args: args{
				ctx:       context.Background(),
				trackerID: expectedTrackerID,
			},
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					trk := makeValidTracker(t, types.Paused)
					trk.ShutdownProcess(expectedProc)
					mock := mockgenerator.NewMockStorage(ctrl)
					mock.EXPECT().FindTracker(gomock.Any(), expectedTrackerID).Times(1).Return(trk, nil)
					mock.EXPECT().Update(gomock.Any(), trk).Return(nil)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					mock := mockgenerator.NewMockProcesses(ctrl)
					mock.EXPECT().Attach(gomock.Any()).Times(1).Return(nil)
					return mock
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.ResumeTracker(tt.args.ctx, tt.args.trackerID); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.ResumeTracker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_Run(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name    string
		fields  mocks
		wantErr error
	}{
		{
			name: "should return error when loading trackers failure",
			fields: mocks{
				bootstraper: func() *mockgenerator.MockBootstraper {
					mock := mockgenerator.NewMockBootstraper(ctrl)
					mock.EXPECT().LoadTrackers(gomock.Any(), gomock.Any()).Return(generator.ErrLoadingTracker)
					return mock
				},
			},
			wantErr: generator.ErrLoadingTracker,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.Run(context.Background()); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_Close(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  mocks
		args    args
		wantErr error
	}{
		{
			name: "should return error when unloading trackers failure",
			fields: mocks{
				bootstraper: func() *mockgenerator.MockBootstraper {
					mock := mockgenerator.NewMockBootstraper(ctrl)
					mock.EXPECT().UnloadTrackers(gomock.Any(), gomock.Any()).Return(generator.ErrUnloadingTracker)
					return mock
				},
			},
			wantErr: generator.ErrUnloadingTracker,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGeneratorFromMocks(tt.fields)
			if err := g.Close(tt.args.ctx); !errors.Is(err, tt.wantErr) {
				t.Errorf("Generator.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
