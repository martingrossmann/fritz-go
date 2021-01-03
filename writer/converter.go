package writer

import (
	"encoding/csv"
	"fmt"
	"github.com/martingrossmann/fritz-go/common"
	"log"
	"os"
	"time"
)

//var csvFileName = ""

func ConvertCsvToInfluxData(conf common.Config) {
	csvFileName := conf.CsvFile
	influxFileName := conf.InfluxImportFile
	log.Println("Convert all data from ", csvFileName, " to InfluxDB import file ", influxFileName)

	csvFile, err := os.OpenFile(csvFileName, os.O_RDONLY, 0644)
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	countData := len(records) - 1
	checkErr(err, "Cannot read the CSV file.")
	log.Println((countData), "data in CSV file.")

	if countData < 1 {
		log.Fatal("No data found to convert")
	}

	dataFormat := "%s,host=%s %s=%si,%s=%si %d"
	timeLayout := "2006-01-02T03:04:05"

	for i, record := range records {
		// Ignore the header data
		if i != 0 {

			// Add 12 am to time stamp as counter.fetchCounterInfo() does the same for InfluxDB
			locBerlin, _ := time.LoadLocation("Europe/Berlin")
			t, err := time.ParseInLocation(timeLayout, record[0]+"T12:00:00", locBerlin)
			if err != nil {
				log.Fatal("Cannot parse time value ", record[0])
			}

			data := fmt.Sprintf(dataFormat,
				conf.InfluxTcMeasurement,
				"fritzbox",
				conf.InfluxTcReceived,
				record[1],
				conf.InfluxTcSent,
				record[2],
				t.Unix())

			log.Println(data)
		}
	}

	//influxFile, err :=os.OpenFile(influxFileName, os.O_CREATE|os.O_RDWR, 0644)
	//checkErr(err, "Cannot create InfluxDB file")

}
