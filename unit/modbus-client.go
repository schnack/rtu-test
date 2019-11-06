package unit

import (
	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

type ModbusClient struct {
	SlaveId   uint8                    `yaml:"slaveId"`
	Port      string                   `yaml:"port"`
	BoundRate int                      `yaml:"boundRate"`
	DataBits  int                      `yaml:"dataBits"`
	Parity    string                   `yaml:"parity"`
	StopBits  int                      `yaml:"stopBits"`
	Timeout   string                   `yaml:"timeout"`
	Tests     map[string][]*ModbusTest `yaml:"tests"`
}

type loger struct {
}

func (l *loger) Write(p []byte) (n int, err error) {
	logrus.Debug(strings.TrimPrefix(string(p), "modbus: "))
	return len(p), nil
}

func (mc *ModbusClient) Run() error {
	handler := modbus.NewRTUClientHandler(mc.Port)
	handler.BaudRate = mc.BoundRate
	handler.DataBits = mc.DataBits
	handler.Parity = mc.Parity
	handler.StopBits = mc.StopBits
	handler.SlaveId = mc.SlaveId
	handler.Timeout = parseDuration(mc.Timeout)
	handler.Logger = log.New(&loger{}, "", 0)
	if err := handler.Connect(); err != nil {
		return err
	}
	defer handler.Close()
	client := modbus.NewClient(handler)

	for group, tests := range mc.Tests {
		logrus.Warnf(">>> GROUP    %s", group)
		for _, test := range tests {
			logrus.Warnf("=== RUN      %s", test.Name)
			test.Before.Print()
			test.Exec(client)
			if test.Check() {
				logrus.Warnf("--- PASS:    %s (%s)", test.Name, test.ResultTime)
				test.Success.Print()
			} else {
				logrus.Errorf("--- FAIL:    %s (%s)%s", test.Name, test.ResultTime, test.String())
				test.Error.Print()
			}
			test.After.Print()
		}
	}

	return nil
}
