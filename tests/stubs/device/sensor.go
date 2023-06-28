package stubdevice

import (
	"fmt"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/gpsgend/internal/types"
)

func Sensors() []types.Sensor {
	sensors := make([]types.Sensor, 0)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("sensor-%d", i)
		sensor, err := types.NewSensor(name, 0, 100, gpsgen.Amplitude4)
		if err != nil {
			panic(err)
		}
		sensors = append(sensors, sensor)
	}
	return sensors
}
