package battery

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/idkso/dwmstatus/modules"
)

func init() {
	c := make(chan string)
	go handleBattery(c)
	modules.Modules = append(modules.Modules, c)
}

func handleBattery(c chan string) {
	for {
		c <- getBattery()
		time.Sleep(time.Second)
	}
}

func getBattery() string {
	f, err := os.Open("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		log.Panic(err)
	}
	r := bufio.NewReader(f)
	s, err := r.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	return strings.TrimSuffix(s, "\n") + "%"
}
