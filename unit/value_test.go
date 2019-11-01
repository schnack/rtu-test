package unit

import (
	"fmt"
	"github.com/schnack/gotest"
	"testing"
)

func TestValue_StringExpectedInt8(t *testing.T) {
	var param int8 = 2
	v := Value{Name: "Test", Int8: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Int8, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedInt16(t *testing.T) {
	var param int16 = 2
	v := Value{Name: "Test", Int16: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Int16, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedInt32(t *testing.T) {
	var param int32 = 2
	v := Value{Name: "Test", Int32: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Int32, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedInt64(t *testing.T) {
	var param int64 = 2
	v := Value{Name: "Test", Int64: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Int64, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedUint8(t *testing.T) {
	var param uint8 = 2
	v := Value{Name: "Test", Uint8: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Uint8, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedUint16(t *testing.T) {
	var param uint16 = 2
	v := Value{Name: "Test", Uint16: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Uint16, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedUint32(t *testing.T) {
	var param uint32 = 2
	v := Value{Name: "Test", Uint32: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Uint32, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedUint64(t *testing.T) {
	var param uint64 = 2
	v := Value{Name: "Test", Uint64: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatDecimal, Uint64, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedFloat32(t *testing.T) {
	var param float32 = 2.2
	v := Value{Name: "Test", Float32: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatFloat, Float32, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedFloat64(t *testing.T) {
	var param float64 = 2.2
	v := Value{Name: "Test", Float64: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatFloat, Float64, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedBool(t *testing.T) {
	var param bool = true
	v := Value{Name: "Test", Bool: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatBool, Bool, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedString(t *testing.T) {
	var param string = "hello"
	v := Value{Name: "Test", Str: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatString, String, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringExpectedByte(t *testing.T) {
	var param string = "0x01 0x0002"
	v := Value{Name: "Test", Byte: &param}

	if err := gotest.Expect(v.StringExpected()).Eq(fmt.Sprintf(FormatByte, Byte, []byte{0x01, 0x00, 0x02})); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotInt8(t *testing.T) {
	var param int8 = 2
	v := Value{Name: "Test", Int8: &param, GotInt8: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Int8, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotInt16(t *testing.T) {
	var param int16 = 2
	v := Value{Name: "Test", Int16: &param, GotInt16: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Int16, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotInt32(t *testing.T) {
	var param int32 = 2
	v := Value{Name: "Test", Int32: &param, GotInt32: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Int32, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotInt64(t *testing.T) {
	var param int64 = 2
	v := Value{Name: "Test", Int64: &param, GotInt64: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Int64, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotUint8(t *testing.T) {
	var param uint8 = 2
	v := Value{Name: "Test", Uint8: &param, GotUint8: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Uint8, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotUint16(t *testing.T) {
	var param uint16 = 2
	v := Value{Name: "Test", Uint16: &param, GotUint16: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Uint16, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotUint32(t *testing.T) {
	var param uint32 = 2
	v := Value{Name: "Test", Uint32: &param, GotUint32: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Uint32, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotUint64(t *testing.T) {
	var param uint64 = 2
	v := Value{Name: "Test", Uint64: &param, GotUint64: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatDecimal, Uint64, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotFloat32(t *testing.T) {
	var param float32 = 2.2
	v := Value{Name: "Test", Float32: &param, GotFloat32: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatFloat, Float32, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotFloat64(t *testing.T) {
	var param float64 = 2.2
	v := Value{Name: "Test", Float64: &param, GotFloat64: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatFloat, Float64, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotBool(t *testing.T) {
	var param bool = true
	v := Value{Name: "Test", Bool: &param, GotBool: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatBool, Bool, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotString(t *testing.T) {
	var param string = "hello"
	v := Value{Name: "Test", Str: &param, GotString: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatString, String, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_StringGotByte(t *testing.T) {
	var param []byte = []byte{0x01, 0x02}
	var param1 string = "0x01, 0x02"
	v := Value{Name: "Test", Byte: &param1, GotByte: param}

	if err := gotest.Expect(v.StringGot()).Eq(fmt.Sprintf(FormatByte, Byte, param)); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt8(t *testing.T) {
	var param int8 = 2
	v := Value{Name: "Test", Int8: &param}
	raw := []byte{0x01, 0x02}

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt8).Eq(int8(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt8).Eq(int8(2)); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckInt8Range(t *testing.T) {
	var paramMin int8 = 2
	var paramMax int8 = 4
	v := Value{Name: "Test", MinInt8: &paramMin, MaxInt8: &paramMax}
	raw := []byte{0x01, 0x02, 0x03, 0x04, 0x05}

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt8).Eq(int8(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt8).Eq(int8(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt8).Eq(int8(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt8).Eq(int8(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(40); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt8).Eq(int8(5)); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckInt16(t *testing.T) {
	var param int16 = 2
	v := Value{Name: "Test", Int16: &param}
	raw := []byte{0x00, 0x01, 0x00, 0x02}

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt16).Eq(int16(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt16).Eq(int16(2)); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckInt16Range(t *testing.T) {
	var paramMin int16 = 2
	var paramMax int16 = 4
	v := Value{Name: "Test", MinInt16: &paramMin, MaxInt16: &paramMax}
	raw := []byte{
		0, 1,
		0, 2,
		0, 3,
		0, 4,
		0, 5,
	}

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt16).Eq(int16(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt16).Eq(int16(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt16).Eq(int16(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt16).Eq(int16(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(80); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt16).Eq(int16(5)); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckInt32(t *testing.T) {
	var param int32 = 2
	v := Value{Name: "Test", Int32: &param}
	raw := []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02}

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt32).Eq(int32(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt32).Eq(int32(2)); err != nil {
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

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt32).Eq(int32(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt32).Eq(int32(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt32).Eq(int32(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt32).Eq(int32(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(160); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt32).Eq(int32(5)); err != nil {
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

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt64).Eq(int64(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt64).Eq(int64(2)); err != nil {
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

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt64).Eq(int64(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt64).Eq(int64(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt64).Eq(int64(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(256); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt64).Eq(int64(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(320); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotInt64).Eq(int64(5)); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckUint8(t *testing.T) {
	var param uint8 = 2
	v := Value{Name: "Test", Uint8: &param}
	raw := []byte{0x01, 0x02}

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint8).Eq(uint8(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint8).Eq(uint8(2)); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint8Range(t *testing.T) {
	var paramMin uint8 = 2
	var paramMax uint8 = 4
	v := Value{Name: "Test", MinUint8: &paramMin, MaxUint8: &paramMax}
	raw := []byte{0x01, 0x02, 0x03, 0x04, 0x05}

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(8); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint8).Eq(uint8(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint8).Eq(uint8(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint8).Eq(uint8(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint8).Eq(uint8(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(40); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint8).Eq(uint8(5)); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckUint16(t *testing.T) {
	var param uint16 = 2
	v := Value{Name: "Test", Uint16: &param}
	raw := []byte{0x00, 0x01, 0x00, 0x02}

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint16).Eq(uint16(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint16).Eq(uint16(2)); err != nil {
		t.Error(err)
	}
}

func TestValue_CheckUint16Range(t *testing.T) {
	var paramMin uint16 = 2
	var paramMax uint16 = 4
	v := Value{Name: "Test", MinUint16: &paramMin, MaxUint16: &paramMax}
	raw := []byte{
		0, 1,
		0, 2,
		0, 3,
		0, 4,
		0, 5,
	}

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(16); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint16).Eq(uint16(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint16).Eq(uint16(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint16).Eq(uint16(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint16).Eq(uint16(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(80); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint16).Eq(uint16(5)); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckUint32(t *testing.T) {
	var param uint32 = 2
	v := Value{Name: "Test", Uint32: &param}
	raw := []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02}

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint32).Eq(uint32(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint32).Eq(uint32(2)); err != nil {
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

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint32).Eq(uint32(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint32).Eq(uint32(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint32).Eq(uint32(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint32).Eq(uint32(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(160); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint32).Eq(uint32(5)); err != nil {
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

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint64).Eq(uint64(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint64).Eq(uint64(2)); err != nil {
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

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint64).Eq(uint64(1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint64).Eq(uint64(2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint64).Eq(uint64(3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(256); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint64).Eq(uint64(4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(320); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotUint64).Eq(uint64(5)); err != nil {
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

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat32).Eq(float32(2.3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat32).Eq(float32(2.4)); err != nil {
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

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(32); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat32).Eq(float32(2.1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat32).Eq(float32(2.2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(96); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat32).Eq(float32(2.3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat32).Eq(float32(2.4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(160); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat32).Eq(float32(2.5)); err != nil {
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

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat64).Eq(float64(2.3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat64).Eq(float64(2.4)); err != nil {
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

	offset := v.Check(raw, 0)
	if err := gotest.Expect(offset).Eq(64); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat64).Eq(float64(2.1)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(128); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat64).Eq(float64(2.2)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(192); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat64).Eq(float64(2.3)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(256); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat64).Eq(float64(2.4)); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(320); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotFloat64).Eq(float64(2.5)); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckBool(t *testing.T) {
	var param = true
	v := Value{Name: "Test", Bool: &param}
	raw := []byte{
		0b00000101,
	}
	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(1); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotBool).Eq(true); err != nil {
		t.Error(err)
	}

	param = false
	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(2); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotBool).Eq(false); err != nil {
		t.Error(err)
	}

	param = true
	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(3); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotBool).Eq(true); err != nil {
		t.Error(err)
	}

}

func TestValue_CheckString(t *testing.T) {
	var param string = "hello"
	v := Value{Name: "Test", Str: &param}
	raw := []byte{
		104, 101, 108, 108, 111,
		104, 101, 108, 109, 111,
	}

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(40); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotString).Eq("hello"); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(80); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotString).Eq("helmo"); err != nil {
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

	offset := v.Check(raw, 0)

	if err := gotest.Expect(offset).Eq(24); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(true); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotByte).Eq([]byte{1, 0, 3}); err != nil {
		t.Error(err)
	}

	offset = v.Check(raw, offset)
	if err := gotest.Expect(offset).Eq(48); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.Pass).Eq(false); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(v.GotByte).Eq([]byte{1, 0, 4}); err != nil {
		t.Error(err)
	}
}

func TestValue_Write(t *testing.T) {
	b, err := (&Value{}).Write()
	if err := gotest.Expect(b).Zero(); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Error("empty value"); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt8(t *testing.T) {
	var v int8 = 1
	b, err := (&Value{Int8: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt16(t *testing.T) {
	var v int16 = 1
	b, err := (&Value{Int16: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt32(t *testing.T) {
	var v int32 = 1
	b, err := (&Value{Int32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt64(t *testing.T) {
	var v int64 = 1
	b, err := (&Value{Int64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint8(t *testing.T) {
	var v uint8 = 1
	b, err := (&Value{Uint8: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint16(t *testing.T) {
	var v uint16 = 1
	b, err := (&Value{Uint16: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint32(t *testing.T) {
	var v uint32 = 1
	b, err := (&Value{Uint32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint64(t *testing.T) {
	var v uint64 = 1
	b, err := (&Value{Uint64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteFloat32(t *testing.T) {
	var v float32 = 1.2
	b, err := (&Value{Float32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{63, 153, 153, 154}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteFloat64(t *testing.T) {
	var v float64 = 1.2
	b, err := (&Value{Float64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{63, 243, 51, 51, 51, 51, 51, 51}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteBoolTrue(t *testing.T) {
	var v bool = true
	b, err := (&Value{Bool: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteBoolFalse(t *testing.T) {
	var v bool = false
	b, err := (&Value{Bool: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteString(t *testing.T) {
	var v string = "test"
	b, err := (&Value{Str: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{116, 101, 115, 116}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteByte(t *testing.T) {
	var v string = "0x01 0x0001 0x000002"
	b, err := (&Value{Byte: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1, 0, 1, 0, 0, 2}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Nil(); err != nil {
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
	if err := gotest.Expect((&Value{Str: &v}).Type()).Eq(String); err != nil {
		t.Error(err)
	}
}

func TestValue_TypeByte(t *testing.T) {
	var v string = "test"
	if err := gotest.Expect((&Value{Byte: &v}).Type()).Eq(Byte); err != nil {
		t.Error(err)
	}
}
