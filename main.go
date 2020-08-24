package main

import (
	_ "github.com/martingrossmann/fritz-go/cmd"
	"github.com/martingrossmann/fritz-go/fritz"
	"github.com/martingrossmann/fritz-go/writer"
	"log"
	"os"

	"github.com/magiconair/properties"
)

type Config struct {
	Host     string `properties:"host,default=http://fritz.box"`
	Password string `properties:"password"`
}

func main() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err, "Cannot open logging file")
	log.SetOutput(file)

	conf := loadConfig()

	fritz := &fritz.FritzBox{
		Host:     conf.Host,
		Insecure: false,
		Passw:    conf.Password,
	}

	fritz.PerformLogin()
	counter, err := fritz.ReadOnlineCounter()
	checkErr(err, "Cannot handle online counter data from Fritz.Box")

	writer.WriteData(counter)

	//cmd.Exec()
}

func loadConfig() Config {
	conf := Config{}
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
