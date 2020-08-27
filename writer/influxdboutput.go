package writer

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/martingrossmann/fritz-go/common"
	"github.com/martingrossmann/fritz-go/fritz"
	"log"
)

func WriteToInnflux(conf common.Config, counter fritz.OnlineCounter) {
	log.Println("Trying to connect to influxdb " + conf.InfluxHost)
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient(conf.InfluxHost, conf.InfluxUser+":"+conf.InfluxPassword)
	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking("", conf.InfluxDB)

	// Create new point with counter data
	p := influxdb2.NewPointWithMeasurement(conf.InfluxTcMeasurement).
		AddTag("host", "fritzbox").
		AddField(conf.InfluxTcSent, counter.DataSent).
		AddField(conf.InfluxTcReceived, counter.DataReceived).
		SetTime(counter.DayOfData)
	writeAPI.WritePoint(context.Background(), p)
	log.Println("Counter data written to influxdb successfully.")
}
