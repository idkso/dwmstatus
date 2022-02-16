package date

import (
	"time"

	"github.com/idkso/dwmstatus/modules"
)

func init() {
	c := make(chan string)
	go handleDate(c)
	modules.Modules = append(modules.Modules, c)
}

func handleDate(c chan string) {
	for {
		c <- getDate()
		time.Sleep(time.Minute)
	}
}

func getDate() string {
	t := time.Now()
	return t.Format("Mon 01-02-06")
}
