package uptime

import (
	"fmt"
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
	u, err := host.Uptime()
	if err != nil {
		log.Panic(err)
	}

	days := 0
	hours := 0
	mins := u / 60

	for mins >= 60 {
		mins -= 60
		hours += 1
	}

	for hours >= 24 {
		hours -= 24
		days += 1
	}
	uptimetxt := ""

	if days > 0 && hours < 1 {
		uptimetxt = fmt.Sprintf("Up: %dd %dm", days, mins)
	} else if days > 0 {
		uptimetxt = fmt.Sprintf("Up: %dd %dh %dm", days, hours, mins)
	} else if hours > 0 {
		uptimetxt = fmt.Sprintf("Up: %dh %dm", hours, mins)
	} else {
		uptimetxt = fmt.Sprintf("Up: %dm", mins)
	}

	return uptimetxt
}
