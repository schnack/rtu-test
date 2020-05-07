package slave

import (
	"github.com/schnack/gotest"
	"github.com/schnack/mbslave"
	"rtu-test/e2e/modbus/master"
	"testing"
)

func TestModbusSlaveTest_Check(t *testing.T) {
	var address uint16 = 0x0000
	var quantity uint16 = 0x0001
	mt := ModbusSlaveTest{
		Name:     "test 1",
		Lifetime: nil,
		Function: "0x02",
		Address:  &address,
		Quantity: &quantity,
		Data:     nil,
	}
	request := mbslave.NewRtuRequest([]byte{0xb1, 0x02, 0x00, 0x00, 0x00, 0x01, 0xa3, 0xfa})
	if err := gotest.Expect(request.Parse()).NotError(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(mt.Check(request, []string{})).Eq(111); err != nil {
		t.Error(err)
	}
}

func TestModbusSlaveTest_getFunction(t *testing.T) {
	mt := ModbusSlaveTest{
		Function: "0x02",
	}

	if err := gotest.Expect(mt.getFunction()).Eq(master.ReadDiscreteInputs); err != nil {
		t.Error(err)
	}
}
