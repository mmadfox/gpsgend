package mongo

import (
	"testing"
	"time"

	"github.com/mmadfox/gpsgend/internal/device"
	stubdevice "github.com/mmadfox/gpsgend/tests/stubs/device"
	"github.com/stretchr/testify/require"
)

func Test_encodeDeviceModel(t *testing.T) {
	dev := stubdevice.Device()
	deviceModel, err := encodeDevice(dev)
	require.NoError(t, err)
	require.NotNil(t, deviceModel)

	batteryChargeTime, err := time.ParseDuration(deviceModel.Battery.ChargeTime)
	require.NoError(t, err)

	require.Equal(t, dev.ID().String(), deviceModel.ID)
	require.Equal(t, dev.UserID(), deviceModel.UserID)
	require.Equal(t, dev.Model().String(), deviceModel.Model)
	require.Equal(t, dev.Color().Hex(), deviceModel.Color)
	require.Equal(t, dev.Description().String(), deviceModel.Description)
	require.Equal(t, dev.Speed().Min(), deviceModel.Speed.Min)
	require.Equal(t, dev.Speed().Max(), deviceModel.Speed.Max)
	require.Equal(t, dev.Speed().Amplitude(), deviceModel.Speed.Amplitude)
	require.Equal(t, dev.Battery().Min(), deviceModel.Battery.Min)
	require.Equal(t, dev.Battery().Max(), deviceModel.Battery.Max)
	require.Equal(t, dev.Battery().ChargeTime(), batteryChargeTime)
	require.Equal(t, dev.Elevation().Min(), deviceModel.Elevation.Min)
	require.Equal(t, dev.Elevation().Max(), deviceModel.Elevation.Max)
	require.Equal(t, dev.Elevation().Amplitude(), deviceModel.Elevation.Amplitude)
	require.Equal(t, dev.Offline().Min(), deviceModel.Offline.Min)
	require.Equal(t, dev.Offline().Max(), deviceModel.Offline.Max)
	require.Equal(t, dev.Props().ToMap(), deviceModel.Props)
	require.EqualValues(t, dev.Status(), deviceModel.Status.ID)
	require.Equal(t, dev.Status().String(), deviceModel.Status.Text)

	sensors := dev.Sensors()
	require.Equal(t, len(sensors), deviceModel.NumSensors)
	totalSensors := 0
	for i := 0; i < len(sensors); i++ {
		totalSensors++
		sensor := sensors[i]
		require.Equal(t, sensor.ID().String(), deviceModel.Sensors[i].ID)
		require.Equal(t, sensor.Name(), deviceModel.Sensors[i].Name)
		require.Equal(t, sensor.Min(), deviceModel.Sensors[i].Min)
		require.Equal(t, sensor.Max(), deviceModel.Sensors[i].Max)
		require.Equal(t, sensor.Amplitude(), deviceModel.Sensors[i].Amplitude)
	}
	require.NotZero(t, totalSensors)

	routes := dev.Routes()
	require.Equal(t, len(routes), deviceModel.NumRoutes)
	totalRoutes := 0
	for i := 0; i < len(routes); i++ {
		totalRoutes++
		route := routes[i]
		require.Equal(t, route.ID().String(), deviceModel.Routes[i].ID)
		require.Equal(t, route.Color().Hex(), deviceModel.Routes[i].Color)
		routeData, err := route.Route().MarshalBinary()
		require.NoError(t, err)
		require.Equal(t, routeData, deviceModel.Routes[i].Route)
	}
	require.NotZero(t, totalRoutes)

	require.Equal(t, dev.CreatedAt().UnixNano(), deviceModel.CreatedAt)
	require.Equal(t, dev.UpdatedAt().UnixNano(), deviceModel.UpdatedAt)
	require.Equal(t, device.Version(dev), deviceModel.Version)
}

func Test_decodeDevice(t *testing.T) {
	dev := stubdevice.Device()
	newModel := func() *deviceModel {
		model, err := encodeDevice(dev)
		require.NoError(t, err)
		return model
	}
	type args struct {
		model func() *deviceModel
	}
	tests := []struct {
		name    string
		args    args
		want    *device.Device
		wantErr bool
	}{
		{
			name: "should return error when device id is invalid",
			args: args{
				model: func() *deviceModel {
					model := newModel()
					model.ID = ""
					return model
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when color is invalid",
			args: args{
				model: func() *deviceModel {
					model := newModel()
					model.Color = "someColor"
					return model
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when sensor id is invalid",
			args: args{
				model: func() *deviceModel {
					model := newModel()
					model.Sensors[0].ID = "someID"
					return model
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when route data is invalid",
			args: args{
				model: func() *deviceModel {
					model := newModel()
					model.Routes[0].Route = []byte(`somedata`)
					return model
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when route id is invalid",
			args: args{
				model: func() *deviceModel {
					model := newModel()
					model.Routes[0].ID = "someID"
					return model
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when route color is invalid",
			args: args{
				model: func() *deviceModel {
					model := newModel()
					model.Routes[0].Color = "someColor"
					return model
				},
			},
			wantErr: true,
		},
		{
			name: "should return device",
			args: args{
				model: newModel,
			},
			want: dev,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeDevice(tt.args.model())
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			require.Equal(t, dev.ID().String(), got.ID().String())
			require.Equal(t, dev.UserID(), got.UserID())
			require.Equal(t, dev.Model(), got.Model())
			require.Equal(t, dev.Color(), got.Color())
			require.Equal(t, dev.Description(), got.Description())
			require.Equal(t, dev.Speed(), got.Speed())
			require.Equal(t, dev.Battery(), got.Battery())
			require.Equal(t, dev.Elevation(), got.Elevation())
			require.Equal(t, dev.Offline(), got.Offline())
			require.Equal(t, dev.Props(), got.Props())
			require.Equal(t, dev.Sensors(), got.Sensors())
			require.Equal(t, dev.CreatedAt().Unix(), got.CreatedAt().Unix())
			require.Equal(t, dev.UpdatedAt().Unix(), got.UpdatedAt().Unix())
			require.Equal(t, device.Version(dev), device.Version(got))

			wantRoutes := dev.Routes()
			gotRoutes := got.Routes()
			require.Equal(t, len(wantRoutes), len(gotRoutes))
			for i := 0; i < len(wantRoutes); i++ {
				wantRoute := wantRoutes[i]
				gotRoute := gotRoutes[i]
				require.Equal(t, wantRoute.ID(), gotRoute.ID())
				require.Equal(t, wantRoute.Color().Hex(), gotRoute.Color().Hex())
				wantRouteData, err := wantRoute.Route().MarshalBinary()
				require.NoError(t, err)
				gotRouteData, err := gotRoute.Route().MarshalBinary()
				require.NoError(t, err)
				require.Equal(t, wantRouteData, gotRouteData)
			}
		})
	}
}
