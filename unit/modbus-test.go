package unit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"strconv"
	"strings"
)

type ModbusTest struct {
	Name          string        `yaml:"name"`
	Before        Message       `yaml:"before"`
	Function      string        `yaml:"function"`
	Address       *uint16       `yaml:"address"`
	Quantity      *uint16       `yaml:"quantity"`
	Write         []Value       `yaml:"write"`
	Expected      []Value       `yaml:"expected"`
	ExpectedError ExpectedError `yaml:"expectedError"`
	Success       Message       `yaml:"success"`
	Error         Message       `yaml:"error"`
	After         Message       `yaml:"after"`
	ResultByte    []byte        `yaml:"-"`
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
	case "writesinglecoil", "5":
		err = mt.WriteSingleCoil(client)
	case "writemultiplecoils", "15":
	case "readinputregisters", "4":
	case "readholdingregisters", "3":
	case "writesingleregister", "6":
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

// TODO возвращать ошибку
func (mt *ModbusTest) GetExceptError() error {
	/* expectError := strings.TrimPrefix(mt.ExpectedError, "0x")
	if expectError == "" {
		return nil
	}

	out, err := strconv.ParseUint(expectError, 16, 8)
	if err != nil {
		log.Fatalf("error test '%s' ExpectError: %s", u.Name, err)
	}
	return &modbus.ModbusError{FunctionCode: u.GetFunction() | 0x80, ExceptionCode: byte(out)}
	*/
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

	res, err := client.WriteSingleCoil(*mt.Address, v)
	// TODO проверка ошибки
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	if v == binary.BigEndian.Uint16(res) {
		return fmt.Errorf("wrong answer received")
	}
	return nil
}

func (mt *ModbusTest) GetWriteByte() ([]byte, error) {
	buff := new(bytes.Buffer)
	for _, val := range mt.Write {
		v, err := val.Write()
		if err != nil {
			return nil, err
		}
		buff.Write(v)
	}
	return buff.Bytes(), nil
}

func (mt *ModbusTest) ReadDiscreteInputs(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	var q uint16
	if mt.Quantity != nil {
		q = *mt.Quantity
	} else {
		// TODO если пустой то высчитывать автоматически
	}
	mt.ResultByte, mt.ResultError = client.ReadDiscreteInputs(*mt.Address, q)
	return nil
}

func (mt *ModbusTest) WriteMultipleCoils(client modbus.Client) error {
	if mt.Address == nil {
		return fmt.Errorf("address is nil")
	}
	var q uint16
	if mt.Quantity != nil {
		q = *mt.Quantity
	} else {
		// TODO если пустой то высчитывать автоматически
	}
	// TODO тут нужно правильно обрабатывать бинарные данные
	buff := new(bytes.Buffer)
	i := 0
	var b uint8 = 0
	for _, val := range mt.Write {
		if val.Type() == Bool {
			b = b | (1 << i)
			i++
		} else {
			i = 7
		}

		if i == 7 {
			buff.Write([]byte{b})
			b = 0
			i = 0
		}
		v, err := val.Write()
		if err != nil {
			return err
		}
		buff.Write(v)
	}
	mt.ResultByte, mt.ResultError = client.WriteMultipleCoils(*mt.Address, q, buff.Bytes())
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
