package unit

import (
	"fmt"
	"github.com/schnack/gotest"
	"testing"
	"time"
)

func TestModbusTest_Check(t *testing.T) {
	var quantity uint16 = 2
	var param1 = false
	var param2 uint8 = 0x01
	modbus := &ModbusTest{
		Name:     "test",
		Quantity: &quantity,
		Function: "WriteMultipleCoils",
		Success:  Message{Message: "success"},
		Error:    Message{Message: "error"},
		Write: []Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
		ResultByte: []byte{0x00, 0x02},
	}
	if err := gotest.Expect(modbus.Check()).Eq(true); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_CheckData(t *testing.T) {
	var quantity uint16 = 2
	var param1 = true
	var param2 uint8 = 0x01
	modbus := &ModbusTest{
		Quantity: &quantity,
		Function: "ReadCoils",
		Expected: []Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
		ResultByte: []byte{0x01, 0x02},
	}
	if err := gotest.Expect(modbus.CheckData()).Eq(true); err != nil {
		t.Error(err)
	}
	param2 = 2
	if err := gotest.Expect(modbus.CheckData()).Eq(false); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_CheckDataWriteMultiple(t *testing.T) {
	var quantity uint16 = 2
	var param1 = false
	var param2 uint8 = 0x01
	modbus := &ModbusTest{
		Quantity: &quantity,
		Function: "WriteMultipleCoils",
		Write: []Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
		ResultByte: []byte{0x00, 0x02},
	}
	if err := gotest.Expect(modbus.CheckData()).Eq(true); err != nil {
		t.Error(err)
	}

	quantity = 3
	if err := gotest.Expect(modbus.CheckData()).Eq(false); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_CheckDataWriteSingleRegister(t *testing.T) {
	var param1 = false
	var param2 uint8 = 0x01
	modbus := &ModbusTest{
		Function: "WriteSingleRegister",
		Write: []Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
		ResultByte: []byte{0x00, 0x02},
	}

	if err := gotest.Expect(modbus.CheckData()).Eq(false); err != nil {
		t.Error(err)
	}

	param2 = 2
	if err := gotest.Expect(modbus.CheckData()).Eq(true); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_CheckDataWriteSingleCoil(t *testing.T) {
	var param1 = true
	modbus := &ModbusTest{
		Function: "WriteSingleCoil",
		Write: []Value{
			{Name: "param1", Bool: &param1},
		},
		ResultByte: []byte{0xFF, 0x00},
	}

	if err := gotest.Expect(modbus.CheckData()).Eq(true); err != nil {
		t.Error(err)
	}

	param1 = false
	if err := gotest.Expect(modbus.CheckData()).Eq(false); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_CheckDuration(t *testing.T) {
	modbus := &ModbusTest{
		ExpectedTime: "1s",
		ResultTime:   1 * time.Second,
	}

	if err := gotest.Expect(modbus.CheckDuration()).Eq(true); err != nil {
		t.Error(err)
	}

	modbus.ResultTime = 2 * time.Second
	if err := gotest.Expect(modbus.CheckDuration()).Eq(false); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_CheckError(t *testing.T) {
	modbus := &ModbusTest{
		ExpectedError: "",
		ResultError:   nil,
	}

	if err := gotest.Expect(modbus.CheckError()).Eq(true); err != nil {
		t.Error(err)
	}

	modbus.ResultError = fmt.Errorf("test error")
	if err := gotest.Expect(modbus.CheckError()).Eq(false); err != nil {
		t.Error(err)
	}

	modbus.ExpectedError = "0x01"
	if err := gotest.Expect(modbus.CheckError()).Eq(false); err != nil {
		t.Error(err)
	}

	modbus.Function = "0x01"
	if err := gotest.Expect(modbus.CheckError()).Eq(false); err != nil {
		t.Error(err)
	}

}

func TestModbusTest_ExecReadCoils(t *testing.T) {
	var Address uint16 = 0
	var Quantity uint16 = 2
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "ReadCoils",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0b00000011}, nil)
	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{3}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "ReadDiscreteInputs",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0b00000011}, nil)

	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{3}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "ReadHoldingRegisters",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)

	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "Read Input Registers",
		Address:  &Address,
		Quantity: &Quantity,
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)

	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "Write Single Coil",
		Address:  &Address,
		Write: []Value{
			{Name: "param1", Bool: &param1},
		},
	}

	client := NewFixtureModBusClient([]byte{0xff, 0x00}, nil)

	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{0xff, 0x00}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "Write Single Register",
		Address:  &Address,
		Write: []Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)

	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "Write Multiple Coils",
		Address:  &Address,
		Quantity: &Quantity,
		Write: []Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)

	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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
	modbus := &ModbusTest{
		Name:     "Test",
		Function: "Write Multiple Registers",
		Address:  &Address,
		Quantity: &Quantity,
		Write: []Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Uint8: &param2},
		},
	}

	client := NewFixtureModBusClient([]byte{0x01, 0x01}, nil)

	modbus.Exec(client)

	if err := gotest.Expect(modbus.ResultByte).Eq([]byte{0x01, 0x01}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultTime > 1).Eq(true); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(modbus.ResultError).Eq(nil); err != nil {
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

	if err := gotest.Expect((&ModbusTest{Function: "0x01"}).getFunction()).Eq(ReadCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "ReadCoils"}).getFunction()).Eq(ReadCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "read coils"}).getFunction()).Eq(ReadCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "bad function"}).getFunction()).Eq(NilFunction); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "ReadDiscreteInputs"}).getFunction()).Eq(ReadDiscreteInputs); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "WriteSingleCoil"}).getFunction()).Eq(WriteSingleCoil); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "WriteMultipleCoils"}).getFunction()).Eq(WriteMultipleCoils); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "ReadInputRegisters"}).getFunction()).Eq(ReadInputRegisters); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "ReadHoldingRegisters"}).getFunction()).Eq(ReadHoldingRegisters); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "WriteSingleRegister"}).getFunction()).Eq(WriteSingleRegister); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "WriteMultipleRegisters"}).getFunction()).Eq(WriteMultipleRegisters); err != nil {
		t.Error(err)
	}
}

func TestModbusTest_getQuantity(t *testing.T) {

	var quantity uint16 = 10
	if err := gotest.Expect((&ModbusTest{Function: "ReadCoils", Quantity: &quantity}).getQuantity()).Eq(uint16(10)); err != nil {
		t.Error(err)
	}

	var param1 int64 = 2
	if err := gotest.Expect((&ModbusTest{Function: "ReadCoils", Expected: []Value{{Int64: &param1}}}).getQuantity()).Eq(uint16(64)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect((&ModbusTest{Function: "ReadInputRegisters", Expected: []Value{{Int64: &param1}}}).getQuantity()).Eq(uint16(4)); err != nil {
		t.Error(err)
	}

}

func TestModbusTest_getWriteData(t *testing.T) {
	var param1 = true
	var param2 uint8 = 1
	var param3 uint16 = 1
	var param4 uint32 = 1
	var param5 uint64 = 1
	modbus := &ModbusTest{}
	modbus.Write = []Value{
		{Bool: &param1},
		{Uint8: &param2},
		{Uint16: &param3},
		{Uint32: &param4},
		{Uint64: &param5},
	}
	data := modbus.getWriteData()
	if err := gotest.Expect(data).Eq([]byte{1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}

}
