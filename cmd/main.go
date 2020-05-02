package main

import (
	"github.com/sirupsen/logrus"
	"github.com/toruta39/remofwd"
)

func main() {
	logrus.Info("Start probing")

	c := remofwd.NewClient()
	f := remofwd.NewForwarder()
	devices, err := c.FetchDevices()
	if err != nil {
		panic(err)
	}

	for _, d := range devices {
		_ = f.SendEvent(d.Name, d.Temperature())
		_ = f.SendEvent(d.Name, d.Motion())
		_ = f.SendEvent(d.Name, d.Humidity())
		_ = f.SendEvent(d.Name, d.Illumination())
	}

	logrus.Info("Finish probing")
}
