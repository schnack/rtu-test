package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"rtu-test/e2e"
)

func main() {
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
			logrus.Fatal(err)
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
		if d.ModbusMaster != nil {
			d.ModbusMaster.Port = *comport
		} else if d.ModbusSlave != nil {
			d.ModbusSlave.Port = *comport
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
