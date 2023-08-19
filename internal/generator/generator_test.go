package generator_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/properties"
	stdtypes "github.com/mmadfox/go-gpsgen/types"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
	mockgenerator "github.com/mmadfox/gpsgend/tests/mocks/generator"
	"github.com/stretchr/testify/require"
)

type mocks struct {
	storage   func() *mockgenerator.MockStorage
	processes func() *mockgenerator.MockProcesses
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
			g := generator.New(
				tt.fields.storage(),
				tt.fields.processes(),
			)
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
			g := generator.New(tt.fields.storage(), tt.fields.processes())
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
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					return nil
				},
				processes: func() *mockgenerator.MockProcesses {
					return nil
				},
			},
			args: args{
				ctx:       context.Background(),
				trackerID: types.NewID(),
				opts:      generator.UpdateTrackerOptions{},
			},
			wantErr: true,
		},
		{
			name: "should return error when options invalid",
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					return nil
				},
				processes: func() *mockgenerator.MockProcesses {
					return nil
				},
			},
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
					mock.EXPECT().FindTracker(gomock.Any(), gomock.Any()).Times(1).Return(nil, generator.ErrTrackNotFound)
					return mock
				},
				processes: func() *mockgenerator.MockProcesses {
					return nil
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
				processes: func() *mockgenerator.MockProcesses {
					return nil
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
				processes: func() *mockgenerator.MockProcesses {
					return nil
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
			g := generator.New(tt.fields.storage(), tt.fields.processes())
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
			fields: mocks{
				storage: func() *mockgenerator.MockStorage {
					return nil
				},
				processes: func() *mockgenerator.MockProcesses {
					return nil
				},
			},
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
				processes: func() *mockgenerator.MockProcesses {
					return nil
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
			g := generator.New(tt.fields.storage(), tt.fields.processes())
			_, err := g.FindTracker(tt.args.ctx, tt.args.trackerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.FindTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
