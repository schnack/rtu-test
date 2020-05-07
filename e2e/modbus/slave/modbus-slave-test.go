package slave

import (
	"encoding/binary"
	"github.com/schnack/mbslave"
	"rtu-test/e2e/common"
	"rtu-test/e2e/modbus/master"
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
	Name        string                     `yaml:"name"`
	Skip        string                     `yaml:"skip"`
	Fatal       string                     `yaml:"fatal"`
	Before      Message                    `yaml:"before"`
	Next        []string                   `yaml:"next"`
	Lifetime    *int                       `yaml:"lifetime"`
	TimeOut     string                     `yaml:"timeout"`
	AutoRun     string                     `yaml:"autorun"`
	Function    string                     `yaml:"function"`
	Address     *uint16                    `yaml:"address"`
	Quantity    *uint16                    `yaml:"quantity"`
	Data        []*common.Value            `yaml:"data"`
	Expected    map[string][]*common.Value `yaml:"expected"`
	AfterWrite  map[string][]*common.Value `yaml:"afterWrite"`
	BeforeWrite map[string][]*common.Value `yaml:"beforeWrite"`
	Success     Message                    `yaml:"success"`
	Error       Message                    `yaml:"error"`
	After       Message                    `yaml:"after"`
}

// Для поиска нужного теста
func (ms *ModbusSlaveTest) Check(request mbslave.Request, nexts []string) (points int) {

	if (ms.Lifetime != nil && *ms.Lifetime <= 0) || ms.Skip != "" {
		return
	}

	if ms.getFunction() == master.NilFunction || ms.getFunction() != master.ModbusFunction(request.GetFunction()) {
		return 0
	}

	if ms.Address == nil || *ms.Address != request.GetAddress() {
		return 0
	}

	switch ms.getFunction() {
	case master.ReadDiscreteInputs, master.ReadCoils, master.ReadInputRegisters, master.ReadHoldingRegisters:
		if ms.Quantity != nil {
			if *ms.Quantity == request.GetQuantity() {
				points += DataPoint + QuantityPoint
			} else {
				return 0
			}
		}
		points += FunctionAndAddressPoint
	case master.WriteSingleCoil:
		if ms.Data != nil {
			countBit := 0
			var expected common.ReportExpected
			for _, v := range ms.Data {
				countBit, expected = v.Check(request.GetData(), 0, "", countBit, 8, binary.BigEndian)
				if !expected.Pass {
					return 0
				}
			}
			points += DataPoint
		}
		points += FunctionAndAddressPoint
	case master.WriteSingleRegister:
		if ms.Data != nil {
			countBit := 0
			var expected common.ReportExpected
			for _, v := range ms.Data {
				countBit, expected = v.Check(request.GetData(), 0, "", countBit, 16, binary.BigEndian)
				if !expected.Pass {
					return 0
				}
			}
			points += DataPoint
		}
		points += FunctionAndAddressPoint
	case master.WriteMultipleCoils, master.WriteMultipleRegisters:
		if ms.Quantity != nil {
			if *ms.Quantity == request.GetQuantity() {
				points += QuantityPoint
			} else {
				return 0
			}
		}
		if ms.Data != nil {
			countBit := 0
			var expected common.ReportExpected
			for _, v := range ms.Data {
				bitSize := 8
				if ms.getFunction() == master.WriteMultipleRegisters {
					bitSize = 16
				}
				countBit, expected = v.Check(request.GetData(), 0, "", countBit, bitSize, binary.BigEndian)
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

func (ms *ModbusSlaveTest) getFunction() master.ModbusFunction {
	mFunc := strings.ReplaceAll(strings.ToLower(ms.Function), " ", "")
	if strings.HasPrefix(mFunc, "0x") {
		if a, err := strconv.ParseInt(strings.TrimPrefix(mFunc, "0x"), 16, 8); err == nil {
			mFunc = strconv.Itoa(int(a))
		}
	}
	switch mFunc {
	case "readdiscreteinputs", "2":
		return master.ReadDiscreteInputs
	case "readcoils", "1":
		return master.ReadCoils
	case "writesinglecoil", "5":
		return master.WriteSingleCoil
	case "writemultiplecoils", "15":
		return master.WriteMultipleCoils
	case "readinputregisters", "4":
		return master.ReadInputRegisters
	case "readholdingregisters", "3":
		return master.ReadHoldingRegisters
	case "writesingleregister", "6":
		return master.WriteSingleRegister
	case "writemultipleregisters", "16":
		return master.WriteMultipleRegisters
	default:
		return master.NilFunction
	}
}
