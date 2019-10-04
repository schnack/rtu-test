package unit

import (
	"fmt"
	"github.com/schnack/gotest"
	"testing"
)

func TestValue_Write(t *testing.T) {
	b, err := (&Value{}).Write()
	if err := gotest.Expect(b).IsZero(); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(fmt.Errorf("empty value")); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt8(t *testing.T) {
	var v int8 = 1
	b, err := (&Value{Int8: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt16(t *testing.T) {
	var v int16 = 1
	b, err := (&Value{Int16: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt32(t *testing.T) {
	var v int32 = 1
	b, err := (&Value{Int32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteInt64(t *testing.T) {
	var v int64 = 1
	b, err := (&Value{Int64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint8(t *testing.T) {
	var v uint8 = 1
	b, err := (&Value{Uint8: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint16(t *testing.T) {
	var v uint16 = 1
	b, err := (&Value{Uint16: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint32(t *testing.T) {
	var v uint32 = 1
	b, err := (&Value{Uint32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteUint64(t *testing.T) {
	var v uint64 = 1
	b, err := (&Value{Uint64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0, 0, 0, 0, 0, 0, 1}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteFloat32(t *testing.T) {
	var v float32 = 1.2
	b, err := (&Value{Float32: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{63, 153, 153, 154}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteFloat64(t *testing.T) {
	var v float64 = 1.2
	b, err := (&Value{Float64: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{63, 243, 51, 51, 51, 51, 51, 51}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteBoolTrue(t *testing.T) {
	var v bool = true
	b, err := (&Value{Bool: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{255, 0}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteBoolFalse(t *testing.T) {
	var v bool = false
	b, err := (&Value{Bool: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{0, 0}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteString(t *testing.T) {
	var v string = "test"
	b, err := (&Value{String: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{116, 101, 115, 116}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}

func TestValue_WriteByte(t *testing.T) {
	var v string = "0x01 0x0001 0x000002"
	b, err := (&Value{Byte: &v}).Write()
	if err := gotest.Expect(b).Eq([]byte{1, 0, 1, 0, 0, 2}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
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
	var v string = "test"
	if err := gotest.Expect((&Value{Byte: &v}).Type()).Eq(Byte); err != nil {
		t.Error(err)
	}
}

func TestValue_parseStringByte(t *testing.T) {
	var v string = "0x01 0x0001 0x000002"
	b, err := (&Value{Byte: &v}).parseStringByte(v)
	if err := gotest.Expect(b).Eq([]byte{1, 0, 1, 0, 0, 2}); err != nil {
		t.Error(err)
	}
	if err := gotest.Expect(err).Eq(nil); err != nil {
		t.Error(err)
	}
}
