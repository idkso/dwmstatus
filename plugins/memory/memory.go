package memory

import (
	"fmt"
	"log"
	"time"

	"github.com/idkso/dwmstatus/modules"
	"github.com/shirou/gopsutil/v3/mem"
)

func init() {
	c := make(chan string)
	go handleMemory(c)
	modules.Modules = append(modules.Modules, c)
}

func handleMemory(c chan string) {
	for {
		c <- getMemory()
		time.Sleep(time.Millisecond * 500)
	}
}

func getMemory() string {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Panic(err)
	}

	return fmt.Sprintf("RAM: %dMiB", v.Used/1024/1024)
}
