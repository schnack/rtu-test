package e2e

import (
	"github.com/schnack/gotest"
	"testing"
)

func TestModbusSlave_Write1Bit(t *testing.T) {
	var param1 = true
	var param2 uint8 = 0xC3
	var param3 = false
	var param4 uint16 = 0xF0_0F
	var param5 uint32 = 0xF0_00_00_0F
	var param6 uint64 = 0xF0_00_00_00_00_00_00_0F
	slave := ModbusSlave{
		Coils: []*Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Address: "0x0002", Uint8: &param2},
			{Name: "param3", Bool: &param3},
			{Name: "param4", Uint16: &param4},
			{Name: "param5", Address: "0x001c", Uint32: &param5},
			{Name: "param6", Address: "0x003d", Uint64: &param6},
		},
	}
	b := make([]byte, 125)

	slave.Write1Bit(b, slave.Coils)

	if err := gotest.Expect(b).Eq([]byte{
		1,
		0,
		1, 1, 0, 0, 0, 0, 1, 1,
		0,
		1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
		0,
		1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
		0,
		1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
	}); err != nil {
		t.Error(err)
	}

}
