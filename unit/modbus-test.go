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

	mfunc := strings.ReplaceAll(strings.ToLower(mt.Function), " ", "")
	if strings.HasPrefix(mfunc, "0x") {
		if a, err := strconv.ParseInt(strings.TrimPrefix(mfunc, "0x"), 16, 8); err == nil {
			mfunc = strconv.Itoa(int(a))
		}
	}
	switch mfunc {
	case "readdiscreteinputs", "2":
		err = mt.ReadDiscreteInputs(client)
	case "readcoils", "1":
		err = mt.ReadCoils(client)
	case "writesinglecoil", "5":
		err = mt.WriteSingleCoil(client)
	case "writemultiplecoils", "15":
	case "readinputregisters", "4":
		err = mt.ReadInputRegisters(client)
	case "readholdingregisters", "3":
		err = mt.ReadHoldingRegisters(client)
	case "writesingleregister", "6":
		err = mt.WriteSingleRegister(client)
	case "writemultipleregisters", "16":
	case "readwritemultipleregisters", "23":
	case "maskwriteregister", "22":
	case "readfifqqueue", "24":
	default:
		return fmt.Errorf("function not found")
	}

	mt.Error.Print()

	mt.Success.Print()

	mt.After.Print()
	return nil
}

func (mt *ModbusTest) ReadCoils(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	if mt.Quantity == nil {
		return fmt.Errorf("quantity is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadCoils(*mt.Address, *mt.Quantity)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) ReadDiscreteInputs(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	if mt.Quantity == nil {
		return fmt.Errorf("quantity is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadDiscreteInputs(*mt.Address, *mt.Quantity)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) ReadHoldingRegisters(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	if mt.Quantity == nil {
		return fmt.Errorf("quantity is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadHoldingRegisters(*mt.Address, *mt.Quantity)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) ReadInputRegisters(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	if mt.Quantity == nil {
		return fmt.Errorf("quantity is nil")
	}
	startTime := time.Now()
	mt.ResultByte, mt.ResultError = client.ReadInputRegisters(*mt.Address, *mt.Quantity)
	mt.ResultTime = time.Since(startTime)
	return nil
}

func (mt *ModbusTest) WriteSingleCoil(client modbus.Client) error {
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

func (mt *ModbusTest) WriteSingleRegister(client modbus.Client) error {
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
