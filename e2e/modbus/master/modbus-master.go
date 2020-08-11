package master

import (
	"fmt"
	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
	"log"
	"rtu-test/e2e/common"
	"rtu-test/e2e/template"
	"strings"
)

type ModbusMaster struct {
	SlaveId   uint8                          `yaml:"slaveId"`
	Port      string                         `yaml:"port"`
	BoundRate int                            `yaml:"boundRate"`
	DataBits  int                            `yaml:"dataBits"`
	Parity    string                         `yaml:"parity"`
	StopBits  int                            `yaml:"stopBits"`
	Timeout   string                         `yaml:"timeout"`
	Filter    string                         `yaml:"filter"`
	Tests     map[string][]*ModbusMasterTest `yaml:"tests"`
}

type loger struct {
}

func (l *loger) Write(p []byte) (n int, err error) {
	logrus.Debug(strings.TrimPrefix(string(p), "modbus: "))
	return len(p), nil
}

// TODO test
func (mc *ModbusMaster) getHandler() *modbus.RTUClientHandler {
	handler := modbus.NewRTUClientHandler(mc.Port)
	handler.BaudRate = mc.BoundRate
	handler.DataBits = mc.DataBits
	handler.Parity = mc.Parity
	handler.StopBits = mc.StopBits
	handler.SlaveId = mc.SlaveId
	handler.Timeout = common.ParseDuration(mc.Timeout)
	handler.IdleTimeout = common.ParseDuration(mc.Timeout)
	handler.Logger = log.New(&loger{}, "", 0)
	return handler
}

// TODO Test
func (mc *ModbusMaster) Run(reports *ReportGroups) error {
	handler := mc.getHandler()
	if err := handler.Connect(); err != nil {
		return fmt.Errorf("open %s: %s", handler.Address, err)
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
		logrus.Warnf(common.Render(template.TestMasterModBusGROUP, report))
		for _, test := range tests {
			if filterTest != "" && filterTest != "all" && filterTest != test.Name {
				continue
			}
			// Подменяем адрес, если он переопределен в тесте
			// Необходимо для управления несколькими устройствами на 1 шине
			if test.SlaveId != 0 {
				handler.SlaveId = test.SlaveId
			}
			report.Tests = append(report.Tests, test.Run(client))
			// Возвращаем адрес по умолчанию
			handler.SlaveId = mc.SlaveId

			// При необходимости закрываем порт
			if test.Disconnect {
				handler.Close()
			}
		}
		reports.ReportGroup = append(reports.ReportGroup, report)
	}

	return nil
}
