package remofwd

import (
	"context"
	"os"
	"strings"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

type Forwarder struct {
	write influxdb2.WriteApiBlocking
}

func NewForwarder() Forwarder {
	client := influxdb2.NewClient(os.Getenv("INFLUXDB_CONN"), "")
	writeApi := client.WriteApiBlocking("", os.Getenv("INFLUXDB_DB"))

	return Forwarder{
		write: writeApi,
	}
}

func (f *Forwarder) SendEvent(device string, event *Event) error {
	p := influxdb2.NewPoint(strings.ToLower(event.Name()), map[string]string{
		"device": device,
	}, map[string]interface{}{
		"value": event.Value(),
	}, event.Timestamp())

	return f.write.WritePoint(context.TODO(), p)
}
