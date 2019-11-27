package e2e

import (
	"encoding/binary"
	"github.com/tbrandon/mbserver"
	"strconv"
	"strings"
)

const (
	FunctionAndAddressPoint = 1
	QuantityPoint           = 10
	DataPoint               = 100
	NextTestPoint           = 1000
)

type ModbusSlaveTest struct {
	Name        string              `yaml:"name"`
	Skip        string              `yaml:"skip"`
	Before      string              `yaml:"before"`
	Next        []string            `yaml:"next"`
	Lifetime    *int                `yaml:"lifetime"`
	TimeOut     string              `yaml:"timeout"`
	AutoRun     string              `yaml:"autorun"`
	Function    string              `yaml:"function"`
	Address     *uint16             `yaml:"address"`
	Quantity    *uint16             `yaml:"quantity"`
	Data        []*Value            `yaml:"data"`
	Expected    map[string][]*Value `yaml:"expected"`
	AfterWrite  map[string][]*Value `yaml:"afterWrite"`
	BeforeWrite map[string][]*Value `yaml:"beforeWrite"`
	Success     string              `yaml:"success"`
	Error       string              `yaml:"error"`
	After       string              `yaml:"after"`
}

// Для поиска нужного теста
func (ms *ModbusSlaveTest) Check(f mbserver.Framer, nexts []string) (points int) {

	if (ms.Lifetime != nil && *ms.Lifetime <= 0) || ms.Skip != "" {
		return
	}

	if ms.getFunction() == NilFunction && ms.getFunction() == ModbusFunction(f.GetFunction()) {
		return 0
	}

	if ms.Address == nil && *ms.Address != binary.BigEndian.Uint16(f.GetData()[0:2]) {
		return 0
	}

	switch ms.getFunction() {
	case ReadDiscreteInputs, ReadCoils, ReadInputRegisters, ReadHoldingRegisters:
		if ms.Quantity != nil {
			if *ms.Quantity == binary.BigEndian.Uint16(f.GetData()[2:4]) {
				points += DataPoint + QuantityPoint
			} else {
				return 0
			}
		}
		points += FunctionAndAddressPoint
	case WriteSingleCoil:
		if ms.Data != nil {
			countBit := 0
			var expected ReportExpected
			for _, v := range ms.Data {
				countBit, expected = v.Check(f.GetData()[2:4], 0, "", countBit)
				if !expected.Pass {
					return 0
				}
			}
			points += DataPoint
		}
		points += FunctionAndAddressPoint
	case WriteSingleRegister:
		if ms.Data != nil {
			countBit := 0
			var expected ReportExpected
			for _, v := range ms.Data {
				countBit, expected = v.Check(f.GetData()[2:4], 0, "", countBit)
				if !expected.Pass {
					return 0
				}
			}
			points += DataPoint
		}
		points += FunctionAndAddressPoint
	case WriteMultipleCoils, WriteMultipleRegisters:
		if ms.Quantity != nil {
			if *ms.Quantity == binary.BigEndian.Uint16(f.GetData()[2:4]) {
				points += QuantityPoint
			} else {
				return 0
			}
		}
		if ms.Data != nil {
			countBit := 0
			var expected ReportExpected
			for _, v := range ms.Data {
				countBit, expected = v.Check(f.GetData()[5:], 0, "", countBit)
				if !expected.Pass {
					return 0
				}
			}
			points += DataPoint
		}

		points += FunctionAndAddressPoint
	}

	if nexts != nil {
		for _, nextName := range nexts {
			if ms.Name == nextName {
				points += NextTestPoint
				break
			}
		}
	}

	return
}

func (ms *ModbusSlaveTest) getFunction() ModbusFunction {
	mFunc := strings.ReplaceAll(strings.ToLower(ms.Function), " ", "")
	if strings.HasPrefix(mFunc, "0x") {
		if a, err := strconv.ParseInt(strings.TrimPrefix(mFunc, "0x"), 16, 8); err == nil {
			mFunc = strconv.Itoa(int(a))
		}
	}
	switch mFunc {
	case "readdiscreteinputs", "2":
		return ReadDiscreteInputs
	case "readcoils", "1":
		return ReadCoils
	case "writesinglecoil", "5":
		return WriteSingleCoil
	case "writemultiplecoils", "15":
		return WriteMultipleCoils
	case "readinputregisters", "4":
		return ReadInputRegisters
	case "readholdingregisters", "3":
		return ReadHoldingRegisters
	case "writesingleregister", "6":
		return WriteSingleRegister
	case "writemultipleregisters", "16":
		return WriteMultipleRegisters
	default:
		return NilFunction
	}
}
