package cpu

import (
	"fmt"
	"log"
	"time"

	"github.com/idkso/dwmstatus/modules"
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

	var total float64
	for _, p := range cper {
		total += p
	}
        perc := "%"
        
        if int(total)/len(cper) < 9 {
          perc = perc + " "
        }

	return fmt.Sprintf("CPU: %d", int(total)/len(cper)) + perc
}
