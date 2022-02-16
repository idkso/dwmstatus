package main

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/idkso/dwmstatus/modules"
	_ "github.com/idkso/dwmstatus/plugins"
)

func main() {
	vals := make([]string, len(modules.Modules))

	for {
		for i, c := range modules.Modules {
			select {
			case m := <-c:
				vals[i] = m
			default:
				continue
			}
		}

		err := exec.Command("xsetroot", "-name", strings.Join(vals, " | ")).Run()
		if err != nil {
			log.Panic(err)
		}

		time.Sleep(time.Millisecond * 500)
	}
}
