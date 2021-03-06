package master

import (
	"encoding/binary"
	"fmt"
	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
	"rtu-test/e2e/common"
	"rtu-test/e2e/template"
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

type ModbusMasterTest struct {
	Name   string  `yaml:"name"`
	Skip   string  `yaml:"skip"`
	Before Message `yaml:"before"`
	// Переопределить адрес устройства, если на одной шине несколько блоков
	SlaveId    uint8           `yaml:"slaveId"`
	Function   string          `yaml:"function"`
	Address    *uint16         `yaml:"address"`
	Quantity   *uint16         `yaml:"quantity"`
	Write      []*common.Value `yaml:"write"`
	Expected   []*common.Value `yaml:"expected"`
	Success    Message         `yaml:"success"`
	Error      Message         `yaml:"error"`
	After      Message         `yaml:"after"`
	Fatal      string          `yaml:"fatal"`
	Disconnect bool            `yaml:"disconnect"`
}

func (mt *ModbusMasterTest) Run(client modbus.Client) ReportMasterTest {
	if err := mt.Validation(); err != nil {
		logrus.Fatal(err)
	}

	report := ReportMasterTest{Name: mt.Name, Pass: true, Skip: mt.Skip}
	logrus.Warn(common.Render(template.TestMasterModBusRUN, report))
	if report.Skip != "" {
		logrus.Warn(common.Render(template.TestMasterModBusSKIP, report))
		return report
	}
	mt.Before.PrintReportMasterTest(report)
	mt.Exec(client, &report)
	mt.Check(&report)
	if report.Pass {
		logrus.Warn(common.Render(template.TestMasterModBusPASS, report))
		mt.Success.PrintReportMasterTest(report)
	} else {
		logrus.Error(common.Render(template.TestMasterModBusFAIL, report))
		mt.Error.PrintReportMasterTest(report)
		if mt.Fatal != "" {
			logrus.Fatal(mt.Fatal)
		}
	}
	mt.After.PrintReportMasterTest(report)
	return report
}

func (mt *ModbusMasterTest) Check(report *ReportMasterTest) {
	countBit := 0
	var expected common.ReportExpected
	for _, v := range mt.Expected {
		bitSize := 8
		switch mt.getFunction() {
		case ReadHoldingRegisters, ReadInputRegisters, WriteSingleRegister, WriteMultipleRegisters:
			bitSize = 16
		}
		countBit, expected = v.Check(report.GotByte, report.GotTime, report.GotError, countBit, bitSize, binary.BigEndian)
		if !expected.Pass {
			report.Pass = false
		}
		report.Expected = append(report.Expected, expected)
	}
}

func (mt *ModbusMasterTest) Exec(client modbus.Client, report *ReportMasterTest) {
	var err error
	switch mt.getFunction() {
	case ReadDiscreteInputs:
		startTime := time.Now()
		if report.GotByte, err = client.ReadDiscreteInputs(*mt.Address, mt.getQuantity()); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	case ReadCoils:
		startTime := time.Now()
		if report.GotByte, err = client.ReadCoils(*mt.Address, mt.getQuantity()); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	case WriteSingleCoil:
		// Special case when writing single coil
		data := binary.BigEndian.Uint16(common.DataSingleCoil(common.ValueToByte(mt.Write)))
		report.Write = append(report.Write, common.ReportWrite{
			Name:    mt.Write[0].Name,
			Type:    common.Bool.String(),
			Data:    fmt.Sprintf("%t", data == 0),
			DataHex: fmt.Sprintf("%04x", data),
			DataBin: fmt.Sprintf("%08b", data),
		})
		startTime := time.Now()
		if report.GotByte, err = client.WriteSingleCoil(*mt.Address, data); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	case WriteMultipleCoils:
		for _, w := range mt.Write {
			report.Write = append(report.Write, w.ReportWrite(binary.BigEndian))
		}
		startTime := time.Now()
		if report.GotByte, err = client.WriteMultipleCoils(*mt.Address, mt.getQuantity(), common.ValueToByte(mt.Write)); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	case ReadInputRegisters:
		startTime := time.Now()
		if report.GotByte, err = client.ReadInputRegisters(*mt.Address, mt.getQuantity()); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	case ReadHoldingRegisters:
		startTime := time.Now()
		if report.GotByte, err = client.ReadHoldingRegisters(*mt.Address, mt.getQuantity()); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	case WriteSingleRegister:
		for _, w := range mt.Write {
			report.Write = append(report.Write, w.ReportWrite(binary.BigEndian))
		}
		startTime := time.Now()
		if report.GotByte, err = client.WriteSingleRegister(*mt.Address, binary.BigEndian.Uint16(common.ValueToByte16(mt.Write))); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	case WriteMultipleRegisters:
		for _, w := range mt.Write {
			report.Write = append(report.Write, w.ReportWrite(binary.BigEndian))
		}
		startTime := time.Now()
		if report.GotByte, err = client.WriteMultipleRegisters(*mt.Address, mt.getQuantity(), common.ValueToByte16(mt.Write)); err != nil {
			report.GotError = err.Error()
		}
		report.GotTime = time.Since(startTime)
	}
}

// TODO
func (mt *ModbusMasterTest) Validation() error {
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

	// переделываем формат ошибки
	for _, v := range mt.Expected {
		if v.Type() == common.Error {
			v.Error = mt.getError(*v.Error)
		}
	}
	return nil
}

func (mt *ModbusMasterTest) getFunction() ModbusFunction {
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

func (mt *ModbusMasterTest) getQuantity() uint16 {

	if mt.Quantity != nil {
		return *mt.Quantity
	}

	// If it is not explicitly specified then we try to determine automatically
	switch mt.getFunction() {
	case ReadDiscreteInputs, ReadCoils:
		return common.CountBit(mt.Expected, false)
	case WriteMultipleCoils:
		return common.CountBit(mt.Write, false)
	case ReadInputRegisters, ReadHoldingRegisters:
		return common.CountBit(mt.Expected, true)
	case WriteMultipleRegisters:
		return common.CountBit(mt.Write, true)
	}
	return 0
}

func (mt *ModbusMasterTest) getError(expected string) *string {
	if mt.getFunction() != NilFunction {
		modbusError := strings.ReplaceAll(strings.ToLower(expected), " ", "")
		if strings.HasPrefix(modbusError, "0x") {
			if a, err := strconv.ParseInt(strings.TrimPrefix(modbusError, "0x"), 16, 8); err == nil {
				modbusError = strconv.Itoa(int(a))
			}
		}
		switch modbusError {
		case "illegalfunction", "1":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 1}).Error()
		case "illegaldataaddress", "2":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 2}).Error()
		case "illegaldatavalue", "3":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 3}).Error()
		case "serverdevicefailure", "4":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 4}).Error()
		case "acknowledge", "5":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 5}).Error()
		case "serverdevicebusy", "6":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 6}).Error()
		case "memoryparityerror", "8":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 8}).Error()
		case "gatewaypathunavailable", "10":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 10}).Error()
		case "gatewaytargetdevicefailedtorespond", "11":
			expected = (&modbus.ModbusError{FunctionCode: byte(mt.getFunction()) | 1<<7, ExceptionCode: 11}).Error()
		}
	}
	return &expected
}
