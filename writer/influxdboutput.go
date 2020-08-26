package writer

import (
	"context"
	"github.com/martingrossmann/fritz-go/common"
	"github.com/martingrossmann/fritz-go/fritz"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func WriteToInnflux(conf common.Config, counter fritz.OnlineCounter) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:9999", "my-token")
	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking("my-org", "my-bucket")
	// Create point using fluent style
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)
}
