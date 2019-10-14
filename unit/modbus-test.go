package unit

import (
	"encoding/binary"
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"strconv"
	"strings"
	"time"
)

type ModbusFunction int

const (
	NilFunction            = ModbusFunction(0)
	ReadCoils              = ModbusFunction(1)
	ReadDiscreteInputs     = ModbusFunction(2)
	ReadHoldingRegisters   = ModbusFunction(3)
	ReadInputRegisters     = ModbusFunction(4)
	WriteSingleCoil        = ModbusFunction(5)
	WriteSingleRegister    = ModbusFunction(6)
	WriteMultipleCoils     = ModbusFunction(15)
	WriteMultipleRegisters = ModbusFunction(16)
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

func (mt *ModbusTest) Exec(client modbus.Client) (err error) {
	log.Printf("Run %s", mt.Name)

	mt.Before.Print()

	switch mt.getFunction() {
	case ReadDiscreteInputs:
		err = mt.readDiscreteInputs(client)
	case ReadCoils:
		err = mt.readCoils(client)
	case WriteSingleCoil:
		err = mt.writeSingleCoil(client)
	case WriteMultipleCoils:
	case ReadInputRegisters:
		err = mt.readInputRegisters(client)
	case ReadHoldingRegisters:
		err = mt.readHoldingRegisters(client)
	case WriteSingleRegister:
		err = mt.writeSingleRegister(client)
	case WriteMultipleRegisters:
	default:
		return fmt.Errorf("function not found")
	}

	mt.Error.Print()

	mt.Success.Print()

	mt.After.Print()
	return nil
}

func (mt *ModbusTest) getQuantity() (uint16, error) {
	if mt.Quantity != nil {
		return *mt.Quantity, nil
		// TODO Возвращать ошибку если количество записываемых или считываемых данных не совпадает
	}

	switch mt.getFunction() {
	case ReadDiscreteInputs, ReadCoils:
		// TODO
	case WriteSingleCoil:
		// TODO
	case WriteMultipleCoils:
		// TODO
	case ReadInputRegisters, ReadHoldingRegisters:
		// TODO
	case WriteMultipleRegisters:
		// TODO
	}
	return 0, fmt.Errorf("quantity is nil")
}

func (mt *ModbusTest) readCoils(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	quantity, err := mt.getQuantity()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadCoils(*mt.Address, quantity)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) readDiscreteInputs(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	quantity, err := mt.getQuantity()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadDiscreteInputs(*mt.Address, quantity)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) readHoldingRegisters(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	quantity, err := mt.getQuantity()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadHoldingRegisters(*mt.Address, quantity)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) readInputRegisters(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	quantity, err := mt.getQuantity()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadInputRegisters(*mt.Address, quantity)
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
	switch mt.Write[0].Type() {
	case Bool:
		if *mt.Write[0].Bool {
			v = 0xff00
		} else {
			v = 0x0000
		}
	case Byte:
		b, err := mt.Write[0].Write()
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		if byteToEq(b, []byte{0xff, 0x00}) {
			v = 0xff00
		} else if byteToEq(b, []byte{0x00, 0x00}) {
			v = 0x0000
		} else {
			return fmt.Errorf("data error. Only supported 0xff00 0x0000")
		}
	default:
		return fmt.Errorf("invalid data type for record")
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
	if len(mt.Write) <= 0 {
		return fmt.Errorf("there is no data to write")
	}

	var v uint16
	var vBytes []byte
	var i int
	var vByte uint8
	for _, w := range mt.Write {
		switch w.Type() {
		case Bool:
			if *w.Bool {
				vByte = vByte | 1<<i
			}
			i++
			if i > 7 {
				vBytes = append(vBytes, vByte)
				vByte = 0
				i = 0
			}
		case Uint16, Int16, Int8, Uint8, Byte, String:
			if i != 0 {
				return fmt.Errorf("data error. the length of the binary type must be 8")
			}

			b, err := w.Write()
			if err != nil {
				return fmt.Errorf("%s", err)
			}
			vBytes = append(vBytes, b...)
		default:
			return fmt.Errorf("data error. Only supported Uint16, Int16, Int8, Uint8, Byte, String, Bool")
		}

		if len(vBytes) >= 2 {
			v = binary.BigEndian.Uint16(vBytes[:2])
			break
		}
	}

	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.WriteSingleRegister(*mt.Address, v)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func byteToEq(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
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
