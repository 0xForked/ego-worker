package main

import (
	"github.com/aasumitro/ego-worker/helper"
	"github.com/aasumitro/ego-worker/subscriber/service"
)

func main() {
	config := helper.GetConfig()
	app := service.Messaging{}
	app.SubscribeMessage(config)
}
