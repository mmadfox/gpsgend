package config

import (
	"fmt"
	"os"
	"time"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/gpsgend/internal/broker"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Service string `yaml:"service"`

	Logger struct {
		Format string `yaml:"format"`
		Level  string `yaml:"level"`
	} `yaml:"logger"`

	EventBroker struct {
		History struct {
			Enable        bool          `yaml:"enable"`
			TimePeriod    time.Duration `yaml:"timePeriod"`
			QueueCapacity int           `yaml:"queueCapacity"`
		} `yaml:"history"`
	} `yaml:"eventBroker"`

	Generator struct {
		FlushInterval time.Duration `yaml:"flushInterval"`
		// default 512
		PacketSize int `yaml:"packetSize"`
		// default num CPUs
		NumWorker int `yaml:"numWorker"`
	} `yaml:"generator"`

	Transport struct {
		GRPC struct {
			Listen string `yaml:"listen"`
		} `yaml:"grpc"`

		HTTP struct {
			Listen string `yaml:"listen"`
		} `yaml:"http"`

		Websocket struct {
			Listen string `yaml:"listen"`
		} `yaml:"websocket"`
	} `yaml:"transport"`

	Storage struct {
		Mongodb struct {
			URI            string `yml:"uri"`
			CollectionName string `yaml:"collectionName"`
			DatabaseName   string `yaml:"databaseName"`
		} `yaml:"mongodb"`
	} `yaml:"storage"`
}

func (c *Config) GeneratorOpts() *gpsgen.Options {
	opts := gpsgen.NewOptions()
	if c.Generator.FlushInterval > 0 {
		opts.Interval = c.Generator.FlushInterval
	} else {
		opts.Interval = 3 * time.Second
	}

	if c.Generator.PacketSize > 0 {
		opts.PacketSize = c.Generator.PacketSize
	}

	if c.Generator.NumWorker > 0 {
		opts.NumWorkers = c.Generator.NumWorker
	}

	return opts
}

func (c *Config) EventBrokerOpts() *broker.Options {
	opts := broker.DefaultOptions()

	opts.HistoryEnable = c.EventBroker.History.Enable
	opts.HistoryQueueCapacity = c.EventBroker.History.QueueCapacity
	opts.HistoryTimePeriod = c.EventBroker.History.TimePeriod

	return opts
}

func FromFile(filename string) (*Config, error) {
	conf := new(Config)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read gpsgend config [%s]: error %w", filename, err)
	}
	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, err
	}
	return conf, err
}
