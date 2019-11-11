package unit

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/sirupsen/logrus"
	"os"
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
	Version      string       `yaml:"version"`
	Name         string       `yaml:"name"`
	Log          string       `yaml:"log"`
	LogLvl       string       `yaml:"logLvl"`
	Description  string       `yaml:"description"`
	ModbusClient ModbusClient `yaml:"modbusClient"`
}

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

func (d *Device) RunTest() {
	logrus.SetFormatter(&logrus.TextFormatter{})

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
		logrus.SetOutput(os.Stdout)
	case LogStderr:
		logrus.SetOutput(os.Stderr)
	default:
		file, err := os.OpenFile(d.Log, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			logrus.Fatal(err)
		}
		defer file.Close()
		logrus.SetOutput(file)
	}

	if err := d.ModbusClient.Run(); err != nil {
		logrus.Fatal(err)
	}
}
