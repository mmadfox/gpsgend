package service_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/random"
	"github.com/mmadfox/gpsgend/internal/device"
	"github.com/mmadfox/gpsgend/internal/service"
	"github.com/mmadfox/gpsgend/internal/types"
	mockdevice "github.com/mmadfox/gpsgend/tests/mocks/device"
	stubdevice "github.com/mmadfox/gpsgend/tests/stubs/device"
	stubservice "github.com/mmadfox/gpsgend/tests/stubs/service"
	"github.com/stretchr/testify/require"
)

func TestService_NewPresetsDevice(t *testing.T) {
	someErr := errors.New("some error")
	type args struct {
		params device.NewPresetsDeviceParams
	}
	tests := []struct {
		name    string
		args    args
		new     func(ctx context.Context) *service.DeviceService
		want    func(device.NewPresetsDeviceParams) *device.Device
		wantErr bool
	}{
		{
			name: "should return an error on storage failure",
			new: func(ctx context.Context) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().Insert(gomock.Any(), gomock.Any()).Times(1).Return(someErr)
				return svc
			},
			wantErr: true,
		},
		{
			name: "should return device when all parameters are empty",
			new: func(ctx context.Context) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().Insert(gomock.Any(), gomock.Any()).Times(1).Return(nil)
				return svc
			},
			want: func(p device.NewPresetsDeviceParams) *device.Device {
				return newPresetsDevice(p)
			},
		},
		{
			name: "should return valid device",
			args: args{
				params: device.NewPresetsDeviceParams{
					Description: "description",
					UserID:      "userID",
					Properties:  map[string]string{"foo": "bar"},
				},
			},
			new: func(ctx context.Context) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().Insert(gomock.Any(), gomock.Any()).Times(1).Return(nil)
				return svc
			},
			want: func(p device.NewPresetsDeviceParams) *device.Device {
				return newPresetsDevice(p)
			},
		},
		{
			name: "should return valid device with routes",
			args: args{
				params: device.NewPresetsDeviceParams{
					WithRandomRoutes: true,
				},
			},
			new: func(ctx context.Context) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().Insert(gomock.Any(), gomock.Any()).Times(1).Return(nil)
				return svc
			},
			want: func(p device.NewPresetsDeviceParams) *device.Device {
				return newPresetsDevice(p)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			svc := tt.new(ctx)
			got, err := svc.NewPresetsDevice(ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.NewPresetsDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			want := tt.want(tt.args.params)
			require.True(t, got.ID() != uuid.Nil)
			require.NotEmpty(t, got.Model().String())
			require.Equal(t, want.Status(), got.Status())
			require.Equal(t, want.UserID(), got.UserID())
			require.Equal(t, want.Description().String(), got.Description().String())
			require.Equal(t, want.Speed().String(), got.Speed().String())
			require.Equal(t, want.Battery().String(), got.Battery().String())
			require.Equal(t, want.Elevation().String(), got.Elevation().String())
			require.Equal(t, want.Offline().String(), got.Offline().String())
			gotProps := got.Props().ToMap()
			for k, v := range want.Props().ToMap() {
				require.Equal(t, v, gotProps[k])
			}
			require.Greater(t, got.CreatedAt().Unix(), int64(0))
			require.Greater(t, got.UpdatedAt().Unix(), int64(0))

			if tt.args.params.WithRandomRoutes {
				require.Len(t, got.Routes(), 3)
			}
		})
	}
}

func TestService_UpdateDevice(t *testing.T) {
	s2p := func(s string) *string {
		return &s
	}
	tests := []struct {
		name       string
		newService func(id uuid.UUID) *service.DeviceService
		params     func(*device.UpdateDeviceParams)
		assert     func(*device.Device)
		wantErr    error
	}{
		{
			name: "should return error when device not found",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(nil, device.ErrDeviceNotFound)
				return svc
			},
			wantErr: device.ErrDeviceNotFound,
		},
		{
			name: "should return error when invalid device model",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(stubdevice.DeviceWithID(id), nil)
				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Model = s2p("")
			},
			wantErr: types.ErrInvalidMinValue,
		},
		{
			name: "should return valid device with changed model",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().Update(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)
				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Model = s2p("Model2")
			},
			assert: func(dev *device.Device) {
				require.Equal(t, "Model2", dev.Model().String())
			},
		},
		{
			name: "should return valid device with changed userID",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().Update(gomock.Any(), dev).
					Times(1).
					Return(nil)
				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.UserID = s2p("UserID2")
			},
			assert: func(dev *device.Device) {
				require.Equal(t, "UserID2", dev.UserID())
			},
		},
		{
			name: "should return error when invalid properties",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				props := make(map[string]string, 20)
				for i := 0; i < 20; i++ {
					pk := fmt.Sprintf("p%d", i)
					props[pk] = pk
				}
				p.Properties = &props
			},
			wantErr: types.ErrInvalidMaxValue,
		},
		{
			name: "should return valid device with changed properties",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().Update(gomock.Any(), dev).
					Times(1).
					Return(nil)
				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				props := make(map[string]string, types.MaxPropertiesValue-1)
				for i := 0; i < types.MaxPropertiesValue-1; i++ {
					pk := fmt.Sprintf("p%d", i)
					props[pk] = pk
				}
				p.Properties = &props
			},
			assert: func(dev *device.Device) {
				props := dev.Props().ToMap()
				for i := 0; i < types.MaxPropertiesValue-1; i++ {
					pk := fmt.Sprintf("p%d", i)
					_, ok := props[pk]
					require.True(t, ok)
				}
			},
		},
		{
			name: "should return error when invalid description",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Description = s2p(strings.Repeat("a", types.MaxDescriptionValue+1))
			},
			wantErr: types.ErrInvalidMaxValue,
		},
		{
			name: "should return valid device with changed description",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().Update(gomock.Any(), dev).
					Times(1).
					Return(nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Description = s2p("description")
			},
			assert: func(dev *device.Device) {
				require.Equal(t, "description", dev.Description().String())
			},
		},
		{
			name: "should return error when invalid speed",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Speed.Max = types.MaxSpeedValue + 1
			},
			wantErr: types.ErrInvalidMaxValue,
		},
		{
			name: "should return valid device with changed speed",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().Update(gomock.Any(), dev).
					Times(1).
					Return(nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Speed.Min = 1
				p.Speed.Max = 2
				p.Speed.Amplitude = 8
			},
			assert: func(dev *device.Device) {
				require.Equal(t, float64(1), dev.Speed().Min())
				require.Equal(t, float64(2), dev.Speed().Max())
				require.Equal(t, 8, dev.Speed().Amplitude())
			},
		},
		{
			name: "should return error when invalid battery",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Battery.ChargeTime = int64(types.MaxBatteryChargeTime.Seconds()) + 1
			},
			wantErr: types.ErrInvalidMaxChargeTime,
		},
		{
			name: "should return error when invalid elevation",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Elevation.Max = types.MaxElevationValue + 1
			},
			wantErr: types.ErrInvalidMaxValue,
		},
		{
			name: "should return error when invalid offline",
			newService: func(id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().FindByID(gomock.Any(), id).
					Times(1).
					Return(dev, nil)

				return svc
			},
			params: func(p *device.UpdateDeviceParams) {
				p.Offline.Max = types.MaxOfflineValue + 10
			},
			wantErr: types.ErrInvalidMaxValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			id := uuid.New()
			params := stubservice.UpdateDeviceParams()
			if tt.params != nil {
				tt.params(&params)
			}
			svc := tt.newService(id)
			got, err := svc.UpdateDevice(ctx, id, params)
			if err != nil {
				require.Equal(t, tt.wantErr, err)
			}
			if tt.assert != nil {
				tt.assert(got)
			}
		})
	}
}

func TestService_DeviceByID(t *testing.T) {
	svc, mock := newService(t)
	ctx := context.Background()
	dev := stubdevice.Device()

	mock.deviceStorage.EXPECT().FindByID(ctx, dev.ID()).Return(dev, nil)

	got, err := svc.DeviceByID(ctx, dev.ID())
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, dev, got)
}

func TestService_RunDevice(t *testing.T) {
	someErr := errors.New("some error")
	tests := []struct {
		name       string
		newService func(context.Context, uuid.UUID) *service.DeviceService
		wantErr    error
	}{
		{
			name: "should return error when device not found",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(nil, device.ErrDeviceNotFound)
				return svc
			},
			wantErr: device.ErrDeviceNotFound,
		},
		{
			name: "should return error when device is running",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				proc, err := dev.NewProcess()
				require.NoError(t, err)
				require.NotNil(t, proc)
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(dev, nil)
				return svc
			},
			wantErr: device.ErrDeviceAlreadyRunning,
		},
		{
			name: "should return error when routes not assigned",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithoutRoutes()
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(dev, nil)
				return svc
			},
			wantErr: device.ErrNoRoutes,
		},
		{
			name: "should return error when updating device in storage on failure",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().
					FindByID(ctx, dev.ID()).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().
					Update(gomock.Any(), dev).
					Times(1).
					Return(someErr)
				mock.generator.EXPECT().
					Attach(gomock.Any()).
					Times(1)
				mock.generator.EXPECT().
					Detach(dev.ID()).
					Times(1)
				return svc
			},
			wantErr: someErr,
		},
		{
			name: "should not return error",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				mock.deviceStorage.EXPECT().
					FindByID(ctx, dev.ID()).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().
					Update(gomock.Any(), dev).
					Times(1).
					Return(nil)
				mock.generator.EXPECT().
					Attach(gomock.Any()).
					Times(1)
				return svc
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			id := uuid.New()
			svc := tt.newService(ctx, id)
			err := svc.RunDevice(ctx, id)
			if err != nil {
				require.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestService_StopDevice(t *testing.T) {
	someErr := errors.New("some error")
	tests := []struct {
		name       string
		newService func(context.Context, uuid.UUID) *service.DeviceService
		wantErr    error
	}{
		{
			name: "should return error when device not found",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(nil, device.ErrDeviceNotFound)
				return svc
			},
			wantErr: device.ErrDeviceNotFound,
		},
		{
			name: "should return error when device is stopped",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				require.True(t, device.Running != dev.Status())
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(dev, nil)
				return svc
			},
			wantErr: device.ErrDeviceAlreadyStopped,
		},
		{
			name: "should return error when updating device in storage on failure",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				// start device
				proc, err := dev.NewProcess()
				require.NoError(t, err)
				require.NotNil(t, proc)

				mock.deviceStorage.EXPECT().
					FindByID(ctx, dev.ID()).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().
					Update(gomock.Any(), dev).
					Times(1).
					Return(someErr)
				mock.generator.EXPECT().
					Detach(dev.ID()).
					Times(1)
				return svc
			},
			wantErr: someErr,
		},
		{
			name: "should not return error",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				// start device
				proc, err := dev.NewProcess()
				require.NoError(t, err)
				require.NotNil(t, proc)

				mock.deviceStorage.EXPECT().
					FindByID(ctx, dev.ID()).
					Times(1).
					Return(dev, nil)
				mock.deviceStorage.EXPECT().
					Update(gomock.Any(), dev).
					Times(1).
					Return(nil)
				mock.generator.EXPECT().
					Detach(dev.ID()).
					Times(1)
				return svc
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			id := uuid.New()
			svc := tt.newService(ctx, id)
			err := svc.StopDevice(ctx, id)
			if err != nil {
				require.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestService_PauseDevice(t *testing.T) {
	tests := []struct {
		name       string
		newService func(context.Context, uuid.UUID) *service.DeviceService
		wantErr    error
	}{
		{
			name: "should return error when device not found",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(nil, device.ErrDeviceNotFound)
				return svc
			},
			wantErr: device.ErrDeviceNotFound,
		},
		{
			name: "should return error when device is not running",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				require.True(t, device.Running != dev.Status())
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(dev, nil)
				return svc
			},
			wantErr: device.ErrDeviceAlreadyStopped,
		},
		{
			name: "should return error when running process not found",
			newService: func(ctx context.Context, id uuid.UUID) *service.DeviceService {
				svc, mock := newService(t)
				dev := stubdevice.DeviceWithID(id)
				proc, err := dev.NewProcess()
				require.NoError(t, err)
				require.NotNil(t, proc)
				mock.deviceStorage.EXPECT().
					FindByID(ctx, id).
					Times(1).
					Return(dev, nil)
				mock.generator.EXPECT().
					Lookup(dev.ID()).
					Return(nil, errors.New("not found"))
				return svc
			},
			wantErr: errors.New("not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			id := uuid.New()
			svc := tt.newService(ctx, id)
			err := svc.PauseDevice(ctx, id)
			if err != nil {
				require.Equal(t, tt.wantErr, err)
			}
		})
	}
}

type mocks struct {
	ctrl          *gomock.Controller
	generator     *mockdevice.MockGenerator
	deviceStorage *mockdevice.MockStorage
	queryStorage  *mockdevice.MockQuery
	publisher     *mockdevice.MockPublisher
}

func newService(t *testing.T) (*service.DeviceService, *mocks) {
	ctrl := gomock.NewController(t)
	m := &mocks{
		ctrl:          ctrl,
		generator:     mockdevice.NewMockGenerator(ctrl),
		deviceStorage: mockdevice.NewMockStorage(ctrl),
		queryStorage:  mockdevice.NewMockQuery(ctrl),
		publisher:     mockdevice.NewMockPublisher(ctrl),
	}
	return service.NewDeviceService(
		m.deviceStorage,
		m.publisher,
		m.generator,
	), m
}

func newPresetsDevice(p device.NewPresetsDeviceParams) *device.Device {
	newDevice, err := device.NewBuilder().
		ID(uuid.New()).
		Model(random.String(8)).
		Status(device.Stopped).
		UserID(p.UserID).
		Description(p.Description).
		Speed(1, 5, gpsgen.Amplitude8).
		Battery(1, 100, 7*time.Hour).
		Elevation(50, 100, gpsgen.Amplitude4).
		Offline(1, 120).
		Props(p.Properties).
		Build()
	if err != nil {
		panic(err)
	}
	return newDevice
}
