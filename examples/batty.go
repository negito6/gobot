package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
)

func main() {
	gbot := gobot.NewGobot()

	api.NewAPI(gbot).Start()

	gbot.AddCommand("echo", func(params map[string]interface{}) interface{} {
		return params["a"]
	})

	loopback := gobot.NewLoopbackAdaptor("loopback")
	ping := gobot.NewPingDriver(loopback, "ping")
	r := gobot.NewRobot("TestBot",
		[]gobot.Connection{loopback},
		[]gobot.Device{ping},
	)

	r.AddCommand("hello", func(params map[string]interface{}) interface{} {
		return fmt.Sprintf("Hello, %v!", params["greeting"])
	})

	gbot.AddRobot(r)
	gbot.Start()
}
