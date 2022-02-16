package uptime

import (
	"log"
	"time"

	"github.com/idkso/dwmstatus/modules"
	"github.com/shirou/gopsutil/v3/host"
)

func init() {
	c := make(chan string)
	go handleUptime(c)
	modules.Modules = append(modules.Modules, c)
}

func handleUptime(c chan string) {
	for {
		c <- getUptime()

		time.Sleep(time.Second)
	}
}

func getUptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		log.Panic(err)
	}

	return time.Unix(int64(uptime), 0).Format("Up: 02d 15h 05m")
}
