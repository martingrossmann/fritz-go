package main

import (
	_ "github.com/martingrossmann/fritz-go/cmd"
	"github.com/martingrossmann/fritz-go/common"
	"github.com/martingrossmann/fritz-go/fritz"
	"github.com/martingrossmann/fritz-go/writer"
	"io"
	"log"
	"os"

	"github.com/magiconair/properties"
)

func main() {
	initLogging()
	conf := loadConfig()

	fritz := &fritz.FritzBox{
		Host:     conf.FritzHost,
		Insecure: false,
		Passw:    conf.FritzPassword,
	}

	fritz.PerformLogin()
	counter, err := fritz.ReadOnlineCounter()
	checkErr(err, "Cannot handle online counter data from Fritz.Box")

	writer.WriteToCSV(conf, counter)

	//cmd.Exec()
}

func initLogging() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err, "Cannot open logging file")
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
}

func loadConfig() common.Config {
	conf := common.Config{}
	p := properties.MustLoadFile("settings.conf", properties.UTF8)
	err := p.Decode(&conf)
	checkErr(err, "Error at parsing settings.conf: ")
	return conf
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
