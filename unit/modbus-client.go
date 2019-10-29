package unit

import (
	"github.com/goburrow/modbus"
)

type ModbusClient struct {
	SlaveId   uint8  `yaml:"slaveId"`
	Port      string `yaml:"port"`
	BoundRate int    `yaml:"boundRate"`
	DataBits  int    `yaml:"dataBits"`
	Parity    string `yaml:"parity"`
	StopBits  int    `yaml:"stopBits"`
	Timeout   string `yaml:"timeout"`
	Tests     Tests  `yaml:"tests"`
}

func (mc *ModbusClient) Run() error {
	handler := modbus.NewRTUClientHandler(mc.Port)
	handler.BaudRate = mc.BoundRate
	handler.DataBits = mc.DataBits
	handler.Parity = mc.Parity
	handler.StopBits = mc.StopBits
	handler.SlaveId = mc.SlaveId
	handler.Timeout = parseDuration(mc.Timeout)
	if err := handler.Connect(); err != nil {
		return err
	}
	defer handler.Close()
	//client := modbus.NewClient(handler)
	//reports := mc.Tests
	return nil
}
