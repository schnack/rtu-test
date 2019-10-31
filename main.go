package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"rtu-test/unit"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	device := unit.Device{}
	if err := device.Load(); err != nil {
		logrus.Fatal(err)
	}
	if err := device.ModbusClient.Run(); err != nil {
		logrus.Fatal(err)
	}
}
