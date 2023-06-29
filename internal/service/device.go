package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/go-gpsgen/random"
	"github.com/mmadfox/go-gpsgen/route"
	"github.com/mmadfox/gpsgend/internal/device"
	"github.com/mmadfox/gpsgend/internal/types"
)

type DeviceService struct {
	deviceStorage device.Storage
	generator     device.Generator
	publisher     device.Publisher
}

func NewDeviceService(
	deviceStorage device.Storage,
	publisher device.Publisher,
	generator device.Generator,
) *DeviceService {
	return &DeviceService{
		publisher:     publisher,
		deviceStorage: deviceStorage,
		generator:     generator,
	}
}

func (s *DeviceService) NewPresetsDevice(ctx context.Context, p device.NewPresetsDeviceParams) (*device.Device, error) {
	now := time.Now()
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
		UpdatedAt(now).
		CreatedAt(now).
		Props(p.Properties).
		Build()
	if err != nil {
		return nil, err
	}

	return s.createDevice(ctx, newDevice, p.WithRandomRoutes)
}

func (s *DeviceService) NewDevice(ctx context.Context, p device.NewDeviceParams) (*device.Device, error) {
	now := time.Now()
	newDevice, err := device.NewBuilder().
		ID(uuid.New()).
		Model(p.Model).
		Status(device.Stopped).
		UserID(p.UserID).
		Description(p.Description).
		Speed(p.Speed.Min, p.Speed.Max, p.Speed.Amplitude).
		Battery(p.Battery.Min, p.Battery.Max, time.Duration(p.Battery.ChargeTime)).
		Elevation(p.Elevation.Min, p.Elevation.Max, p.Elevation.Amplitude).
		Offline(p.Offline.Min, p.Offline.Max).
		Props(p.Properties).
		UpdatedAt(now).
		CreatedAt(now).
		Build()
	if err != nil {
		return nil, err
	}

	return s.createDevice(ctx, newDevice, p.WithRandomRoutes)
}

func (s *DeviceService) createDevice(ctx context.Context, newDevice *device.Device, randomRoutes bool) (*device.Device, error) {
	if randomRoutes {
		for i := 0; i < 3; i++ {
			newRoute, err := route.Generate()
			if err != nil {
				return nil, err
			}
			if err := newDevice.AddRoute(device.NewRoute(newRoute)); err != nil {
				return nil, err
			}
		}
	}

	newDevice.AssignColor()

	if err := s.deviceStorage.Insert(ctx, newDevice); err != nil {
		return nil, err
	}

	return newDevice, nil
}

func (s *DeviceService) UpdateDevice(ctx context.Context, deviceID uuid.UUID, p device.UpdateDeviceParams) (*device.Device, error) {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return nil, err
	}

	if p.Model != nil {
		model, err := types.NewModel(*p.Model)
		if err != nil {
			return nil, err
		}
		dev.ChangeModel(model)
	}

	if p.UserID != nil {
		dev.ChangeUserID(*p.UserID)
	}

	if p.Properties != nil {
		props, err := types.NewProperties(*p.Properties)
		if err != nil {
			return nil, err
		}
		dev.ChangeProps(props)
	}

	if p.Description != nil {
		descr, err := types.NewDescription(*p.Description)
		if err != nil {
			return nil, err
		}
		dev.ChangeDescription(descr)
	}

	if p.Speed != nil {
		speed, err := types.NewSpeed(
			p.Speed.Min,
			p.Speed.Max,
			p.Speed.Amplitude,
		)
		if err != nil {
			return nil, err
		}
		dev.ChangeSpeed(speed)
	}

	if p.Battery != nil {
		battery, err := types.NewBattery(
			p.Battery.Min,
			p.Battery.Max,
			time.Duration(p.Battery.ChargeTime)*time.Second,
		)
		if err != nil {
			return nil, err
		}
		dev.ChangeBattery(battery)
	}

	if p.Elevation != nil {
		elevation, err := types.NewElevation(
			p.Elevation.Min,
			p.Elevation.Max,
			p.Elevation.Amplitude,
		)
		if err != nil {
			return nil, err
		}
		dev.ChangeElevation(elevation)
	}

	if p.Offline != nil {
		offline, err := types.NewOffline(
			p.Offline.Min,
			p.Offline.Max,
		)
		if err != nil {
			return nil, err
		}
		dev.ChangeOfflineDuraiton(offline)
	}

	return dev, s.deviceStorage.Update(ctx, dev)
}

func (s *DeviceService) DeviceByID(ctx context.Context, deviceID uuid.UUID) (*device.Device, error) {
	return s.deviceStorage.FindByID(ctx, deviceID)
}

func (s *DeviceService) RunDevice(ctx context.Context, deviceID uuid.UUID) error {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	proc, err := dev.NewProcess()
	if err != nil {
		return err
	}
	s.bindDeviceProcess(proc)
	s.generator.Attach(proc)

	if err := s.deviceStorage.Update(ctx, dev); err != nil {
		s.generator.Detach(proc.ID())
		return err
	}

	return nil
}

func (s *DeviceService) StopDevice(ctx context.Context, deviceID uuid.UUID) error {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	if err := dev.RemoveProcess(); err != nil {
		return err
	}

	s.generator.Detach(dev.ID())

	return s.deviceStorage.Update(ctx, dev)
}

func (s *DeviceService) PauseDevice(ctx context.Context, deviceID uuid.UUID) error {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	if !dev.IsRunning() {
		return device.ErrDeviceAlreadyStopped
	}

	proc, err := s.generator.Lookup(dev.ID())
	if err != nil {
		return err
	}

	if err := dev.PauseProcess(proc); err != nil {
		return err
	}

	s.generator.Detach(dev.ID())

	return s.deviceStorage.Update(ctx, dev)
}

func (s *DeviceService) ResumeDevice(ctx context.Context, deviceID uuid.UUID) error {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	proc, err := dev.ResumeDevice()
	if err != nil {
		return err
	}
	s.bindDeviceProcess(proc)
	s.generator.Attach(proc)

	if err := s.deviceStorage.Update(ctx, dev); err != nil {
		s.generator.Detach(dev.ID())
		return err
	}

	return nil
}

func (s *DeviceService) RemoveDevice(ctx context.Context, deviceID uuid.UUID) error {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	if dev.IsRunning() {
		s.generator.Detach(dev.ID())
	}

	return s.deviceStorage.Delete(ctx, deviceID)
}

func (s *DeviceService) AddRoutes(ctx context.Context, deviceID uuid.UUID, routes []*navigator.Route) error {
	if len(routes) == 0 {
		return fmt.Errorf("routes are empty for device %s", deviceID)
	}

	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(routes); i++ {
		route := device.NewRoute(routes[i])
		if err := dev.AddRoute(route); err != nil {
			return err
		}
	}

	return s.deviceStorage.Update(ctx, dev)
}

func (s *DeviceService) Routes(ctx context.Context, deviceID uuid.UUID) ([]*device.Route, error) {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return nil, err
	}

	return dev.Routes(), nil
}

func (s *DeviceService) RemoveRoute(ctx context.Context, deviceID uuid.UUID, routeID uuid.UUID) error {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	if err := dev.RemoveRoute(routeID); err != nil {
		return err
	}

	return s.deviceStorage.Update(ctx, dev)
}

func (s *DeviceService) Sensors(ctx context.Context, deviceID uuid.UUID) ([]types.Sensor, error) {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return nil, err
	}

	return dev.Sensors(), nil
}

func (s *DeviceService) AddSensor(ctx context.Context, deviceID uuid.UUID, p device.AddSensorParams) ([]types.Sensor, error) {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return nil, err
	}

	if err := dev.AddSensor(p.Name, p.Min, p.Max, p.Amplitude); err != nil {
		return nil, err
	}

	if err := s.deviceStorage.Update(ctx, dev); err != nil {
		return nil, err
	}

	return dev.Sensors(), nil
}

func (s *DeviceService) RemoveSensor(ctx context.Context, deviceID uuid.UUID, sensorID uuid.UUID) error {
	dev, err := s.deviceStorage.FindByID(ctx, deviceID)
	if err != nil {
		return err
	}

	if err := dev.RemoveSensor(sensorID); err != nil {
		return err
	}

	return s.deviceStorage.Update(ctx, dev)
}

func (s *DeviceService) Bootstrap(ctx context.Context) error {
	status := []device.Status{device.Stored}
	return s.deviceStorage.WalkByStatus(ctx, status,
		func(dev *device.Device) error {
			proc, err := dev.ResumeDevice()
			if err != nil {
				return err
			}

			s.bindDeviceProcess(proc)
			s.generator.Attach(proc)

			return s.deviceStorage.Update(ctx, dev)
		})
}

func (s *DeviceService) Close(ctx context.Context) error {
	status := []device.Status{device.Running}
	return s.deviceStorage.WalkByStatus(ctx, status,
		func(dev *device.Device) error {
			proc, _ := s.generator.Lookup(dev.ID())
			if proc == nil {
				return nil
			}

			s.bindDeviceProcess(proc)

			if err := dev.StoreProcess(proc); err != nil {
				return err
			}

			return s.deviceStorage.Update(ctx, dev)
		})
}

func (s *DeviceService) bindDeviceProcess(proc *gpsgen.Device) {
	ctx := context.Background()
	proc.OnStateChangeBytes = func(data []byte) {
		_ = s.publisher.Publish(ctx, data)
	}
}
