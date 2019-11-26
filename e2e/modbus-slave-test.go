package e2e

import (
	"github.com/tbrandon/mbserver"
	"strconv"
	"strings"
)

type ModbusSlaveTest struct {
	Name     string              `yaml:"name"`
	Skip     string              `yaml:"skip"`
	Before   string              `yaml:"before"`
	Next     []string            `yaml:"next"`
	Lifetime int                 `yaml:"lifetime"`
	TimeOut  string              `yaml:"timeout"`
	AutoRun  string              `yaml:"autorun"`
	Function string              `yaml:"function"`
	Address  *uint16             `yaml:"address"`
	Quantity *uint16             `yaml:"quantity"`
	Data     []*Value            `yaml:"data"`
	Expected map[string][]*Value `yaml:"expected"`
	Write    map[string][]*Value `yaml:"write"`
	Success  string              `yaml:"success"`
	Error    string              `yaml:"error"`
	After    string              `yaml:"after"`
}

func (ms *ModbusSlaveTest) Check(f mbserver.Framer) bool {
	return false
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
