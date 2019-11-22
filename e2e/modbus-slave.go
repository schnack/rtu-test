package e2e

import (
	"encoding/binary"
	"github.com/goburrow/serial"
	"github.com/sirupsen/logrus"
	"github.com/tbrandon/mbserver"
)

type ModbusSlave struct {
	SlaveId   uint8  `yaml:"slaveId"`
	Port      string `yaml:"port"`
	BoundRate int    `yaml:"boundRate"`
	DataBits  int    `yaml:"dataBits"`
	Parity    string `yaml:"parity"`
	StopBits  int    `yaml:"stopBits"`
	Timeout   string `yaml:"timeout"`

	Coils            []*Value `yaml:"coils"`
	DiscreteInput    []*Value `yaml:"discreteInput"`
	HoldingRegisters []*Value `yaml:"holdingRegisters"`
	InputRegisters   []*Value `yaml:"InputRegisters"`

	Tests []*ModbusSlaveTest `yaml:"tests"`
}

func (ms *ModbusSlave) Write1Bit(s []byte, v []*Value) {
	var address uint16 = 0
	for i := range v {
		if v[i].Address != "" {
			rawAddress, err := parseStringByte(v[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			address = binary.BigEndian.Uint16(rawAddress)
		}

		data := v[i].Write()
		for b := range data {
			if len(s) < int(address) {
				logrus.Fatal("ModBus tables overflow")
			}

			if v[i].Type() == Bool {
				s[address] = byte(b)
				address++
			} else {
				for ii := 0; ii <= 7; ii++ {
					if (b & (1 << ii)) != 0 {
						s[address] = 1
					} else {
						s[address] = 0
					}
					address++
				}
			}
		}
	}
}

func (ms *ModbusSlave) getServer() *mbserver.Server {
	s := mbserver.NewServer()
	ms.Write1Bit(s.Coils, ms.Coils)
	ms.Write1Bit(s.DiscreteInputs, ms.Coils)
	s.RegisterFunctionHandler(1, ms.ReadCoils)
	s.RegisterFunctionHandler(2, ms.ReadDiscreteInputs)
	s.RegisterFunctionHandler(3, ms.ReadHoldingRegisters)
	s.RegisterFunctionHandler(4, ms.ReadInputRegisters)
	s.RegisterFunctionHandler(5, ms.WriteSingleCoil)
	s.RegisterFunctionHandler(6, ms.WriteHoldingRegister)
	s.RegisterFunctionHandler(15, ms.WriteMultipleCoils)
	s.RegisterFunctionHandler(16, ms.WriteHoldingRegisters)
	return s
}

func (ms *ModbusSlave) Run() {
	s := ms.getServer()

	err := s.ListenRTU(&serial.Config{
		Address:  ms.Port,
		BaudRate: ms.BoundRate,
		DataBits: ms.DataBits,
		StopBits: ms.StopBits,
		Parity:   ms.Parity,
		Timeout:  parseDuration(ms.Timeout),
	})
	if err != nil {
		logrus.Fatalf("failed to listen, got %v\n", err)
	}
}

func (ms *ModbusSlave) ReadCoils(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.ReadCoils(s, f)
}

func (ms *ModbusSlave) ReadDiscreteInputs(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.ReadDiscreteInputs(s, f)
}

func (ms *ModbusSlave) ReadHoldingRegisters(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.ReadHoldingRegisters(s, f)
}

func (ms *ModbusSlave) ReadInputRegisters(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.ReadInputRegisters(s, f)
}

func (ms *ModbusSlave) WriteSingleCoil(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.WriteSingleCoil(s, f)
}

func (ms *ModbusSlave) WriteHoldingRegister(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.WriteHoldingRegister(s, f)
}

func (ms *ModbusSlave) WriteMultipleCoils(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.WriteMultipleCoils(s, f)
}

func (ms *ModbusSlave) WriteHoldingRegisters(s *mbserver.Server, f mbserver.Framer) ([]byte, *mbserver.Exception) {
	return mbserver.WriteHoldingRegisters(s, f)
}
