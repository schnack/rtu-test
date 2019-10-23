package unit

import (
	"encoding/binary"
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"math"
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

func (mt *ModbusTest) CheckData() (reports []Report) {
	switch mt.getFunction() {
	case ReadCoils, ReadDiscreteInputs, ReadHoldingRegisters, ReadInputRegisters:
		for _, v := range mt.Expected {
			report := Report{Pass: true, Name: v.Name, Type: v.Type()}
			switch report.Type {
			case Nil:
				report.Pass = true
			case Int8:
				report.Expected = []byte{uint8(*v.Int8)}
			case Int8Range:
				report.ExpectedMin = []byte{uint8(*v.MinInt8)}
				report.ExpectedMax = []byte{uint8(*v.MaxInt8)}
			case Int16:
				report.Expected = make([]byte, 2)
				binary.BigEndian.PutUint16(report.Expected, uint16(*v.Int16))
			case Int16Range:
				report.ExpectedMin = make([]byte, 2)
				binary.BigEndian.PutUint16(report.ExpectedMin, uint16(*v.MinInt16))
				report.ExpectedMax = make([]byte, 2)
				binary.BigEndian.PutUint16(report.ExpectedMax, uint16(*v.MaxInt16))
			case Int32:
				report.Expected = make([]byte, 4)
				binary.BigEndian.PutUint32(report.Expected, uint32(*v.Int32))
			case Int32Range:
				report.ExpectedMin = make([]byte, 4)
				binary.BigEndian.PutUint32(report.ExpectedMin, uint32(*v.MinInt32))
				report.ExpectedMax = make([]byte, 4)
				binary.BigEndian.PutUint32(report.ExpectedMax, uint32(*v.MaxInt32))
			case Int64:
				report.Expected = make([]byte, 8)
				binary.BigEndian.PutUint64(report.Expected, uint64(*v.Int64))
			case Int64Range:
				report.ExpectedMin = make([]byte, 8)
				binary.BigEndian.PutUint64(report.ExpectedMin, uint64(*v.MinInt64))
				report.ExpectedMax = make([]byte, 8)
				binary.BigEndian.PutUint64(report.ExpectedMax, uint64(*v.MaxInt64))
			case Uint8:
				report.Expected = []byte{*v.Uint8}
			case Uint8Range:
				report.ExpectedMin = []byte{*v.MinUint8}
				report.ExpectedMax = []byte{*v.MaxUint8}
			case Uint16:
				report.Expected = make([]byte, 2)
				binary.BigEndian.PutUint16(report.Expected, *v.Uint16)
			case Uint16Range:
				report.ExpectedMin = make([]byte, 2)
				binary.BigEndian.PutUint16(report.ExpectedMin, *v.MinUint16)
				report.ExpectedMax = make([]byte, 2)
				binary.BigEndian.PutUint16(report.ExpectedMax, *v.MaxUint16)
			case Uint32:
				report.Expected = make([]byte, 4)
				binary.BigEndian.PutUint32(report.Expected, *v.Uint32)
			case Uint32Range:
				report.ExpectedMin = make([]byte, 4)
				binary.BigEndian.PutUint32(report.ExpectedMin, *v.MinUint32)
				report.ExpectedMax = make([]byte, 4)
				binary.BigEndian.PutUint32(report.ExpectedMax, *v.MaxUint32)
			case Uint64:
				report.Expected = make([]byte, 8)
				binary.BigEndian.PutUint64(report.Expected, *v.Uint64)
			case Uint64Range:
				report.ExpectedMin = make([]byte, 8)
				binary.BigEndian.PutUint64(report.ExpectedMin, *v.MinUint64)
				report.ExpectedMax = make([]byte, 8)
				binary.BigEndian.PutUint64(report.ExpectedMax, *v.MaxUint64)
			case Float32:
				report.Expected = make([]byte, 4)
				binary.BigEndian.PutUint32(report.Expected, math.Float32bits(*v.Float32))
			case Float32Range:
				report.ExpectedMin = make([]byte, 4)
				binary.BigEndian.PutUint32(report.ExpectedMin, math.Float32bits(*v.Float32))
				report.ExpectedMax = make([]byte, 4)
				binary.BigEndian.PutUint32(report.ExpectedMax, math.Float32bits(*v.Float32))
			case Float64:
				report.Expected = make([]byte, 8)
				binary.BigEndian.PutUint64(report.Expected, math.Float64bits(*v.Float64))
			case Float64Range:
				report.ExpectedMin = make([]byte, 8)
				binary.BigEndian.PutUint64(report.ExpectedMin, math.Float64bits(*v.Float64))
				report.ExpectedMax = make([]byte, 8)
				binary.BigEndian.PutUint64(report.ExpectedMax, math.Float64bits(*v.Float64))
			case Bool:
			case String:
				report.Expected = []byte(*v.String)
			case Byte:
				var err error
				if report.Expected, err = v.GetByte(); err != nil {
					log.Panicf("%s", err)
				}
			}
			reports = append(reports, report)
		}
	case WriteSingleCoil:
		report := Report{Pass: true, Type: Byte, Got: mt.ResultByte}
		report.Expected = dataSingleCoil(mt.getWriteData())
		if !byteToEq(report.Expected, report.Got) {
			report.Pass = false
		}
		reports = append(reports, report)
	case WriteSingleRegister:
		report := Report{Pass: true, Type: Byte, Got: mt.ResultByte}
		report.Expected = mt.getWriteData()
		if !byteToEq(report.Expected[:2], report.Got) {
			report.Pass = false
		}
		reports = append(reports, report)
	case WriteMultipleCoils, WriteMultipleRegisters:
		report := Report{Pass: true, Type: Uint16, Got: mt.ResultByte, Expected: make([]byte, 2)}
		binary.BigEndian.PutUint16(report.Expected, mt.getQuantity())
		resultQuantity := binary.BigEndian.Uint16(report.Got)

		if mt.getQuantity() != resultQuantity {
			report.Pass = false
		}
		reports = append(reports, report)
	}
	return
}

func (mt *ModbusTest) CheckDuration() Report {
	expectTime := parseDuration(mt.ExpectedTime)
	report := Report{Name: "Execution time", Type: String, Expected: []byte(expectTime.String()), Got: []byte(mt.ResultTime.String()), Pass: true}

	if mt.ResultTime > expectTime {
		report.Pass = false
	}
	return report
}

func (mt *ModbusTest) CheckError() Report {
	report := Report{Name: "ModBus error", Type: String}
	if mt.ResultError != nil {
		report.Got = []byte(mt.ResultError.Error())
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
	report.Expected = []byte(errorText)

	report.Pass = string(report.Expected) == string(report.Got)

	return report
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

	data := dataSingleCoil(mt.getWriteData())
	if len(data) != 2 {
		return fmt.Errorf("data error. Only supported 1, 0, 0xff00, 0x0000")
	}
	v := binary.BigEndian.Uint16(data[:2])

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
