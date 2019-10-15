package unit

import (
	"github.com/schnack/gotest"
	"testing"
)

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
	data, err := modbus.getWriteData()
	if err := gotest.Expect(data).Eq([]byte{1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
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
