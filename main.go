package main

import (
	"fmt"
	_ "github.com/martingrossmann/fritz-go/cmd"
	"github.com/martingrossmann/fritz-go/fritz"
	"log"
)

func main() {
	fritz := &fritz.FritzBox{
		Host:     "http://fritz.box",
		Insecure: false,
	}

	fritz.PerformLogin("...")
	counter, err := fritz.ReadOnlineCounter()
	if err != nil {
		log.Fatal("Cannot handle online counter data from Fritz.Box", err)
	}

	fmt.Println(counter)

	//cmd.Exec()
}
