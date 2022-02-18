package wttr

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/idkso/dwmstatus/modules"
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
	request, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "https://wttr.in/?format=%c%h+%t+%w+%m", nil)
	if err != nil {
		log.Panic(err)
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return "Weather: " + string(body)
}
