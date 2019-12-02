package e2e

import (
	"fmt"
	"github.com/schnack/gotest"
	"testing"
	"time"
)

func TestValue_CheckTime(t *testing.T) {
	var param string = "2ms"
	v := Value{Name: "Test", Time: &param}
	raw := []byte{0x01, 0x02}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(0); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Time.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2ms"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("00000000001e8480"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[0000000000000000000000000000000000000000000111101000010010000000]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1s"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("000000003b9aca00"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[0000000000000000000000000000000000111011100110101100101000000000]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Nanosecond, "", offset, 8)
	if err := gotest.Expect(offset).Eq(0); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1ns"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckError(t *testing.T) {
	var param string = "error"
	v := Value{Name: "Test", Error: &param}
	raw := []byte{}

	offset, report := v.Check(raw, time.Second, "test", 0, 8)

	if err := gotest.Expect(offset).Eq(0); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Error.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("error"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("6572726f72"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[01100101 01110010 01110010 01101111 01110010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("74657374"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01110100 01100101 01110011 01110100]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Nanosecond, "error", offset, 8)
	if err := gotest.Expect(offset).Eq(0); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("error"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt8(t *testing.T) {
	var param int8 = 2
	v := Value{Name: "Test", Int8: &param}
	raw := []byte{0x01, 0x02}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("02"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt8Range(t *testing.T) {
	var paramMin int8 = 2
	var paramMax int8 = 4
	v := Value{Name: "Test", MinInt8: &paramMin, MaxInt8: &paramMax}
	raw := []byte{0x01, 0x02, 0x03, 0x04, 0x05}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "02", "04")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000010]", "[00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(40); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt8Min(t *testing.T) {
	var paramMin int8 = 2
	v := Value{Name: "Test", MinInt8: &paramMin}
	raw := []byte{0x01, 0x02, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "02", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt8Max(t *testing.T) {
	var paramMax int8 = 2
	v := Value{Name: "Test", MaxInt8: &paramMax}
	raw := []byte{0x01, 0x02, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "02")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt16(t *testing.T) {
	var param int16 = 2
	v := Value{Name: "Test", Int16: &param}
	raw := []byte{
		0x00, 0x01,
		0x00, 0x02}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("0002"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000000 00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt16Range(t *testing.T) {
	var paramMin int16 = 2
	var paramMax int16 = 4
	v := Value{Name: "Test", MinInt16: &paramMin, MaxInt16: &paramMax}
	raw := []byte{
		0, 0x01,
		0, 0x02,
		0, 0x03,
		0, 0x04,
		0, 0x05}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0002", "0004")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000010]", "[00000000 00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(80); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt16Min(t *testing.T) {
	var paramMin int16 = 2
	v := Value{Name: "Test", MinInt16: &paramMin}
	raw := []byte{
		0, 0x01,
		0, 0x02,
		0, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0002", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt16Max(t *testing.T) {
	var paramMax int16 = 2
	v := Value{Name: "Test", MaxInt16: &paramMax}
	raw := []byte{
		0, 0x01,
		0, 0x02,
		0, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "0002")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000000 00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt32(t *testing.T) {
	var param int32 = 2
	v := Value{Name: "Test", Int32: &param}
	raw := []byte{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("00000002"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000000 00000000 00000000 00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt32Range(t *testing.T) {
	var paramMin int32 = 2
	var paramMax int32 = 4
	v := Value{Name: "Test", MinInt32: &paramMin, MaxInt32: &paramMax}
	raw := []byte{
		0, 0, 0, 1,
		0, 0, 0, 2,
		0, 0, 0, 3,
		0, 0, 0, 4,
		0, 0, 0, 5,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "00000002", "00000004")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000010]", "[00000000 00000000 00000000 00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(160); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt32Min(t *testing.T) {
	var paramMin int32 = 2
	v := Value{Name: "Test", MinInt32: &paramMin}
	raw := []byte{
		0, 0, 0, 1,
		0, 0, 0, 2,
		0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "00000002", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt32Max(t *testing.T) {
	var paramMax int32 = 2
	v := Value{Name: "Test", MaxInt32: &paramMax}
	raw := []byte{
		0, 0, 0, 1,
		0, 0, 0, 2,
		0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "00000002")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000000 00000000 00000000 00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt64(t *testing.T) {
	var param int64 = 2
	v := Value{Name: "Test", Int64: &param}
	raw := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("0000000000000002"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt64Range(t *testing.T) {
	var paramMin int64 = 2
	var paramMax int64 = 4
	v := Value{Name: "Test", MinInt64: &paramMin, MaxInt64: &paramMax}
	raw := []byte{
		0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 0, 0, 3,
		0, 0, 0, 0, 0, 0, 0, 4,
		0, 0, 0, 0, 0, 0, 0, 5,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0000000000000002", "0000000000000004")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]", "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(256); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(320); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt64Min(t *testing.T) {
	var paramMin int64 = 2
	v := Value{Name: "Test", MinInt64: &paramMin}
	raw := []byte{
		0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0000000000000002", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt64Max(t *testing.T) {
	var paramMax int64 = 2
	v := Value{Name: "Test", MaxInt64: &paramMax}
	raw := []byte{
		0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Int64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "0000000000000002")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint8(t *testing.T) {
	var param uint8 = 2
	v := Value{Name: "Test", Uint8: &param}
	raw := []byte{0x01, 0x02}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("02"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint8Range(t *testing.T) {
	var paramMin uint8 = 2
	var paramMax uint8 = 4
	v := Value{Name: "Test", MinUint8: &paramMin, MaxUint8: &paramMax}
	raw := []byte{0x01, 0x02, 0x03, 0x04, 0x05}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "02", "04")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000010]", "[00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(40); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint8Min(t *testing.T) {
	var paramMin uint8 = 2
	v := Value{Name: "Test", MinUint8: &paramMin}
	raw := []byte{0x01, 0x02, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "02", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint8Max(t *testing.T) {
	var paramMax uint8 = 2
	v := Value{Name: "Test", MaxUint8: &paramMax}
	raw := []byte{0x01, 0x02, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "02")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint16(t *testing.T) {
	var param uint16 = 2
	v := Value{Name: "Test", Uint16: &param}
	raw := []byte{0x00, 0x01, 0x00, 0x02}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("0002"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000000 00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint16Range(t *testing.T) {
	var paramMin uint16 = 2
	var paramMax uint16 = 4
	v := Value{Name: "Test", MinUint16: &paramMin, MaxUint16: &paramMax}
	raw := []byte{
		0, 0x01,
		0, 0x02,
		0, 0x03,
		0, 0x04,
		0, 0x05}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0002", "0004")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000010]", "[00000000 00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(80); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint16Min(t *testing.T) {
	var paramMin uint16 = 2
	v := Value{Name: "Test", MinUint16: &paramMin}
	raw := []byte{
		0, 0x01,
		0, 0x02,
		0, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0002", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint16Max(t *testing.T) {
	var paramMax uint16 = 2
	v := Value{Name: "Test", MaxUint16: &paramMax}
	raw := []byte{
		0, 0x01,
		0, 0x02,
		0, 0x03}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "0002")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000000 00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint32(t *testing.T) {
	var param uint32 = 2
	v := Value{Name: "Test", Uint32: &param}
	raw := []byte{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("00000002"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000000 00000000 00000000 00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint32Range(t *testing.T) {
	var paramMin uint32 = 2
	var paramMax uint32 = 4
	v := Value{Name: "Test", MinUint32: &paramMin, MaxUint32: &paramMax}
	raw := []byte{
		0, 0, 0, 1,
		0, 0, 0, 2,
		0, 0, 0, 3,
		0, 0, 0, 4,
		0, 0, 0, 5,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "00000002", "00000004")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000010]", "[00000000 00000000 00000000 00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(160); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint32Min(t *testing.T) {
	var paramMin uint32 = 2
	v := Value{Name: "Test", MinUint32: &paramMin}
	raw := []byte{
		0, 0, 0, 1,
		0, 0, 0, 2,
		0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "00000002", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint32Max(t *testing.T) {
	var paramMax uint32 = 2
	v := Value{Name: "Test", MaxUint32: &paramMax}
	raw := []byte{
		0, 0, 0, 1,
		0, 0, 0, 2,
		0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "00000002")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000000 00000000 00000000 00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint64(t *testing.T) {
	var param uint64 = 2
	v := Value{Name: "Test", Uint64: &param}
	raw := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("0000000000000002"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint64Range(t *testing.T) {
	var paramMin uint64 = 2
	var paramMax uint64 = 4
	v := Value{Name: "Test", MinUint64: &paramMin, MaxUint64: &paramMax}
	raw := []byte{
		0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 0, 0, 3,
		0, 0, 0, 0, 0, 0, 0, 4,
		0, 0, 0, 0, 0, 0, 0, 5,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "4")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0000000000000002", "0000000000000004")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]", "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000100]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(256); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("4"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(320); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("5"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint64Min(t *testing.T) {
	var paramMin uint64 = 2
	v := Value{Name: "Test", MinUint64: &paramMin}
	raw := []byte{
		0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "0000000000000002", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint64Max(t *testing.T) {
	var paramMax uint64 = 2
	v := Value{Name: "Test", MaxUint64: &paramMax}
	raw := []byte{
		0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 0, 0, 3,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Uint64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "0000000000000002")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("3"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat32(t *testing.T) {
	var param float32 = 2.4
	v := Value{Name: "Test", Float32: &param}
	raw := []byte{
		0x40, 0x13, 0x33, 0x33,
		0x40, 0x19, 0x99, 0x9A,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2.400000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("4019999a"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[01000000 00011001 10011001 10011010]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("40133333"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00010011 00110011 00110011]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.400000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat32Range(t *testing.T) {
	var paramMin float32 = 2.2
	var paramMax float32 = 2.4
	v := Value{Name: "Test", MinFloat32: &paramMin, MaxFloat32: &paramMax}
	raw := []byte{
		0x40, 0x06, 0x66, 0x66,
		0x40, 0x0c, 0xcc, 0xcd,
		0x40, 0x13, 0x33, 0x33,
		0x40, 0x19, 0x99, 0x9A,
		0x40, 0x20, 0x00, 0x00,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2.200000", "2.400000")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "400ccccd", "4019999a")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[01000000 00001100 11001100 11001101]", "[01000000 00011001 10011001 10011010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.100000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("40066666"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00000110 01100110 01100110]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.200000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.400000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(160); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.500000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat32Min(t *testing.T) {
	var paramMin float32 = 2.2
	v := Value{Name: "Test", MinFloat32: &paramMin}
	raw := []byte{
		0x40, 0x06, 0x66, 0x66,
		0x40, 0x0c, 0xcc, 0xcd,
		0x40, 0x13, 0x33, 0x33,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2.200000", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "400ccccd", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[01000000 00001100 11001100 11001101]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.100000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("40066666"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00000110 01100110 01100110]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.200000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat32Max(t *testing.T) {
	var paramMax float32 = 2.2
	v := Value{Name: "Test", MaxFloat32: &paramMax}
	raw := []byte{
		0x40, 0x06, 0x66, 0x66,
		0x40, 0x0c, 0xcc, 0xcd,
		0x40, 0x13, 0x33, 0x33,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2.200000")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "400ccccd")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[01000000 00001100 11001100 11001101]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.100000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("40066666"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00000110 01100110 01100110]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.200000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat64(t *testing.T) {
	var param float64 = 2.4
	v := Value{Name: "Test", Float64: &param}
	raw := []byte{
		0x40, 0x02, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
		0x40, 0x03, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("2.400000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("4003333333333333"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[01000000 00000011 00110011 00110011 00110011 00110011 00110011 00110011]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("4002666666666666"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00000010 01100110 01100110 01100110 01100110 01100110 01100110]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.400000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat64Range(t *testing.T) {
	var paramMin float64 = 2.2
	var paramMax float64 = 2.4
	v := Value{Name: "Test", MinFloat64: &paramMin, MaxFloat64: &paramMax}
	raw := []byte{
		0x40, 0x00, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd,
		0x40, 0x01, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a,
		0x40, 0x02, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
		0x40, 0x03, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
		0x40, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2.200000", "2.400000")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "400199999999999a", "4003333333333333")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[01000000 00000001 10011001 10011001 10011001 10011001 10011001 10011010]", "[01000000 00000011 00110011 00110011 00110011 00110011 00110011 00110011]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.100000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("4000cccccccccccd"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00000000 11001100 11001100 11001100 11001100 11001100 11001101]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.200000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(256); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.400000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(320); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.500000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat64Min(t *testing.T) {
	var paramMin float64 = 2.2
	v := Value{Name: "Test", MinFloat64: &paramMin}
	raw := []byte{
		0x40, 0x00, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd,
		0x40, 0x01, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a,
		0x40, 0x02, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "2.200000", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "400199999999999a", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "[01000000 00000001 10011001 10011001 10011001 10011001 10011001 10011010]", "")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.100000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("4000cccccccccccd"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00000000 11001100 11001100 11001100 11001100 11001100 11001101]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.200000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckFloat64Max(t *testing.T) {
	var paramMax float64 = 2.2
	v := Value{Name: "Test", MaxFloat64: &paramMax}
	raw := []byte{
		0x40, 0x00, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd,
		0x40, 0x01, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a,
		0x40, 0x02, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Float64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq(fmt.Sprintf(FormatRange, "", "2.200000")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq(fmt.Sprintf(FormatRange, "", "400199999999999a")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq(fmt.Sprintf(FormatRange, "", "[01000000 00000001 10011001 10011001 10011001 10011001 10011001 10011010]")); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.100000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("4000cccccccccccd"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[01000000 00000000 11001100 11001100 11001100 11001100 11001100 11001101]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.200000"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("2.300000"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckBool(t *testing.T) {
	var param = true
	v := Value{Name: "Test", Bool: &param}
	raw := []byte{0b00000000, 0b00000101}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(1); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Bool.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("true"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("true"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("1"); err != nil {
		t.Error(err)
	}
	param = false
	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(2); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("false"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(3); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("true"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckString(t *testing.T) {
	var param string = "hello"
	v := Value{Name: "Test", String: &param}
	raw := []byte{
		104, 101, 108, 108, 111,
		104, 101, 108, 109, 111,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(40); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(String.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("hello"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("68656c6c6f"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[1101000 1100101 1101100 1101100 1101111]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("hello"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("68656c6c6f"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[1101000 1100101 1101100 1101100 1101111]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(80); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("helmo"); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckByte(t *testing.T) {
	var param string = "0x01 0x0003"
	v := Value{Name: "Test", Byte: &param}
	raw := []byte{
		0x01, 0x00, 0x03,
		0x01, 0x00, 0x04,
	}

	offset, report := v.Check(raw, time.Second, "", 0, 8)

	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("Test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Type).Eq(Byte.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Expected).Eq("01 00 03"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedHex).Eq("010003"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.ExpectedBin).Eq("[00000001 00000000 00000011]"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("01 00 03"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotHex).Eq("010003"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.GotBin).Eq("[00000001 00000000 00000011]"); err != nil {
		t.Error(err)
	}

	offset, report = v.Check(raw, time.Second, "", offset, 8)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Got).Eq("01 00 04"); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt8(t *testing.T) {
	var v int8 = 1
	b := (&Value{Name: "test", Int8: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt16(t *testing.T) {
	var v int16 = 1
	b := (&Value{Name: "test", Int16: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt32(t *testing.T) {
	var v int32 = 1
	b := (&Value{Name: "test", Int32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt64(t *testing.T) {
	var v int64 = 1
	b := (&Value{Name: "test", Int64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint8(t *testing.T) {
	var v uint8 = 1
	b := (&Value{Name: "test", Uint8: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint16(t *testing.T) {
	var v uint16 = 1
	b := (&Value{Name: "test", Uint16: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint32(t *testing.T) {
	var v uint32 = 1
	b := (&Value{Uint32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint64(t *testing.T) {
	var v uint64 = 1
	b := (&Value{Name: "test", Uint64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteFloat32(t *testing.T) {
	var v float32 = 1.2
	b := (&Value{Name: "test", Float32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{63, 153, 153, 154}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteFloat64(t *testing.T) {
	var v float64 = 1.2
	b := (&Value{Name: "test", Float64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{63, 243, 51, 51, 51, 51, 51, 51}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteBoolTrue(t *testing.T) {
	var v bool = true
	b := (&Value{Name: "test", Bool: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteBoolFalse(t *testing.T) {
	var v bool = false
	b := (&Value{Name: "test", Bool: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteString(t *testing.T) {
	var v string = "test"
	b := (&Value{Name: "test", String: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{116, 101, 115, 116}); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteByte(t *testing.T) {
	var v string = "0x01 0x0001 0x000002"
	b := (&Value{Name: "test", Byte: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1, 0, 1, 0, 0, 2}); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteInt8(t *testing.T) {
	var v int8 = 1
	report := (&Value{Name: "test", Int8: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Int8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteInt16(t *testing.T) {
	var v int16 = 1
	report := (&Value{Name: "test", Int16: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Int16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteInt32(t *testing.T) {
	var v int32 = 1
	report := (&Value{Name: "test", Int32: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Int32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteInt64(t *testing.T) {
	var v int64 = 1
	report := (&Value{Name: "test", Int64: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Int64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteUint8(t *testing.T) {
	var v uint8 = 1
	report := (&Value{Name: "test", Uint8: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Uint8.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteUint16(t *testing.T) {
	var v uint16 = 1
	report := (&Value{Name: "test", Uint16: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Uint16.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("0001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000000 00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteUint32(t *testing.T) {
	var v uint32 = 1
	report := (&Value{Name: "test", Uint32: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Uint32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("00000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}

}

func TestValue_ReportWriteUint64(t *testing.T) {
	var v uint64 = 1
	report := (&Value{Name: "test", Uint64: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Uint64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("0000000000000001"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteFloat32(t *testing.T) {
	var v float32 = 1.2
	report := (&Value{Name: "test", Float32: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Float32.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1.200000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("3f99999a"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00111111 10011001 10011001 10011010]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteFloat64(t *testing.T) {
	var v float64 = 1.2
	report := (&Value{Name: "test", Float64: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Float64.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("1.200000"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("3ff3333333333333"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00111111 11110011 00110011 00110011 00110011 00110011 00110011 00110011]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteBoolTrue(t *testing.T) {
	var v bool = true
	report := (&Value{Name: "test", Bool: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Bool.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("true"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("01"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000001]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteBoolFalse(t *testing.T) {
	var v bool = false
	report := (&Value{Name: "test", Bool: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Bool.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("false"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("00"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000000]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteString(t *testing.T) {
	var v string = "test"
	report := (&Value{Name: "test", String: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(String.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("74657374"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[01110100 01100101 01110011 01110100]"); err != nil {
		t.Error(err)
	}
}

func TestValue_ReportWriteByte(t *testing.T) {
	var v string = "0x01 0x0001 0x000002"
	report := (&Value{Name: "test", Byte: &v}).ReportWrite()

	if err := gotest.Expect(report.Type).Eq(Byte.String()); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Name).Eq("test"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.Data).Eq("01 00 01 00 00 02"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataHex).Eq("010001000002"); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(report.DataBin).Eq("[00000001 00000000 00000001 00000000 00000000 00000010]"); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeEmpty(t *testing.T) {
	if err := gotest.Expect((&Value{}).Type()).Eq(Nil); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeInt8(t *testing.T) {
	var v int8 = 1
	if err := gotest.Expect((&Value{Int8: &v}).Type()).Eq(Int8); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeInt16(t *testing.T) {
	var v int16 = 1
	if err := gotest.Expect((&Value{Int16: &v}).Type()).Eq(Int16); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeInt32(t *testing.T) {
	var v int32 = 1
	if err := gotest.Expect((&Value{Int32: &v}).Type()).Eq(Int32); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeInt64(t *testing.T) {
	var v int64 = 1
	if err := gotest.Expect((&Value{Int64: &v}).Type()).Eq(Int64); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeUint8(t *testing.T) {
	var v uint8 = 1
	if err := gotest.Expect((&Value{Uint8: &v}).Type()).Eq(Uint8); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeUint16(t *testing.T) {
	var v uint16 = 1
	if err := gotest.Expect((&Value{Uint16: &v}).Type()).Eq(Uint16); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeUint32(t *testing.T) {
	var v uint32 = 1
	if err := gotest.Expect((&Value{Uint32: &v}).Type()).Eq(Uint32); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeUint64(t *testing.T) {
	var v uint64 = 1
	if err := gotest.Expect((&Value{Uint64: &v}).Type()).Eq(Uint64); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeFloat32(t *testing.T) {
	var v float32 = 1
	if err := gotest.Expect((&Value{Float32: &v}).Type()).Eq(Float32); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeFloat64(t *testing.T) {
	var v float64 = 1
	if err := gotest.Expect((&Value{Float64: &v}).Type()).Eq(Float64); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeBool(t *testing.T) {
	var v bool = true
	if err := gotest.Expect((&Value{Bool: &v}).Type()).Eq(Bool); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeString(t *testing.T) {
	var v string = "test"
	if err := gotest.Expect((&Value{String: &v}).Type()).Eq(String); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeByte(t *testing.T) {
	var v string = "0x01020304"
	if err := gotest.Expect((&Value{Byte: &v}).Type()).Eq(Byte); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeTime(t *testing.T) {
	var v string = "1s"
	if err := gotest.Expect((&Value{Time: &v}).Type()).Eq(Time); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeError(t *testing.T) {
	var v string = "error"
	if err := gotest.Expect((&Value{Error: &v}).Type()).Eq(Error); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBitEmpty(t *testing.T) {
	if err := gotest.Expect((&Value{}).LengthBit()).Eq(0); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBit8(t *testing.T) {
	var v int8 = 1
	if err := gotest.Expect((&Value{Int8: &v}).LengthBit()).Eq(8); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBit16(t *testing.T) {
	var v int16 = 1
	if err := gotest.Expect((&Value{Int16: &v}).LengthBit()).Eq(16); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBit32(t *testing.T) {
	var v int32 = 1
	if err := gotest.Expect((&Value{Int32: &v}).LengthBit()).Eq(32); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBit64(t *testing.T) {
	var v int64 = 1
	if err := gotest.Expect((&Value{Int64: &v}).LengthBit()).Eq(64); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBitBool(t *testing.T) {
	var v bool = true
	if err := gotest.Expect((&Value{Bool: &v}).LengthBit()).Eq(1); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBitString(t *testing.T) {
	var v string = "test"
	if err := gotest.Expect((&Value{String: &v}).LengthBit()).Eq(8 * 4); err != nil {
		t.Error(err)
	}
}

func TestValue_LengthBitByte(t *testing.T) {
	var v string = "0x01020304"
	if err := gotest.Expect((&Value{Byte: &v}).LengthBit()).Eq(8 * 4); err != nil {
		t.Error(err)
	}
}
