package fritz

import (
	"io/ioutil"
	"log"
	"net/http"
)

var fritzUri = "http://fritz.box"

func Connect() {

	log.Print("Connect to ", fritzUri)
	resp, err := http.Get(fritzUri + "/login_sid.lua")
	//resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
