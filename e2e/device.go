package e2e

import (
	"context"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"sync"
)

var instanceDevice *Device
var onceDevice sync.Once

const (
	LogStdout = "stdout"
	LogStderr = "stderr"
)

func Init() *Device {
	onceDevice.Do(func() {
		instanceDevice = &Device{}
	})
	return instanceDevice
}

type Device struct {
	Version      string        `yaml:"version"`
	Name         string        `yaml:"name"`
	Log          string        `yaml:"log"`
	LogLvl       string        `yaml:"logLvl"`
	Description  string        `yaml:"description"`
	ExitMessage  Message       `yaml:"exitMessage"`
	ModbusMaster *ModbusMaster `yaml:"modbusMaster"`
	ModbusSlave  *ModbusSlave  `yaml:"modbusSlave"`
	Slave        *Slave        `yaml:"slave"`
	Master       *Master       `yaml:"master"`
}

// Load - загружает конфигурацию лога
func (d *Device) Load(s string) error {
	file, err := os.Open(s)
	if err != nil {
		return fmt.Errorf("device load config: %s", err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(d); err != nil {
		return fmt.Errorf("parse yaml error: %s", err)
	}
	return nil
}

func (d *Device) RunTest(ctx context.Context) {
	format := &logrus.TextFormatter{}

	logrus.SetFormatter(format)
	// Настраиваем уровень логирования
	switch d.LogLvl {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}

	switch d.Log {
	case "", "off":
		logrus.SetOutput(os.Stderr)
		logrus.SetLevel(logrus.PanicLevel)
	case LogStdout:
		if runtime.GOOS == "windows" {
			format.ForceColors = true
			logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
		} else {
			logrus.SetOutput(os.Stdout)
		}
	case LogStderr:
		if runtime.GOOS == "windows" {
			format.ForceColors = true
			logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stderr))
		} else {
			logrus.SetOutput(os.Stderr)
		}
	default:
		file, err := os.OpenFile(d.Log, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			logrus.Fatal(err)
		}
		defer file.Close()
		logrus.SetOutput(file)
	}

	switch {
	case d.ModbusMaster != nil:
		report := ReportGroups{
			Name:        d.Name,
			Description: d.Description,
		}
		// Вывод отчета в конце выполнения программы
		logrus.RegisterExitHandler(func() { d.ExitMessage.PrintReportMasterGroups(report) })

		// Запуск тестов
		if err := d.ModbusMaster.Run(&report); err != nil {
			logrus.Fatalf("Exit app modbus master: %s", err)
		}
	case d.ModbusSlave != nil:
		if err := d.ModbusSlave.Run(); err != nil {
			logrus.Fatalf("Exit app modbus slave: %s", err)
		}
	case d.Slave != nil:
		if err := d.Slave.Run(); err != nil {
			logrus.Fatalf("Exit app slave: %s", err)
		}
	case d.Master != nil:
		// TODO
	default:
		logrus.Fatal("configuration file not found")
	}

}
