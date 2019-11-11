package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"rtu-test/unit"
)

func main() {
	var comport = flag.String("p", "", "serial port")
	var filter = flag.String("f", "", "filter")
	var logs = flag.String("l", "", "log")
	var logLvl = flag.String("lvl", "info", "logLvl")
	var help = flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	d := unit.Init()
	fileNames := flag.Args()

	if len(fileNames) == 0 {
		fileNames = append(fileNames, "test.yml")
	}

	for _, fileName := range fileNames {
		if err := d.Load(fileName); err != nil {
			logrus.Fatal(err)
		}
	}

	if *logs != "" {
		d.Log = *logs
	}

	if *logLvl != "" {
		d.LogLvl = *logLvl
	}

	if *comport != "" {
		d.ModbusClient.Port = *comport
	}

	if *filter != "" {
		d.ModbusClient.Filter = *filter
	}

	d.RunTest()
}
