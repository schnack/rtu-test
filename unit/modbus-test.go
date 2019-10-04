package unit

import (
	"bytes"
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
	b, err := mt.GetWriteByte()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	mt.ResultByte, mt.ResultError = client.WriteMultipleCoils(*mt.Address, q, b)
	return nil
}
