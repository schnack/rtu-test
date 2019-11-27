package e2e

import (
	"encoding/binary"
	"github.com/goburrow/serial"
	"github.com/sirupsen/logrus"
	"github.com/tbrandon/mbserver"
	"sync"
	"time"
)

const (
	CoilsTable            = "coils"
	DiscreteInputTable    = "discreteInput"
	HoldingRegistersTable = "holdingRegisters"
	InputRegistersTable   = "inputRegisters"
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
	InputRegisters   []*Value `yaml:"inputRegisters"`

	Tests []*ModbusSlaveTest `yaml:"tests"`

	muCoils            sync.Mutex `yaml:"-"`
	muDiscreteInput    sync.Mutex `yaml:"-"`
	muHoldingRegisters sync.Mutex `yaml:"-"`
	muInputRegisters   sync.Mutex `yaml:"-"`

	currentTest *ModbusSlaveTest `yaml:"-"`
}

func (ms *ModbusSlave) getServer() *mbserver.Server {
	s := mbserver.NewServer()
	ms.Write1Bit(s.Coils, ms.Coils, &ms.muCoils)
	ms.Write1Bit(s.DiscreteInputs, ms.DiscreteInput, &ms.muDiscreteInput)
	ms.Write16Bit(s.HoldingRegisters, ms.HoldingRegisters, &ms.muHoldingRegisters)
	ms.Write16Bit(s.InputRegisters, ms.InputRegisters, &ms.muInputRegisters)
	s.RegisterFunctionHandler(1, ms.ActionHandler)
	s.RegisterFunctionHandler(2, ms.ActionHandler)
	s.RegisterFunctionHandler(3, ms.ActionHandler)
	s.RegisterFunctionHandler(4, ms.ActionHandler)
	s.RegisterFunctionHandler(5, ms.ActionHandler)
	s.RegisterFunctionHandler(6, ms.ActionHandler)
	s.RegisterFunctionHandler(15, ms.ActionHandler)
	s.RegisterFunctionHandler(16, ms.ActionHandler)
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

func (ms *ModbusSlave) ActionHandler(s *mbserver.Server, f mbserver.Framer) (result []byte, exp *mbserver.Exception) {

	var test *ModbusSlaveTest
	max := 0
	var next []string

	if ms.currentTest != nil && ms.currentTest.Next != nil {
		next = ms.currentTest.Next
	}

	for i := range ms.Tests {
		ball := ms.Tests[i].Check(f, next)

		if ball != 0 && ball > max {
			test = ms.Tests[i]
			max = ball
		}
	}

	if test != nil && test.Lifetime != nil {
		*test.Lifetime--
	}

	// Before
	if test != nil && test.BeforeWrite != nil {
		if v, ok := test.BeforeWrite[CoilsTable]; ok {
			ms.Write1Bit(s.Coils, v, &ms.muCoils)
		}
		if v, ok := test.BeforeWrite[DiscreteInputTable]; ok {
			ms.Write1Bit(s.DiscreteInputs, v, &ms.muDiscreteInput)
		}
		if v, ok := test.BeforeWrite[HoldingRegistersTable]; ok {
			ms.Write16Bit(s.HoldingRegisters, v, &ms.muHoldingRegisters)
		}
		if v, ok := test.BeforeWrite[InputRegistersTable]; ok {
			ms.Write16Bit(s.InputRegisters, v, &ms.muInputRegisters)
		}
	}

	switch ModbusFunction(f.GetFunction()) {
	case ReadCoils:
		result, exp = mbserver.ReadCoils(s, f)
	case ReadDiscreteInputs:
		result, exp = mbserver.ReadDiscreteInputs(s, f)
	case ReadHoldingRegisters:
		result, exp = mbserver.ReadHoldingRegisters(s, f)
	case ReadInputRegisters:
		result, exp = mbserver.ReadInputRegisters(s, f)
	case WriteSingleCoil:
		result, exp = mbserver.WriteSingleCoil(s, f)
	case WriteSingleRegister:
		result, exp = mbserver.WriteHoldingRegister(s, f)
	case WriteMultipleCoils:
		result, exp = mbserver.WriteMultipleCoils(s, f)
	case WriteMultipleRegisters:
		result, exp = mbserver.WriteHoldingRegisters(s, f)
	}

	// after
	if test != nil && test.AfterWrite != nil {
		if v, ok := test.AfterWrite[CoilsTable]; ok {
			ms.Write1Bit(s.Coils, v, &ms.muCoils)
		}
		if v, ok := test.AfterWrite[DiscreteInputTable]; ok {
			ms.Write1Bit(s.DiscreteInputs, v, &ms.muDiscreteInput)
		}
		if v, ok := test.AfterWrite[HoldingRegistersTable]; ok {
			ms.Write16Bit(s.HoldingRegisters, v, &ms.muHoldingRegisters)
		}
		if v, ok := test.AfterWrite[InputRegistersTable]; ok {
			ms.Write16Bit(s.InputRegisters, v, &ms.muInputRegisters)
		}
	}
	// TODO message
	// TODO write
	// TODO message
	ms.currentTest = test
	time.Sleep(parseDuration(ms.currentTest.TimeOut))
	return
}

func (ms *ModbusSlave) Write1Bit(s []byte, v []*Value, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
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
		for _, b := range data {
			if len(s) <= int(address) {
				logrus.Fatal("ModBus tables overflow")
			}
			if v[i].Type() == Bool {
				s[address] = b
				address++
			} else {
				for ii := 7; ii >= 0; ii-- {
					if len(s) <= int(address) {
						logrus.Fatal("ModBus tables overflow")
					}
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

func (ms *ModbusSlave) Write16Bit(s []uint16, v []*Value, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	var address uint16 = 0

	var vBytes uint16 = 0
	current := 0

	for i := range v {
		if v[i].Address != "" {
			// Сбрасываем счетчик бит
			if current != 0 {
				if len(s) <= int(address) {
					logrus.Fatal("ModBus tables overflow")
				}
				s[address] = vBytes
				address++
				vBytes = 0
				current = 0
			}

			rawAddress, err := parseStringByte(v[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			address = binary.BigEndian.Uint16(rawAddress)

		} else if current >= 16 {
			if len(s) <= int(address) {
				logrus.Fatal("ModBus tables overflow")
			}
			s[address] = vBytes
			address++
			vBytes = 0
			current = 0
		}

		switch v[i].Type() {
		case Bool:
			vBytes |= 1 << current
			current++
		default:
			data := v[i].Write()
			if current < 8 && current != 0 {
				current += 8 - (current % 8)
			}

			if current < 16 && current != 0 && !(len(data) == 1 && current == 8) {
				if len(s) <= int(address) {
					logrus.Fatal("ModBus tables overflow")
				}
				s[address] = vBytes
				address++
				vBytes = 0
				current = 0
			}

			for _, b := range data {
				if current >= 16 {
					if len(s) <= int(address) {
						logrus.Fatal("ModBus tables overflow")
					}
					s[address] = vBytes
					address++
					vBytes = 0
					current = 0
				}
				if current/8 == 0 {
					vBytes |= uint16(b) << 8
				} else {
					vBytes |= uint16(b)
				}
				current += 8
			}
		}
	}
	if current != 0 {
		if len(s) <= int(address) {
			logrus.Fatal("ModBus tables overflow")
		}
		s[address] = vBytes
	}
}
