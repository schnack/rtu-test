package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"rtu-test/e2e"
)

func main() {

	fmt.Printf("Welcome to RTU-TEST v%s\n", Version)

	var comport = flag.String("p", "", "serial port")
	var filter = flag.String("f", "", "filter")
	var logs = flag.String("l", "", "log")
	var logLvl = flag.String("lvl", "", "logLvl")
	var help = flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	d := e2e.Init()
	fileNames := flag.Args()

	if len(fileNames) == 0 {
		fileNames = append(fileNames, "test.yml")
	}

	// Загружаем конфигурацию
	for _, fileName := range fileNames {
		if err := d.Load(fileName); err != nil {
			fmt.Printf("Loading configuration: %s\nError: %s", fileName, err)
			return
		} else {
			fmt.Printf("Loading configuration: %s\n", fileName)
		}
	}

	// Заменяем путь для вывода лога
	if *logs != "" {
		d.Log = *logs
	}

	// Заменяем уровень лога
	if *logLvl != "" {
		d.LogLvl = *logLvl
	}

	// Заменяем comport
	if *comport != "" {
		switch {
		case d.ModbusMaster != nil:
			d.ModbusMaster.Port = *comport
		case d.ModbusSlave != nil:
			d.ModbusSlave.Port = *comport
		case d.CustomSlave != nil:
			d.CustomSlave.Port = *comport
		}
	}

	// Заменяем фильтр
	if *filter != "" {
		if d.ModbusMaster != nil {
			d.ModbusMaster.Filter = *filter
		}
	}

	// Запускаем тесты
	d.RunTest(context.Background())
	//
	logrus.Exit(0)
}
