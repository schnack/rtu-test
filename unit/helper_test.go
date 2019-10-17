package unit

import (
	"github.com/schnack/gotest"
	"testing"
	"time"
)

func Test_countBit(t *testing.T) {

	var param1 int64 = 2
	count, err := countBit([]Value{{Int64: &param1}}, false)
	if err := gotest.Expect(count).Eq(uint16(64)); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}

	count, err = countBit([]Value{{Int64: &param1}}, true)
	if err := gotest.Expect(count).Eq(uint16(4)); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}

}

func Test_valueToByte(t *testing.T) {
	var param1 = true
	var param2 uint8 = 1
	var param3 uint16 = 1
	var param4 uint32 = 1
	var param5 uint64 = 1
	values := []Value{
		{Bool: &param1},
		{Uint8: &param2},
		{Uint16: &param3},
		{Uint32: &param4},
		{Uint64: &param5},
	}
	data, err := valueToByte(values)
	if err := gotest.Expect(data).Eq([]byte{1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}

}

func Test_parsePauseNs(t *testing.T) {
	d, tp := parseDuration("1 ns")
	if err := gotest.Expect(d).Eq(time.Duration(1)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("ns"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseUs(t *testing.T) {
	d, tp := parseDuration("1 us")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Microsecond); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("us"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseMs(t *testing.T) {
	d, tp := parseDuration("1 ms")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Millisecond); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("ms"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseS(t *testing.T) {
	d, tp := parseDuration("1 s")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Second); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("s"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseM(t *testing.T) {
	d, tp := parseDuration("1 m")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Minute); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("m"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseH(t *testing.T) {
	d, tp := parseDuration("1 h")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Hour); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("h"); err != nil {
		t.Error(err)
	}
}

func Test_parsePause(t *testing.T) {
	d, tp := parseDuration("1")
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Second); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("s"); err != nil {
		t.Error(err)
	}
}

func Test_parsePauseEnter(t *testing.T) {
	d, tp := parseDuration("enter")
	if err := gotest.Expect(d).Eq(time.Duration(-1)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq(""); err != nil {
		t.Error(err)
	}
}
