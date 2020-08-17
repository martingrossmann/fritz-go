package main

import (
	"fmt"
	_ "github.com/martingrossmann/fritz-go/cmd"
	"github.com/martingrossmann/fritz-go/fritz"
	"log"

	"github.com/magiconair/properties"
)

type Config struct {
	Host     string `properties:"host,default=http://fritz.box"`
	Password string `properties:"password"`
}

func main() {
	conf := loadConfig()

	fritz := &fritz.FritzBox{
		Host:     conf.Host,
		Insecure: false,
		Passw:    conf.Password,
	}

	fritz.PerformLogin()
	counter, err := fritz.ReadOnlineCounter()
	if err != nil {
		log.Fatal("Cannot handle online counter data from Fritz.Box", err)
	}

	fmt.Println(counter)

	//cmd.Exec()
}

func loadConfig() Config {
	conf := Config{}
	p := properties.MustLoadFile("settings.conf", properties.UTF8)
	err := p.Decode(&conf)
	if err != nil {
		log.Fatal("Error at parsing settings.conf: ", err)
	}
	return conf
}
