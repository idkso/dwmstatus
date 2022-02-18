package wttr

import (
	"dwmstatus/modules"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init() {
	c := make(chan string)
	go handleWTTR(c)
	modules.Modules = append(modules.Modules, c)
}

func handleWTTR(c chan string) {
	for {
		c <- getWTTR()
		time.Sleep(time.Hour)
	}
}

func getWTTR() string {
	resp, err := http.Get("https://wttr.in/?format=%c%h+%t+%w+%m")
	if err != nil {
		log.Panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return "Weather: " + string(body)
}
