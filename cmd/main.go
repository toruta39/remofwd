package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/toruta39/remofwd"
)

var client remofwd.Client
var forwarder remofwd.Forwarder

func Probe() {
	logrus.Info("Start probing")

	devices, err := client.FetchDevices()
	if err != nil {
		panic(err)
	}

	for _, d := range devices {
		_ = forwarder.SendEvent(d.Name, d.Temperature())
		_ = forwarder.SendEvent(d.Name, d.Motion())
		_ = forwarder.SendEvent(d.Name, d.Humidity())
		_ = forwarder.SendEvent(d.Name, d.Illumination())
	}

	logrus.Info("Finish probing")
}

func main() {
	client = remofwd.NewClient()
	forwarder = remofwd.NewForwarder()

	for true {
		Probe()
		time.Sleep(30 * time.Second)
	}
}
