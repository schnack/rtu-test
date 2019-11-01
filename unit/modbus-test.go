package unit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/goburrow/modbus"
	"strconv"
	"strings"
	"time"
)

type ModbusFunction int

const (
	NilFunction            = ModbusFunction(0)
	ReadCoils              = ModbusFunction(modbus.FuncCodeReadCoils)
	ReadDiscreteInputs     = ModbusFunction(modbus.FuncCodeReadDiscreteInputs)
	ReadHoldingRegisters   = ModbusFunction(modbus.FuncCodeReadHoldingRegisters)
	ReadInputRegisters     = ModbusFunction(modbus.FuncCodeReadInputRegisters)
	WriteSingleCoil        = ModbusFunction(modbus.FuncCodeWriteSingleCoil)
	WriteSingleRegister    = ModbusFunction(modbus.FuncCodeWriteSingleRegister)
	WriteMultipleCoils     = ModbusFunction(modbus.FuncCodeWriteMultipleCoils)
	WriteMultipleRegisters = ModbusFunction(modbus.FuncCodeWriteMultipleRegisters)
)

type ModbusTest struct {
	Name          string        `yaml:"name"`
	Before        Message       `yaml:"before"`
	Function      string        `yaml:"function"`
	Address       *uint16       `yaml:"address"`
	Quantity      *uint16       `yaml:"quantity"`
	Write         []Value       `yaml:"write"`
	Expected      []Value       `yaml:"expected"`
	ExpectedError string        `yaml:"expectedError"`
	ExpectedTime  string        `yaml:"expectedTime"`
	Success       Message       `yaml:"success"`
	Error         Message       `yaml:"error"`
	After         Message       `yaml:"after"`
	ResultByte    []byte        `yaml:"-"`
	ResultTime    time.Duration `yaml:"-"`
	ResultError   error         `yaml:"-"`
}

func (mt *ModbusTest) String() string {
	buff := bytes.NewBufferString("\n")
	if !mt.CheckError() {
		buff.WriteString("Error:\n")
		buff.WriteString(fmt.Sprintf("\texpected: %s\n", mt.StringErrorExpected()))
		buff.WriteString(fmt.Sprintf("\t     got: %s\n", mt.StringErrorGot()))
	}

	if !mt.CheckDuration() {
		buff.WriteString("Time:\n")
		buff.WriteString(fmt.Sprintf("\texpected: %s\n", mt.StringTimeExpected()))
		buff.WriteString(fmt.Sprintf("\t     got: %s\n", mt.StringTimeGot()))
	}

	if !mt.CheckData() {
		buff.WriteString("Data\n")
		switch mt.getFunction() {
		case ReadDiscreteInputs, ReadCoils, ReadHoldingRegisters, ReadInputRegisters:
			for _, v := range mt.Expected {
				if v.Pass {
					continue
				}
				buff.WriteString(fmt.Sprintf("%s:\n", v.Name))
				buff.WriteString(fmt.Sprintf("\texpected: %s\n", v.StringExpected()))
				buff.WriteString(fmt.Sprintf("\t     got: %s\n", v.StringGot()))
			}
		case WriteSingleCoil:
			buff.WriteString(fmt.Sprintf("\texpected: 0x%04x\n", dataSingleCoil(mt.getWriteData())))
			if mt.ResultByte == nil {
				buff.WriteString("\t     got: \n")
			} else {
				buff.WriteString(fmt.Sprintf("\t     got: 0x%04x\n", mt.ResultByte))
			}

		case WriteSingleRegister:
			buff.WriteString(fmt.Sprintf("\texpected: 0x%04x\n", mt.getWriteData()[:2]))
			if mt.ResultByte == nil {
				buff.WriteString("\t     got: \n")
			} else {
				buff.WriteString(fmt.Sprintf("\t     got: 0x%04x\n", mt.ResultByte[:2]))
			}
		case WriteMultipleRegisters, WriteMultipleCoils:
			buff.WriteString(fmt.Sprintf("\texpected: 0x%04x\n", mt.getQuantity()))
			if mt.ResultByte == nil {
				buff.WriteString("\t     got: \n")
			} else {
				buff.WriteString(fmt.Sprintf("\t     got: 0x%04x\n", binary.BigEndian.Uint16(mt.ResultByte)))
			}

		}
	}
	return buff.String()
}

const FormatDuration = "%[1]s"
const FormatError = "%[1]s"

func (mt *ModbusTest) StringTimeExpected() string {
	return fmt.Sprintf(FormatDuration, parseDuration(mt.ExpectedTime).String())
}

func (mt *ModbusTest) StringTimeGot() string {
	return fmt.Sprintf(FormatDuration, mt.ResultTime.String())
}

func (mt *ModbusTest) StringErrorExpected() string {
	return fmt.Sprintf(FormatError, mt.getError())
}

func (mt *ModbusTest) StringErrorGot() string {
	return fmt.Sprintf(FormatError, mt.ResultError.Error())
}

func (mt *ModbusTest) Check() bool {
	if mt.CheckError() && mt.CheckData() && mt.CheckDuration() {
		return true
	}
	return false
}

func (mt *ModbusTest) CheckData() bool {
	switch mt.getFunction() {
	case ReadCoils, ReadDiscreteInputs, ReadHoldingRegisters, ReadInputRegisters:
		countBit := 0
		pass := true
		for _, v := range mt.Expected {
			countBit = v.Check(mt.ResultByte, countBit)
			if !v.Pass {
				pass = false
			}
		}
		return pass
	case WriteSingleCoil:
		if !byteToEq(dataSingleCoil(mt.getWriteData()), mt.ResultByte) {
			return false
		}
	case WriteSingleRegister:
		if !byteToEq(mt.getWriteData()[:2], mt.ResultByte) {
			return false
		}
	case WriteMultipleCoils, WriteMultipleRegisters:
		if mt.ResultByte == nil {
			return false
		}
		if mt.getQuantity() != binary.BigEndian.Uint16(mt.ResultByte) {
			return false
		}
	}
	return true
}

func (mt *ModbusTest) CheckDuration() bool {
	if mt.ResultTime > parseDuration(mt.ExpectedTime) {
		return false
	}
	return true
}

func (mt *ModbusTest) CheckError() bool {
	var got string
	if mt.ResultError != nil {
		got = mt.ResultError.Error()
	}
	return mt.getError() == got
}

func (mt *ModbusTest) Exec(client modbus.Client) {
	switch mt.getFunction() {
	case ReadDiscreteInputs:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.ReadDiscreteInputs(*mt.Address, mt.getQuantity())
		mt.ResultTime = time.Since(startTime)
	case ReadCoils:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.ReadCoils(*mt.Address, mt.getQuantity())
		mt.ResultTime = time.Since(startTime)
	case WriteSingleCoil:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.WriteSingleCoil(*mt.Address, binary.BigEndian.Uint16(dataSingleCoil(mt.getWriteData())))
		mt.ResultTime = time.Since(startTime)
	case WriteMultipleCoils:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.WriteMultipleCoils(*mt.Address, mt.getQuantity(), mt.getWriteData())
		mt.ResultTime = time.Since(startTime)
	case ReadInputRegisters:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.ReadInputRegisters(*mt.Address, mt.getQuantity())
		mt.ResultTime = time.Since(startTime)
	case ReadHoldingRegisters:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.ReadHoldingRegisters(*mt.Address, mt.getQuantity())
		mt.ResultTime = time.Since(startTime)
	case WriteSingleRegister:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.WriteSingleRegister(*mt.Address, binary.BigEndian.Uint16(mt.getWriteData()))
		mt.ResultTime = time.Since(startTime)
	case WriteMultipleRegisters:
		startTime := time.Now()
		mt.ResultByte, mt.ResultError = client.WriteMultipleRegisters(*mt.Address, mt.getQuantity(), mt.getWriteData())
		mt.ResultTime = time.Since(startTime)
	}
}

// TODO
func (mt *ModbusTest) Validation() error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	switch mt.getFunction() {
	case ReadDiscreteInputs, WriteMultipleRegisters, ReadCoils, ReadHoldingRegisters, ReadInputRegisters, WriteMultipleCoils:
	case WriteSingleCoil, WriteSingleRegister:
		if len(mt.Write) <= 0 {
			return fmt.Errorf("there is no data to write")
		}
	}
	return nil
}

func (mt *ModbusTest) getFunction() ModbusFunction {
	mFunc := strings.ReplaceAll(strings.ToLower(mt.Function), " ", "")
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

func (mt *ModbusTest) getQuantity() uint16 {

	if mt.Quantity != nil {
		return *mt.Quantity
	}

	// If it is not explicitly specified then we try to determine automatically
	switch mt.getFunction() {
	case ReadDiscreteInputs, ReadCoils:
		if bits, err := countBit(mt.Expected, false); err == nil {
			return bits
		}
	case WriteMultipleCoils:
		if bits, err := countBit(mt.Write, false); err == nil {
			return bits
		}
	case ReadInputRegisters, ReadHoldingRegisters:
		if bits, err := countBit(mt.Expected, true); err == nil {
			return bits
		}
	case WriteMultipleRegisters:
		if bits, err := countBit(mt.Write, true); err == nil {
			return bits
		}
	}
	return 0
}

func (mt *ModbusTest) getError() string {
	expected := mt.ExpectedError
	if mt.getFunction() != NilFunction {
		modbusError := strings.ReplaceAll(strings.ToLower(mt.ExpectedError), " ", "")
		if strings.HasPrefix(modbusError, "0x") {
			if a, err := strconv.ParseInt(strings.TrimPrefix(modbusError, "0x"), 16, 8); err == nil {
				modbusError = strconv.Itoa(int(a))
			}
		}
		switch modbusError {
		case "illegalfunction", "1":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 1}).Error()
		case "illegaldataaddress", "2":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 2}).Error()
		case "illegaldatavalue", "3":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 3}).Error()
		case "serverdevicefailure", "4":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 4}).Error()
		case "acknowledge", "5":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 5}).Error()
		case "serverdevicebusy", "6":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 6}).Error()
		case "memoryparityerror", "8":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 8}).Error()
		case "gatewaypathunavailable", "10":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 10}).Error()
		case "gatewaytargetdevicefailedtorespond", "11":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 11}).Error()
		}
	}
	return expected
}

func (mt *ModbusTest) getWriteData() []byte {
	data, err := valueToByte(mt.Write)
	if err != nil {
		return nil
	}
	return data
}
