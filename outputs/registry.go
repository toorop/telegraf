package outputs

import (
	"github.com/toorop/influxdb/client"
)

type Output interface {
	Connect() error
	Close() error
	Description() string
	SampleConfig() string
	Write(client.BatchPoints) error
}

type Creator func() Output

var Outputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Outputs[name] = creator
}
