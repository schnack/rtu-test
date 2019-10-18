package unit

import (
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

func (mt *ModbusTest) CheckData() error {
	switch mt.getFunction() {
	case ReadCoils, ReadDiscreteInputs, ReadHoldingRegisters, ReadInputRegisters:
		// TODO
	case WriteSingleCoil, WriteSingleRegister, WriteMultipleCoils, WriteMultipleRegisters:
		// TODO
	}
	return nil
}

func (mt *ModbusTest) CheckDuration() error {
	expectTime := parseDuration(mt.ExpectedTime)
	if mt.ResultTime > expectTime {
		return fmt.Errorf("\nexpected: %s\n     got: %s\n", expectTime.String(), mt.ResultTime.String())
	}
	return nil
}

func (mt *ModbusTest) CheckError() error {
	if mt.ExpectedError == "" && mt.ResultError == nil {
		return nil
	} else if mt.ExpectedError == "" && mt.ResultError != nil {
		return fmt.Errorf("\nexpected:\n     got: %s\n", mt.ResultError.Error())
	}

	errorText := mt.ExpectedError
	if mt.getFunction() != NilFunction {
		modbusError := strings.ReplaceAll(strings.ToLower(mt.ExpectedError), " ", "")
		if strings.HasPrefix(modbusError, "0x") {
			if a, err := strconv.ParseInt(strings.TrimPrefix(modbusError, "0x"), 16, 8); err == nil {
				modbusError = strconv.Itoa(int(a))
			}
		}
		switch modbusError {
		case "illegalfunction", "1":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 1}).Error()
		case "illegaldataaddress", "2":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 2}).Error()
		case "illegaldatavalue", "3":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 3}).Error()
		case "serverdevicefailure", "4":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 4}).Error()
		case "acknowledge", "5":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 5}).Error()
		case "serverdevicebusy", "6":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 6}).Error()
		case "memoryparityerror", "8":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 8}).Error()
		case "gatewaypathunavailable", "10":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 10}).Error()
		case "gatewaytargetdevicefailedtorespond", "11":
			errorText = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()), ExceptionCode: 11}).Error()
		}
	}

	if mt.ResultError.Error() != errorText {
		return fmt.Errorf("\nexpected: %s\n     got: %s\n", errorText, mt.ResultError.Error())
	}

	return nil
}

func (mt *ModbusTest) Exec(client modbus.Client) (err error) {
	switch mt.getFunction() {
	case ReadDiscreteInputs:
		err = mt.readDiscreteInputs(client)
	case ReadCoils:
		err = mt.readCoils(client)
	case WriteSingleCoil:
		err = mt.writeSingleCoil(client)
	case WriteMultipleCoils:
		err = mt.writeMultipleCoils(client)
	case ReadInputRegisters:
		err = mt.readInputRegisters(client)
	case ReadHoldingRegisters:
		err = mt.readHoldingRegisters(client)
	case WriteSingleRegister:
		err = mt.writeSingleRegister(client)
	case WriteMultipleRegisters:
		err = mt.writeMultipleRegisters(client)
	default:
		err = fmt.Errorf("function not found")
	}
	return
}

func (mt *ModbusTest) readCoils(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadCoils(*mt.Address, mt.getQuantity())
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) readDiscreteInputs(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadDiscreteInputs(*mt.Address, mt.getQuantity())
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) readHoldingRegisters(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadHoldingRegisters(*mt.Address, mt.getQuantity())
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) readInputRegisters(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadInputRegisters(*mt.Address, mt.getQuantity())
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) writeSingleCoil(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	if len(mt.Write) <= 0 {
		return fmt.Errorf("there is no data to write")
	}

	var v uint16 = 0
	data := mt.getWriteData()
	if len(data) > 1 && (data[0] == 0xff || data[0] == 0x00) && data[0] == 0x00 {
		v = binary.BigEndian.Uint16(data[:2])
	} else if len(data) > 0 && data[0] == 0x01 || data[0] == 0x00 {
		v = 0x0000
		if data[0] == 0x01 {
			v = 0xff00
		}
	} else {
		return fmt.Errorf("data error. Only supported 1, 0, 0xff00, 0x0000")
	}

	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.WriteSingleCoil(*mt.Address, v)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) writeSingleRegister(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}

	data := mt.getWriteData()
	if len(data) < 2 {
		return fmt.Errorf("invalid data type for record")
	}
	v := binary.BigEndian.Uint16(data[:2])

	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.WriteSingleRegister(*mt.Address, v)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) writeMultipleCoils(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}

	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.WriteMultipleCoils(*mt.Address, mt.getQuantity(), mt.getWriteData())
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) writeMultipleRegisters(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}

	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.WriteMultipleRegisters(*mt.Address, mt.getQuantity(), mt.getWriteData())
	mt.ResultTime = time.Since(startTime)
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

func (mt *ModbusTest) getWriteData() []byte {
	data, err := valueToByte(mt.Write)
	if err != nil {
		return nil
	}
	return data
}
