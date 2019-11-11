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
	Filter    string                   `yaml:"filter"`
	Tests     map[string][]*ModbusTest `yaml:"tests"`
}

type loger struct {
}

func (l *loger) Write(p []byte) (n int, err error) {
	logrus.Debug(strings.TrimPrefix(string(p), "modbus: "))
	return len(p), nil
}

func (mc *ModbusClient) getHandler() *modbus.RTUClientHandler {
	handler := modbus.NewRTUClientHandler(mc.Port)
	handler.BaudRate = mc.BoundRate
	handler.DataBits = mc.DataBits
	handler.Parity = mc.Parity
	handler.StopBits = mc.StopBits
	handler.SlaveId = mc.SlaveId
	handler.Timeout = parseDuration(mc.Timeout)
	handler.Logger = log.New(&loger{}, "", 0)
	return handler
}

func (mc *ModbusClient) Run() error {
	handler := mc.getHandler()
	if err := handler.Connect(); err != nil {
		return err
	}
	defer handler.Close()
	client := modbus.NewClient(handler)

	filterGroup := ""
	filterTest := ""
	filter := strings.Split(mc.Filter, ":")
	if len(filter) > 1 {
		filterGroup = filter[0]
		filterTest = filter[1]
	} else {
		filterGroup = filter[0]
	}

	for group, tests := range mc.Tests {
		if filterGroup != "" && filterGroup != "all" && filterGroup != group {
			continue
		}
		report := ReportGroup{Name: group}
		logrus.Warnf(render(TestGROUP, report))
		for _, test := range tests {
			if filterTest != "" && filterTest != "all" && filterTest != test.Name {
				continue
			}
			report.Tests = append(report.Tests, test.Run(client))
		}
	}

	return nil
}
