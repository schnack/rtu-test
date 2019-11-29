package e2e

import (
	"github.com/schnack/gotest"
	"sync"
	"testing"
)

func TestModbusSlave_Expect1Bit(t *testing.T) {
	var param1 = true
	var param2 uint8 = 0x17
	var param3 = false
	var param4 uint16 = 0x17_17
	var param5 uint32 = 0x17_17_17_17
	var param6 uint64 = 0x17_17_17_17_17_17_17_17

	values := []*Value{
		{Name: "param1", Bool: &param1},
		{Name: "param2", Address: "0x0002", Uint8: &param2},
		{Name: "param3", Bool: &param3},
		{Name: "param4", Uint16: &param4},
		{Name: "param5", Address: "0x001c", Uint32: &param5},
		{Name: "param6", Address: "0x003d", Uint64: &param6},
	}
	slave := ModbusSlave{}
	b := []byte{
		1,
		0,
		1, 1, 1, 0, 1, 0, 0, 0,
		0,
		1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0,
		0,
		1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0,
		0,
		1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0,
	}

	reports := slave.Expect1Bit(b, values, new(sync.Mutex))

	if err := gotest.Expect(len(reports)).Eq(6); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[0].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[0].Got).Eq("true"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[1].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[1].GotHex).Eq("17"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[2].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[2].Got).Eq("false"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[3].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[3].GotHex).Eq("1717"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[4].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[4].GotHex).Eq("17171717"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[5].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[5].GotHex).Eq("1717171717171717"); err != nil {
		t.Error(err)
	}
}

func TestModbusSlave_Expect16Bit(t *testing.T) {
	var param1 = true
	var param2 uint8 = 0x17
	var param3 uint8 = 0x18
	var param4 = true
	var param5 uint16 = 0x17_17
	var param6 uint32 = 0x17_00_00_17
	var param7 uint64 = 0x17_00_00_00_00_00_00_17
	slave := ModbusSlave{}
	values := []*Value{
		{Name: "param1", Bool: &param1},
		{Name: "param2", Address: "0x0002", Uint8: &param2},
		{Name: "param3", Uint8: &param3},
		{Name: "param4", Bool: &param4},
		{Name: "param5", Uint16: &param5},
		{Name: "param6", Address: "0x0006", Uint32: &param6},
		{Name: "param7", Address: "0x0009", Uint64: &param7},
	}
	b := []uint16{
		0x0001, //1
		0,
		0x1718, //2 - 3
		0x0001, //4
		0x1717, //5
		0,
		0x1700, //6
		0x0017, //6
		0,
		0x1700, //7
		0,      //7
		0,      //7
		0x0017, //7
	}

	reports := slave.Expect16Bit(b, values, new(sync.Mutex))

	if err := gotest.Expect(len(reports)).Eq(7); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[0].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[0].Got).Eq("true"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[1].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[1].GotHex).Eq("17"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[2].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[2].GotHex).Eq("18"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[3].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[3].Got).Eq("true"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[4].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[4].GotHex).Eq("1717"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[5].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[5].GotHex).Eq("17000017"); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[6].Pass).True(); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(reports[6].GotHex).Eq("1700000000000017"); err != nil {
		t.Error(err)
	}

}

func TestModbusSlave_Write1Bit(t *testing.T) {
	var param1 = true
	var param2 uint8 = 0x17
	var param3 = false
	var param4 uint16 = 0x17_17
	var param5 uint32 = 0x17_17_17_17
	var param6 uint64 = 0x17_17_17_17_17_17_17_17
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

	slave.Write1Bit(b, slave.Coils, new(sync.Mutex))

	if err := gotest.Expect(b).Eq([]byte{
		1,
		0,
		1, 1, 1, 0, 1, 0, 0, 0,
		0,
		1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0,
		0,
		1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0,
		0,
		1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0,
	}); err != nil {
		t.Error(err)
	}
}

func TestModbusSlave_Write16Bit(t *testing.T) {
	var param1 = true
	var param2 uint8 = 0x17
	var param3 uint8 = 0x18
	var param4 = true
	var param5 uint16 = 0x17_17
	var param6 uint32 = 0x17_00_00_17
	var param7 uint64 = 0x17_00_00_00_00_00_00_17
	slave := ModbusSlave{
		Coils: []*Value{
			{Name: "param1", Bool: &param1},
			{Name: "param2", Address: "0x0002", Uint8: &param2},
			{Name: "param3", Uint8: &param3},
			{Name: "param4", Bool: &param4},
			{Name: "param5", Uint16: &param5},
			{Name: "param6", Address: "0x0006", Uint32: &param6},
			{Name: "param7", Address: "0x0009", Uint64: &param7},
		},
	}
	b := make([]uint16, 13)

	slave.Write16Bit(b, slave.Coils, new(sync.Mutex))

	if err := gotest.Expect(b).Eq([]uint16{
		0x0001,
		0,
		0x1718,
		0x0001,
		0x1717,
		0,
		0x1700,
		0x0017,
		0,
		0x1700,
		0,
		0,
		0x0017,
	}); err != nil {
		t.Error(err)
	}
}
