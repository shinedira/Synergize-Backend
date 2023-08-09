package main

import (
	"runtime"
	"synergize/backend-test/bootstrap"
	"synergize/backend-test/pkg/facades"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	bootstrap.Boot()

	go facades.Route.Start(facades.Config.GetString("app.host"))

	select {}
}
