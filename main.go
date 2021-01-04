package main

import (
	"github.com/martingrossmann/fritz-go/cmd"
	"io"
	"log"
	"os"
)

func main() {
	initLogging()
	//conf := loadConfig()

	//fritz := &fritz.FritzBox{
	//	Host:     conf.FritzHost,
	//	Insecure: false,
	//	Passw:    conf.FritzPassword,
	//}
	//
	//fritz.PerformLogin()
	//counter, err := fritz.ReadOnlineCounter()
	//checkErr(err, "Cannot handle online counter data from Fritz.Box")
	//
	//if conf.CsvActive {
	//	writer.WriteToCSV(conf, counter)
	//}
	//
	//if conf.InfluxActive {
	//	writer.WriteToInnflux(conf, counter)
	//}

	//writer.ConvertCsvToInfluxData(conf)

	cmd.Exec()
}

func initLogging() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Cannot open logging file", err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
}
