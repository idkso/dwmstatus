package cpu

import (
	"fmt"
	"log"
	"time"

	"dwmstatus/modules"
	"github.com/shirou/gopsutil/v3/cpu"
)

func init() {
	c := make(chan string)
	go handleCPU(c)
	modules.Modules = append(modules.Modules, c)
}

func handleCPU(c chan string) {
	for {
		c <- getCPU()
		time.Sleep(time.Millisecond * 500)
	}
}

func getCPU() string {
	cper, err := cpu.Percent(0, true)
	if err != nil {
		log.Panic(err)
	}
	var total float64 = 0
	for _, p := range cper {
		total += p
	}
	usage := int(total)/len(cper)
	
	idk := fmt.Sprintf("CPU: %d", int(total)/len(cper))

	if usage > 9 {
		return idk + "%"
	} else {
		return idk + "% "
	}
}
