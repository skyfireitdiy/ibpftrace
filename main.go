package main

import (
	"log"
	"runtime"

	"go.skyfire.com/gui"
	"go.skyfire.com/utils"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			traceback := make([]byte, 10240)
			runtime.Stack(traceback, true)
			log.Println(string(traceback))
		}
	}()
	if !utils.EnvCheck() {
		log.Fatal("Environment check failed")
	}

	gui.InitGuiData()

	gui.GuiLoop()
}
