package master

import (
	"github.com/schnack/gotest"
	"rtu-test/e2e/common"
	"testing"
	"time"
)

func TestModbusTest_Run(t *testing.T) {
	var param uint16 = 2
	var errorString string = ""
	var timeString string = "2s"
	var Address uint16 = 0
	var Quantity uint16 = 2
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "ReadHoldingRegisters",
		Address:  &Address,
		Quantity: &Quantity,
		Expected: []*common.Value{
			{Name: "param", Uint16: &param},
			{Name: "error", Error: &errorString},
			{Name: "time", Time: &timeString},
		},
	}

	client := NewFixtureModBusClient([]byte{0x00, 0x02}, nil)
	report := modbus.Run(client)

	if err := gotest.Expect(report.Expected[0].Name).Eq("param"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected[0].Pass).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.Expected[1].Name).Eq("error"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected[1].Pass).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.Expected[2].Name).Eq("time"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected[2].Pass).Eq(true); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_Check(t *testing.T) {
	var param uint8 = 2
	var errorString string = ""
	var timeString string = "2s"
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "ReadCoils",
		Expected: []*common.Value{
			{Name: "param", Uint8: &param},
			{Name: "error", Error: &errorString},
			{Name: "time", Time: &timeString},
		},
	}
	report := ReportMasterTest{GotByte: []byte{0x02}, GotError: errorString, GotTime: time.Second}
	modbus.Check(&report)

	if err := gotest.Expect(report.Expected[0].Name).Eq("param"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected[0].Pass).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.Expected[1].Name).Eq("error"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected[1].Pass).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.Expected[2].Name).Eq("time"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected[2].Pass).Eq(true); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecReadCoils(t *testing.T) {
	var Address uint16 = 0
	var Quantity uint16 = 2
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "ReadCoils",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0b00000011}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{3}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(Quantity); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecReadDiscreteInputs(t *testing.T) {
	var Address uint16 = 0
	var Quantity uint16 = 2
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "ReadDiscreteInputs",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0b00000011}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{3}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(Quantity); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecReadHoldingRegisters(t *testing.T) {
	var Address uint16 = 0
	var Quantity uint16 = 2
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "ReadHoldingRegisters",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(Quantity); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecReadInputRegisters(t *testing.T) {
	var Address uint16 = 0
	var Quantity uint16 = 2
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "Read Input Registers",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(Quantity); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecWriteSingleCoil(t *testing.T) {
	var param1 = true
	var Address uint16 = 0
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "Write Single Coil",
		Address:  &Address,
		Write: []*common.Value{
			{Name: "param1", Bool: &param1},
		},
	}

	client := NewFixtureModBusClient([]byte{0xff, 0x00}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{0xff, 0x00}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(uint16(0)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.SingleValue).Eq(uint16(0xff00)); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecWriteSingleRegister(t *testing.T) {
	var param1 = true
	var param2 uint8 = 1
	var Address uint16 = 0
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "Write Single Register",
		Address:  &Address,
		Write: []*common.Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(uint16(0)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.SingleValue).Eq(uint16(0x0101)); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecWriteMultipleCoils(t *testing.T) {
	var param1 = true
	var param2 uint8 = 1
	var Address uint16 = 0
	var Quantity uint16 = 16
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "Write Multiple Coils",
		Address:  &Address,
		Quantity: &Quantity,
		Write: []*common.Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(uint16(16)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Value).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_ExecWriteMultipleRegisters(t *testing.T) {
	var param1 = true
	var param2 uint8 = 1
	var Address uint16 = 0
	var Quantity uint16 = 1
	modbus := &ModbusMasterTest{
		Name:     "Test",
		Function: "Write Multiple Registers",
		Address:  &Address,
		Quantity: &Quantity,
		Write: []*common.Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)
	report := ReportMasterTest{}
	modbus.Exec(client, &report)

	if err := gotest.Expect(report.GotByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotTime > 0).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(report.GotError).Eq(""); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Address).Eq(Address); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Quantity).Eq(uint16(1)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(client.Value).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_getFunction(t *testing.T) {

	if err := gotest.Expect((&ModbusMasterTest{Function: "0x01"}).getFunction()).Eq(ReadCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "ReadCoils"}).getFunction()).Eq(ReadCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "read coils"}).getFunction()).Eq(ReadCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "bad function"}).getFunction()).Eq(NilFunction); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "ReadDiscreteInputs"}).getFunction()).Eq(ReadDiscreteInputs); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "WriteSingleCoil"}).getFunction()).Eq(WriteSingleCoil); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "WriteMultipleCoils"}).getFunction()).Eq(WriteMultipleCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "ReadInputRegisters"}).getFunction()).Eq(ReadInputRegisters); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "ReadHoldingRegisters"}).getFunction()).Eq(ReadHoldingRegisters); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "WriteSingleRegister"}).getFunction()).Eq(WriteSingleRegister); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "WriteMultipleRegisters"}).getFunction()).Eq(WriteMultipleRegisters); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_getQuantity(t *testing.T) {

	var quantity uint16 = 10
	if err := gotest.Expect((&ModbusMasterTest{Function: "ReadCoils", Quantity: &quantity}).getQuantity()).Eq(uint16(10)); err != nil {
		t.Error(err)
	}

	var param1 int64 = 2
	if err := gotest.Expect((&ModbusMasterTest{Function: "ReadCoils", Expected: []*common.Value{{Int64: &param1}}}).getQuantity()).Eq(uint16(64)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusMasterTest{Function: "ReadInputRegisters", Expected: []*common.Value{{Int64: &param1}}}).getQuantity()).Eq(uint16(4)); err != nil {
		t.Error(err)
	}

}

func TestModbusTest_getError(t *testing.T) {

	if err := gotest.Expect(*(&ModbusMasterTest{Function: "ReadCoils"}).getError("0x01")).Eq("modbus: exception '1' (illegal function), function '129'"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(*(&ModbusMasterTest{Function: "ReadCoils"}).getError("test")).Eq("test"); err != nil {
		t.Error(err)
	}

}
