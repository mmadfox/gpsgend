package stubdevice

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/gpsgend/internal/device"
)

func DeviceWithID(id uuid.UUID) *device.Device {
	newDevice, err := device.NewBuilder().
		ID(id).
		Status(device.Stopped).
		UserID(uuid.NewString()).
		Color(Color()).
		Model("testDevice").
		Props(map[string]string{"foo": "bar"}).
		Description("some description").
		Speed(1, 7, gpsgen.Amplitude4).
		Battery(1, 100, time.Hour).
		Elevation(1, 150, gpsgen.Amplitude16).
		Offline(1, 30).
		Routes(Routes()).
		Sensors(Sensors()).
		CreatedAt(time.Now().Add(-time.Hour)).
		UpdatedAt(time.Now()).
		Build()
	if err != nil {
		panic(err)
	}

	device.SetVersion(newDevice, 5)

	return newDevice
}

func DeviceWithoutRoutes() *device.Device {
	newDevice, err := device.NewBuilder().
		ID(uuid.New()).
		Status(device.Stopped).
		UserID(uuid.NewString()).
		Color(Color()).
		Model("testDevice").
		Props(map[string]string{"foo": "bar"}).
		Description("some description").
		Speed(1, 7, gpsgen.Amplitude4).
		Battery(1, 100, time.Hour).
		Elevation(1, 150, gpsgen.Amplitude16).
		Offline(1, 30).
		Sensors(Sensors()).
		CreatedAt(time.Now().Add(-time.Hour)).
		UpdatedAt(time.Now()).
		Build()
	if err != nil {
		panic(err)
	}

	device.SetVersion(newDevice, 5)

	return newDevice
}

func Device() *device.Device {
	newDevice, err := device.NewBuilder().
		ID(uuid.New()).
		Status(device.Stopped).
		UserID(uuid.NewString()).
		Color(Color()).
		Model("testDevice").
		Props(map[string]string{"foo": "bar"}).
		Description("some description").
		Speed(1, 7, gpsgen.Amplitude4).
		Battery(1, 100, time.Hour).
		Elevation(1, 150, gpsgen.Amplitude16).
		Offline(1, 30).
		Routes(Routes()).
		Sensors(Sensors()).
		CreatedAt(time.Now().Add(-time.Hour)).
		UpdatedAt(time.Now()).
		Build()
	if err != nil {
		panic(err)
	}

	device.SetVersion(newDevice, 5)

	return newDevice
}

func Color() colorful.Color {
	color, err := colorful.Hex("#ff0000")
	if err != nil {
		panic(err)
	}
	return color
}
