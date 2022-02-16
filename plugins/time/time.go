package time

import (
	"time"

	"github.com/idkso/dwmstatus/modules"
)

func init() {
	c := make(chan string)
	go handleTime(c)
	modules.Modules = append(modules.Modules, c)
}

func handleTime(c chan string) {
	for {
		c <- getTime()
		time.Sleep(time.Second)
	}
}

func getTime() string {
	t := time.Now()
	return t.Format("3:04:05 PM")
}
