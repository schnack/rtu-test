package slave

import (
	"encoding/binary"
	"github.com/schnack/mbslave"
	"github.com/sirupsen/logrus"
	"math"
	"rtu-test/e2e/common"
	"rtu-test/e2e/modbus/master"
	"rtu-test/e2e/template"
	"strings"
	"time"
)

const (
	CoilsTable            = "coils"
	DiscreteInputTable    = "discreteInput"
	HoldingRegistersTable = "holdingRegisters"
	InputRegistersTable   = "inputRegisters"
)

type ModbusSlave struct {
	SlaveId        uint8  `yaml:"slaveId"`
	Port           string `yaml:"port"`
	BoundRate      int    `yaml:"boundRate"`
	DataBits       int    `yaml:"dataBits"`
	Parity         string `yaml:"parity"`
	StopBits       int    `yaml:"stopBits"`
	SilentInterval string `yaml:"silentInterval"`

	Coils            []*common.Value `yaml:"coils"`
	DiscreteInput    []*common.Value `yaml:"discreteInput"`
	HoldingRegisters []*common.Value `yaml:"holdingRegisters"`
	InputRegisters   []*common.Value `yaml:"inputRegisters"`

	Tests []*ModbusSlaveTest `yaml:"tests"`

	DataModel *mbslave.DefaultDataModel `yaml:"-"`

	currentTest *ModbusSlaveTest `yaml:"-"`
}

func (ms *ModbusSlave) getServer() *mbslave.Server {
	//# Parity: N - None, E - Even, O - Odd (default E)
	parity := mbslave.EvenParity
	switch strings.ToLower(ms.Parity) {
	case "n":
		parity = mbslave.NoParity
	case "e":
		parity = mbslave.EvenParity
	case "o":
		parity = mbslave.OddParity
	}

	stopBits := mbslave.TwoStopBits

	switch ms.StopBits {
	case 1:
		stopBits = mbslave.OneStopBit
	case 2:
		stopBits = mbslave.TwoStopBits
	}

	config := &mbslave.Config{
		Port:                 ms.Port,
		BaudRate:             ms.BoundRate,
		DataBits:             ms.DataBits,
		Parity:               parity,
		StopBits:             stopBits,
		SilentInterval:       common.ParseDuration(ms.SilentInterval),
		SlaveId:              ms.SlaveId,
		SizeDiscreteInputs:   math.MaxUint16,
		SizeCoils:            math.MaxUint16,
		SizeInputRegisters:   math.MaxUint16,
		SizeHoldingRegisters: math.MaxUint16,
	}
	ms.DataModel = mbslave.NewDefaultDataModel(config)
	transport := mbslave.NewRtuTransport(config)
	s := mbslave.NewServer(transport, ms.DataModel)

	ms.Write1Bit(CoilsTable, ms.Coils)
	ms.Write1Bit(DiscreteInputTable, ms.DiscreteInput)
	ms.Write16Bit(HoldingRegistersTable, ms.HoldingRegisters)
	ms.Write16Bit(InputRegistersTable, ms.InputRegisters)
	ms.DataModel.SetFunction(mbslave.FuncReadCoils, ms.ActionHandler)
	ms.DataModel.SetFunction(mbslave.FuncReadDiscreteInputs, ms.ActionHandler)
	ms.DataModel.SetFunction(mbslave.FuncReadHoldingRegisters, ms.ActionHandler)
	ms.DataModel.SetFunction(mbslave.FuncReadInputRegisters, ms.ActionHandler)
	ms.DataModel.SetFunction(mbslave.FuncWriteSingleCoil, ms.ActionHandler)
	ms.DataModel.SetFunction(mbslave.FuncWriteSingleRegister, ms.ActionHandler)
	ms.DataModel.SetFunction(mbslave.FuncWriteMultipleCoils, ms.ActionHandler)
	ms.DataModel.SetFunction(mbslave.FuncWriteMultipleRegisters, ms.ActionHandler)
	return s
}

func (ms *ModbusSlave) Run() error {
	s := ms.getServer()
	ms.autorun()
	return s.Listen()
}

func (ms *ModbusSlave) autorun() {
	for _, test := range ms.Tests {
		if test.AutoRun == "" || test.Skip != "" {
			continue
		}
		go func(t *ModbusSlaveTest) {
			autorun := strings.Split(t.AutoRun, "/")
			delay := autorun[0]
			timer := autorun[len(autorun)-1]
			time.Sleep(common.ParseDuration(delay))
			tiker := time.NewTicker(common.ParseDuration(timer))
			reports := ReportSlaveTest{
				Name: t.Name,
			}
			for _ = range tiker.C {
				if t.Lifetime != nil {
					if *t.Lifetime <= 0 {
						tiker.Stop()
						return
					}
					*t.Lifetime--
				}

				ms.before(t, reports)
				ms.expected(t, reports)
				ms.after(t, reports)

			}
		}(test)
	}
}

func (ms *ModbusSlave) before(test *ModbusSlaveTest, reports ReportSlaveTest) {
	if test == nil || test.Skip != "" {
		return
	}

	test.Before.PrintReportSlaveTest(reports)

	if test.BeforeWrite == nil {
		return
	}
	if v, ok := test.BeforeWrite[CoilsTable]; ok {
		ms.Write1Bit(CoilsTable, v)
	}
	if v, ok := test.BeforeWrite[DiscreteInputTable]; ok {
		ms.Write1Bit(DiscreteInputTable, v)
	}
	if v, ok := test.BeforeWrite[HoldingRegistersTable]; ok {
		ms.Write16Bit(HoldingRegistersTable, v)
	}
	if v, ok := test.BeforeWrite[InputRegistersTable]; ok {
		ms.Write16Bit(InputRegistersTable, v)
	}
}

func (ms *ModbusSlave) expected(test *ModbusSlaveTest, reports ReportSlaveTest) {
	if test == nil || test.Skip != "" || test.Expected == nil {
		return
	}

	logrus.Warn(common.Render(template.TestSlaveRUN, reports))

	if v, ok := test.Expected[CoilsTable]; ok {
		reports.ExpectedCoils, reports.Pass = ms.Expect1Bit(CoilsTable, v)
	}
	if v, ok := test.Expected[DiscreteInputTable]; ok {
		reports.ExpectedDiscreteInput, reports.Pass = ms.Expect1Bit(DiscreteInputTable, v)
	}
	if v, ok := test.Expected[HoldingRegistersTable]; ok {
		reports.ExpectedHoldingRegisters, reports.Pass = ms.Expect16Bit(HoldingRegistersTable, v)
	}
	if v, ok := test.Expected[InputRegistersTable]; ok {
		reports.ExpectedInputRegisters, reports.Pass = ms.Expect16Bit(InputRegistersTable, v)
	}

	if reports.Pass {
		logrus.Warn(common.Render(template.TestSlavePASS, reports))
		test.Success.PrintReportSlaveTest(reports)
	} else {
		logrus.Error(common.Render(template.TestSlaveFAIL, reports))
		test.Error.PrintReportSlaveTest(reports)
		if test.Fatal != "" {
			logrus.Fatal(test.Fatal)
		}
	}
}

func (ms *ModbusSlave) after(test *ModbusSlaveTest, reports ReportSlaveTest) {
	if test == nil || test.Skip != "" {
		return
	}

	if test.AfterWrite != nil {
		if v, ok := test.AfterWrite[CoilsTable]; ok {
			ms.Write1Bit(CoilsTable, v)
		}
		if v, ok := test.AfterWrite[DiscreteInputTable]; ok {
			ms.Write1Bit(DiscreteInputTable, v)
		}
		if v, ok := test.AfterWrite[HoldingRegistersTable]; ok {
			ms.Write16Bit(HoldingRegistersTable, v)
		}
		if v, ok := test.AfterWrite[InputRegistersTable]; ok {
			ms.Write16Bit(InputRegistersTable, v)
		}
	}

	test.After.PrintReportSlaveTest(reports)
}

func (ms *ModbusSlave) ActionHandler(request mbslave.Request, response mbslave.Response) {
	reports := ReportSlaveTest{}

	var test *ModbusSlaveTest
	max := 0
	var next []string

	if ms.currentTest != nil && ms.currentTest.Next != nil {
		next = ms.currentTest.Next
	}

	for i := range ms.Tests {
		ball := ms.Tests[i].Check(request, next)

		if ball != 0 && ball > max {
			test = ms.Tests[i]
			max = ball
		}
	}

	if test != nil {
		reports.Name = test.Name
		if test.Skip != "" {
			logrus.Warn(common.Render(template.TestSlaveSkip, reports))
		} else {
			if test.Lifetime != nil {
				*test.Lifetime--
			}
		}
	}

	ms.before(test, reports)

	switch master.ModbusFunction(request.GetFunction()) {
	case master.ReadCoils:
		ms.DataModel.ReadCoils(request, response)
	case master.ReadDiscreteInputs:
		ms.DataModel.ReadDiscreteInputs(request, response)
	case master.ReadHoldingRegisters:
		ms.DataModel.ReadHoldingRegisters(request, response)
	case master.ReadInputRegisters:
		ms.DataModel.ReadInputRegisters(request, response)
	case master.WriteSingleCoil:
		ms.DataModel.WriteSingleCoil(request, response)
	case master.WriteSingleRegister:
		ms.DataModel.WriteSingleRegister(request, response)
	case master.WriteMultipleCoils:
		ms.DataModel.WriteMultipleCoils(request, response)
	case master.WriteMultipleRegisters:
		ms.DataModel.WriteMultipleRegisters(request, response)
	}

	ms.expected(test, reports)
	ms.after(test, reports)

	if test != nil && test.Skip == "" {
		ms.currentTest = test
		time.Sleep(common.ParseDuration(test.TimeOut))
	}
	return
}

func (ms *ModbusSlave) Expect1Bit(table string, v []*common.Value) (reports []common.ReportExpected, pass bool) {
	pass = true
	var address uint16 = 0

	var getFunc func(address uint16) bool
	var countRegisters int

	switch table {
	case CoilsTable:
		countRegisters = ms.DataModel.LengthCoils()
		getFunc = ms.DataModel.GetCoils
	case DiscreteInputTable:
		countRegisters = ms.DataModel.LengthDiscreteInputs()
		getFunc = ms.DataModel.GetDiscreteInputs
	default:
		logrus.Fatalf("%s table is not supported", table)
		return
	}

	for i := range v {
		if v[i].Address != "" {
			rawAddress, err := common.ParseStringByte(v[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			address = binary.BigEndian.Uint16(rawAddress)
		}

		var buf []byte
		if v[i].LengthBit()%8 != 0 {
			buf = make([]byte, (v[i].LengthBit()/8)+1)
		} else {
			buf = make([]byte, v[i].LengthBit()/8)
		}

		for ii := 0; ii < v[i].LengthBit(); ii++ {
			if countRegisters <= int(address) {
				logrus.Fatal("ModBus tables overflow")
			}
			if getFunc(address) {
				buf[ii/8] |= 1 << (ii % 8)
			}
			address++
		}
		_, report := v[i].Check(buf, 0, "", 0, 8, binary.BigEndian)

		if !report.Pass {
			pass = false
		}
		reports = append(reports, report)
	}
	return
}

func (ms *ModbusSlave) Expect16Bit(table string, v []*common.Value) (reports []common.ReportExpected, pass bool) {
	pass = true
	var address uint16 = 0

	var getFunc func(address uint16) uint16

	switch table {
	case HoldingRegistersTable:
		getFunc = ms.DataModel.GetHoldingRegisters
	case InputRegistersTable:
		getFunc = ms.DataModel.GetInputRegisters
	default:
		logrus.Fatalf("%s table is not supported", table)
		return
	}

	countBit := 0
	for i := range v {

		if countBit != 0 {
			switch v[i].Type() {
			case common.Bool:
			case common.Uint8, common.Int8:
			case common.String, common.Byte:
				if len(v[i].Write(binary.BigEndian)) == 1 && countBit%8 != 0 {
					countBit += 8 - (countBit % 8)
				} else {
					address++
					countBit = 0
				}
			default:
				address++
				countBit = 0

			}
		}

		if v[i].Address != "" {
			rawAddress, err := common.ParseStringByte(v[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			address = binary.BigEndian.Uint16(rawAddress)
			countBit = 0
		}

		var buf []byte
		if v[i].LengthBit()%16 != 0 {
			if v[i].LengthBit()%16 > 7 {
				buf = make([]byte, (v[i].LengthBit()/8)+1)
			} else {
				buf = make([]byte, (v[i].LengthBit()/8)+2)
			}
		} else {
			buf = make([]byte, v[i].LengthBit()/8)
		}

		for ii := range buf {
			if ii%2 == 0 {
				continue
			}
			b := make([]byte, 2)
			binary.BigEndian.PutUint16(b, getFunc(address))
			buf[ii-1] = b[0]
			buf[ii] = b[1]
			if len(buf)-1 != ii {
				address++
			}
		}

		_, report := v[i].Check(buf, 0, "", countBit, 16, binary.BigEndian)

		switch v[i].Type() {
		case common.Bool:
			countBit++
			if countBit >= 16 {
				address++
				countBit = 0
			}
		case common.Uint8, common.Int8:
			if countBit%8 != 0 {
				countBit += 8 - (countBit % 8)
			}
			countBit += 8
			if countBit >= 16 {
				address++
				countBit = 0
			}
		case common.String, common.Byte:
			if len(v[i].Write(binary.BigEndian)) == 1 {
				if countBit%8 != 0 {
					countBit += 8 - (countBit % 8)
				}
				countBit += 8
				if countBit >= 16 {
					address++
					countBit = 0
				}
			} else {
				address++
				countBit = 0
			}
		default:
			address++
			countBit = 0
		}

		if !report.Pass {
			pass = false
		}
		reports = append(reports, report)
	}
	return
}

func (ms *ModbusSlave) Write1Bit(table string, v []*common.Value) {
	var address uint16 = 0

	var setFunc func(address uint16, value bool) error
	var countRegisters int

	switch table {
	case CoilsTable:
		countRegisters = ms.DataModel.LengthCoils()
		setFunc = ms.DataModel.SetCoils
	case DiscreteInputTable:
		countRegisters = ms.DataModel.LengthDiscreteInputs()
		setFunc = ms.DataModel.SetDiscreteInputs
	default:
		logrus.Fatalf("%s table is not supported", table)
		return
	}

	for i := range v {
		if v[i].Address != "" {
			rawAddress, err := common.ParseStringByte(v[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			address = binary.BigEndian.Uint16(rawAddress)
		}

		data := v[i].Write(binary.BigEndian)
		for _, b := range data {
			if countRegisters <= int(address) {
				logrus.Fatal("ModBus tables overflow")
			}
			if v[i].Type() == common.Bool {
				if err := setFunc(address, b != 0); err != nil {
					logrus.Fatalf("%s", err)
				}
				address++
			} else {
				for ii := 0; ii < 8; ii++ {
					if countRegisters <= int(address) {
						logrus.Fatal("ModBus tables overflow")
					}
					if err := setFunc(address, (b&(1<<ii)) != 0); err != nil {
						logrus.Fatalf("%s", err)
					}
					address++
				}
			}
		}
	}
}

func (ms *ModbusSlave) Write16Bit(table string, v []*common.Value) {
	var address uint16 = 0

	var setFunc func(address uint16, value uint16) error
	var countRegisters int

	switch table {
	case HoldingRegistersTable:
		countRegisters = ms.DataModel.LengthHoldingRegisters()
		setFunc = ms.DataModel.SetHoldingRegisters
	case InputRegistersTable:
		countRegisters = ms.DataModel.LengthInputRegisters()
		setFunc = ms.DataModel.SetInputRegisters
	default:
		logrus.Fatalf("%s table is not supported", table)
		return
	}

	var vBytes uint16 = 0
	current := 0

	for i := range v {
		if v[i].Address != "" {
			// Сбрасываем счетчик бит
			if current != 0 {
				if countRegisters <= int(address) {
					logrus.Fatal("ModBus tables overflow")
				}
				if err := setFunc(address, vBytes); err != nil {
					logrus.Fatalf("%s", err)
				}
				address++
				vBytes = 0
				current = 0
			}

			rawAddress, err := common.ParseStringByte(v[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			address = binary.BigEndian.Uint16(rawAddress)

		} else if current >= 16 {
			if countRegisters <= int(address) {
				logrus.Fatal("ModBus tables overflow")
			}
			if err := setFunc(address, vBytes); err != nil {
				logrus.Fatalf("%s", err)
			}
			address++
			vBytes = 0
			current = 0
		}

		switch v[i].Type() {
		case common.Bool:
			vBytes |= 1 << current
			current++
		default:
			data := v[i].Write(binary.BigEndian)
			if current < 8 && current != 0 {
				current += 8 - (current % 8)
			}

			if current < 16 && current != 0 && !(len(data) == 1 && current == 8) {
				if countRegisters <= int(address) {
					logrus.Fatal("ModBus tables overflow")
				}
				if err := setFunc(address, vBytes); err != nil {
					logrus.Fatalf("%s", err)
				}
				address++
				vBytes = 0
				current = 0
			}

			for _, b := range data {
				if current >= 16 {
					if countRegisters <= int(address) {
						logrus.Fatal("ModBus tables overflow")
					}
					if err := setFunc(address, vBytes); err != nil {
						logrus.Fatalf("%s", err)
					}
					address++
					vBytes = 0
					current = 0
				}
				if current/8 == 0 {
					vBytes |= uint16(b) << 8
				} else {
					vBytes |= uint16(b)
				}
				current += 8
			}
		}
	}
	if current != 0 {
		if countRegisters <= int(address) {
			logrus.Fatal("ModBus tables overflow")
		}
		if err := setFunc(address, vBytes); err != nil {
			logrus.Fatalf("%s", err)
		}
	}
}
