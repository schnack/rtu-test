package common

import (
	"github.com/schnack/gotest"
	"testing"
	"time"
)

func Test_parseStringByte(t *testing.T) {
	var v string = "0x01 0x0001 0x000002"
	b, err := ParseStringByte(v)
	if err := gotest.Expect(b).Eq([]byte{1, 0, 1, 0, 0, 2}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func Test_dataSingleCoil(t *testing.T) {
	if err := gotest.Expect(DataSingleCoil([]byte{0x01})).Eq([]byte{0xff, 0x00}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(DataSingleCoil([]byte{0x00})).Eq([]byte{0x00, 0x00}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(DataSingleCoil([]byte{0xff, 0x00})).Eq([]byte{0xff, 0x00}); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(DataSingleCoil([]byte{0x00, 0x00})).Eq([]byte{0x00, 0x00}); err != nil {
		t.Error(err)
	}
}

func Test_countBit(t *testing.T) {

	var param1 int64 = 2
	count := CountBit([]*Value{{Int64: &param1}}, false)
	if err := gotest.Expect(count).Eq(uint16(64)); err != nil {
		t.Error(err)
	}

	count = CountBit([]*Value{{Int64: &param1}}, true)
	if err := gotest.Expect(count).Eq(uint16(4)); err != nil {
		t.Error(err)
	}

}

func Test_valueToByte(t *testing.T) {
	var param1 = true
	var param2 uint8 = 1
	var param3 uint16 = 1
	var param4 uint32 = 1
	var param5 uint64 = 1
	values := []*Value{
		{Bool: &param1},
		{Uint8: &param2},
		{Uint16: &param3},
		{Uint32: &param4},
		{Uint64: &param5},
	}
	data := ValueToByte(values)
	if err := gotest.Expect(data).Eq([]byte{1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}

}

func Test_valueToByte16(t *testing.T) {
	var param1 = true
	var param2 uint8 = 1
	var param3 uint16 = 1
	var param4 uint32 = 1
	var param5 uint64 = 1
	values := []*Value{
		{Bool: &param1},
		{Uint8: &param2},
		{Uint16: &param3},
		{Uint32: &param4},
		{Uint64: &param5},
	}
	data := ValueToByte16(values)
	if err := gotest.Expect(data).Eq([]byte{1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}

}

func Test_parsePauseNs(t *testing.T) {
	d := ParseDuration("1 ns")
	if err := gotest.Expect(d).Eq(time.Duration(1)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("1ns"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseUs(t *testing.T) {
	d := ParseDuration("1 us")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Microsecond); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("1µs"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseMs(t *testing.T) {
	d := ParseDuration("1 ms")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Millisecond); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(d.String()).Eq("1ms"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseS(t *testing.T) {
	d := ParseDuration("1 s")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Second); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("1s"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseM(t *testing.T) {
	d := ParseDuration("1 m")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Minute); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("1m0s"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseH(t *testing.T) {
	d := ParseDuration("1 h")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Hour); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("1h0m0s"); err != nil {
		t.Error(err)
	}
}

func Test_parsePause(t *testing.T) {
	d := ParseDuration("1")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Second); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("1s"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseEnter(t *testing.T) {
	d := ParseDuration("enter")
	if err := gotest.Expect(d).Eq(time.Duration(-1)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("-1ns"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseEmpty(t *testing.T) {
	d := ParseDuration("")
	if err := gotest.Expect(d).Eq(time.Duration(0)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(d.String()).Eq("0s"); err != nil {
		t.Error(err)
	}
}
