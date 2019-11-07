package unit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

type TypeValue int

func (tv TypeValue) String() string {
	switch tv {
	case Int8, Int8Range:
		return "int8"
	case Int16, Int16Range:
		return "int16"
	case Int32, Int32Range:
		return "int32"
	case Int64, Int64Range:
		return "int64"
	case Uint8, Uint8Range:
		return "uint8"
	case Uint16, Uint16Range:
		return "uint16"
	case Uint32, Uint32Range:
		return "uint32"
	case Uint64, Uint64Range:
		return "uint64"
	case Float32, Float32Range:
		return "float32"
	case Float64, Float64Range:
		return "float64"
	case Bool:
		return "bool"
	case String:
		return "string"
	case Byte:
		return "byte"
	default:
		return "nil"
	}
}

const (
	Nil = TypeValue(iota)
	Int8
	Int8Range
	Int16
	Int16Range
	Int32
	Int32Range
	Int64
	Int64Range
	Uint8
	Uint8Range
	Uint16
	Uint16Range
	Uint32
	Uint32Range
	Uint64
	Uint64Range
	Float32
	Float32Range
	Float64
	Float64Range
	Bool
	String
	Byte
)

type Value struct {
	Name string `yaml:"name"`

	Int8    *int8 `yaml:"int8"`
	MaxInt8 *int8 `yaml:"maxInt8"`
	MinInt8 *int8 `yaml:"minInt8"`

	Int16    *int16 `yaml:"int16"`
	MaxInt16 *int16 `yaml:"maxInt16"`
	MinInt16 *int16 `yaml:"minInt16"`

	Int32    *int32 `yaml:"int32"`
	MaxInt32 *int32 `yaml:"maxInt32"`
	MinInt32 *int32 `yaml:"minInt32"`

	Int64    *int64 `yaml:"int64"`
	MaxInt64 *int64 `yaml:"maxInt64"`
	MinInt64 *int64 `yaml:"minInt64"`

	Uint8    *uint8 `yaml:"uint8"`
	MaxUint8 *uint8 `yaml:"maxUint8"`
	MinUint8 *uint8 `yaml:"minUint8"`

	Uint16    *uint16 `yaml:"uint16"`
	MaxUint16 *uint16 `yaml:"maxUint16"`
	MinUint16 *uint16 `yaml:"minUint16"`

	Uint32    *uint32 `yaml:"uint32"`
	MaxUint32 *uint32 `yaml:"maxUint32"`
	MinUint32 *uint32 `yaml:"minUint32"`

	Uint64    *uint64 `yaml:"uint64"`
	MaxUint64 *uint64 `yaml:"maxUint64"`
	MinUint64 *uint64 `yaml:"minUint64"`

	Float32    *float32 `yaml:"float32"`
	MaxFloat32 *float32 `yaml:"maxFloat32"`
	MinFloat32 *float32 `yaml:"minFloat32"`

	Float64    *float64 `yaml:"float64"`
	MaxFloat64 *float64 `yaml:"maxFloat64"`
	MinFloat64 *float64 `yaml:"minFloat64"`

	Bool *bool `yaml:"bool"`

	String *string `yaml:"string"`

	Byte *string `yaml:"byte"`

	GotInt8    *int8    `yaml:"-"`
	GotInt16   *int16   `yaml:"-"`
	GotInt32   *int32   `yaml:"-"`
	GotInt64   *int64   `yaml:"-"`
	GotUint8   *uint8   `yaml:"-"`
	GotUint16  *uint16  `yaml:"-"`
	GotUint32  *uint32  `yaml:"-"`
	GotUint64  *uint64  `yaml:"-"`
	GotFloat32 *float32 `yaml:"-"`
	GotFloat64 *float64 `yaml:"-"`
	GotBool    *bool    `yaml:"-"`
	GotString  *string  `yaml:"-"`
	GotByte    []byte   `yaml:"-"`

	Pass bool `yaml:"-"`
}

const FormatRange = "%s..%s"
const FormatDecimal = "(%[1]s) %[2]d"
const FormatDecimalRange = "(%[1]s) %[2]d..%[3]d"
const FormatDecimalMore = "> (%[1]s) %[2]d"
const FormatDecimalLess = "< (%[1]s)%[2]d"
const FormatFloat = "(%[1]s) %[2]f"
const FormatFloatRange = "(%[1]s) %[2]f..%[3]f"
const FormatFloatMore = "> (%[1]s) %[2]f"
const FormatFloatLess = "< (%[1]s) %[2]f"
const FormatString = "(%[1]s) %[2]s"
const FormatByte = "(%[1]s) {0x%02[2]x} [b%08[2]b]"
const FormatBool = "(%[1]s) %[2]t"
const FormatNil = "(nil)"

func (v *Value) StringGot() string {
	switch v.Type() {
	case Int8, Int8Range:
		if v.GotInt8 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatDecimal, Int8, *v.GotInt8)
	case Int16, Int16Range:
		if v.GotInt16 == nil {
			return FormatNil
		}
		return fmt.Sprintf(fmt.Sprintf(FormatDecimal, Int16, *v.GotInt16))
	case Int32, Int32Range:
		if v.GotInt32 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatDecimal, Int32, *v.GotInt32)
	case Int64, Int64Range:
		if v.GotInt64 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatDecimal, Int64, *v.GotInt64)
	case Uint8, Uint8Range:
		if v.GotUint8 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatDecimal, Uint8, *v.GotUint8)
	case Uint16, Uint16Range:
		if v.GotUint16 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatDecimal, Uint16, *v.GotUint16)
	case Uint32, Uint32Range:
		if v.GotUint32 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatDecimal, Uint32, *v.GotUint32)
	case Uint64, Uint64Range:
		if v.GotUint64 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatDecimal, Uint64, *v.GotUint64)
	case Float32, Float32Range:
		if v.GotFloat32 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatFloat, Float32, *v.GotFloat32)
	case Float64, Float64Range:
		if v.GotFloat64 == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatFloat, Float64, *v.GotFloat64)
	case Bool:
		if v.GotBool == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatBool, Bool, *v.GotBool)
	case String:
		if v.GotString == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatString, String, *v.GotString)
	case Byte:
		if v.GotByte == nil {
			return FormatNil
		}
		return fmt.Sprintf(FormatByte, Byte, v.GotByte)
	default:
		return FormatNil
	}
}

func (v *Value) StringExpected() string {
	switch v.Type() {
	case Int8:
		return fmt.Sprintf(FormatDecimal, Int8, *v.Int8)
	case Int8Range:
		if v.MinInt8 == nil {
			return fmt.Sprintf(FormatDecimalLess, Int8, *v.MaxInt8)
		} else if v.MaxInt8 == nil {
			return fmt.Sprintf(FormatDecimalMore, Int8, *v.MinInt8)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Int8, *v.MinInt8, *v.MaxInt8)
		}
	case Int16:
		return fmt.Sprintf(fmt.Sprintf(FormatDecimal, Int16, *v.Int16))
	case Int16Range:
		if v.MinInt16 == nil {
			return fmt.Sprintf(FormatDecimalLess, Int16, *v.MaxInt16)
		} else if v.MaxInt16 == nil {
			return fmt.Sprintf(FormatDecimalMore, Int16, *v.MinInt16)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Int16, *v.MinInt16, *v.MaxInt16)
		}
	case Int32:
		return fmt.Sprintf(FormatDecimal, Int32, *v.Int32)
	case Int32Range:
		if v.MinInt32 == nil {
			return fmt.Sprintf(FormatDecimalLess, Int32, *v.MaxInt32)
		} else if v.MaxInt32 == nil {
			return fmt.Sprintf(FormatDecimalMore, Int32, *v.MinInt32)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Int32, *v.MinInt32, *v.MaxInt32)
		}
	case Int64:
		return fmt.Sprintf(FormatDecimal, Int64, *v.Int64)
	case Int64Range:
		if v.MinInt64 == nil {
			return fmt.Sprintf(FormatDecimalLess, Int64, *v.MaxInt64)
		} else if v.MaxInt64 == nil {
			return fmt.Sprintf(FormatDecimalMore, Int64, *v.MinInt64)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Int64, *v.MinInt64, *v.MaxInt64)
		}
	case Uint8:
		return fmt.Sprintf(FormatDecimal, Uint8, *v.Uint8)
	case Uint8Range:
		if v.MinUint8 == nil {
			return fmt.Sprintf(FormatDecimalLess, Uint8, *v.MaxUint8)
		} else if v.MaxUint8 == nil {
			return fmt.Sprintf(FormatDecimalMore, Uint8, *v.MinUint8)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Uint8, *v.MinUint8, *v.MaxUint8)
		}
	case Uint16:
		return fmt.Sprintf(FormatDecimal, Uint16, *v.Uint16)
	case Uint16Range:
		if v.MinUint16 == nil {
			return fmt.Sprintf(FormatDecimalLess, Uint16, *v.MaxUint16)
		} else if v.MaxUint16 == nil {
			return fmt.Sprintf(FormatDecimalMore, Uint16, *v.MinUint16)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Uint16, *v.MinUint16, *v.MaxUint16)
		}
	case Uint32:
		return fmt.Sprintf(FormatDecimal, Uint32, *v.Uint32)
	case Uint32Range:
		if v.MinUint32 == nil {
			return fmt.Sprintf(FormatDecimalLess, Uint32, *v.MaxUint32)
		} else if v.MaxUint32 == nil {
			return fmt.Sprintf(FormatDecimalMore, Uint32, *v.MinUint32)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Uint32, *v.MinUint32, *v.MaxUint32)
		}
	case Uint64:
		return fmt.Sprintf(FormatDecimal, Uint64, *v.Uint64)
	case Uint64Range:
		if v.MinUint64 == nil {
			return fmt.Sprintf(FormatDecimalLess, Uint64, *v.MaxUint64)
		} else if v.MaxUint64 == nil {
			return fmt.Sprintf(FormatDecimalMore, Uint64, *v.MinUint64)
		} else {
			return fmt.Sprintf(FormatDecimalRange, Uint64, *v.MinUint64, *v.MaxUint64)
		}
	case Float32:
		return fmt.Sprintf(FormatFloat, Float32, *v.Float32)
	case Float32Range:
		if v.MinFloat32 == nil {
			return fmt.Sprintf(FormatFloatLess, Float32, *v.MaxFloat32)
		} else if v.MaxFloat32 == nil {
			return fmt.Sprintf(FormatFloatMore, Float32, *v.MinFloat32)
		} else {
			return fmt.Sprintf(FormatFloatRange, Float32, *v.MinFloat32, *v.MaxFloat32)
		}
	case Float64:
		return fmt.Sprintf(FormatFloat, Float64, *v.Float64)
	case Float64Range:
		if v.MinFloat64 == nil {
			return fmt.Sprintf(FormatFloatLess, Float64, *v.MaxFloat64)
		} else if v.MaxFloat64 == nil {
			return fmt.Sprintf(FormatFloatMore, Float64, *v.MinFloat64)
		} else {
			return fmt.Sprintf(FormatFloatRange, Float64, *v.MinFloat64, *v.MaxFloat64)
		}
	case Bool:
		return fmt.Sprintf(FormatBool, Bool, *v.Bool)
	case String:
		return fmt.Sprintf(FormatString, String, *v.String)
	case Byte:
		expected, err := parseStringByte(*v.Byte)
		if err != nil {
			logrus.Fatal(err)
		}
		return fmt.Sprintf(FormatByte, Byte, expected)
	default:
		return ""
	}
}

func (v *Value) Check(raw []byte, currentBit int) (offsetBit int, report ReportExpected) {
	report.Name = v.Name
	report.Pass = true
	report.Type = v.Type().String()
	switch v.Type() {
	case Nil:
		offsetBit = 0
	case Int8:
		report.Expected = fmt.Sprintf("%d", *v.Int8)
		report.ExpectedHex = fmt.Sprintf("%02x", *v.Int8)
		report.ExpectedBin = fmt.Sprintf("%08b", *v.Int8)

		offsetBit = currentBit + (currentBit % 8) + 8
		if len(raw) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int8(raw[currentBit/8 : offsetBit/8][0])
		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%02x", got)
		report.GotBin = fmt.Sprintf("%08b", got)

		report.Pass = got != *v.Int8

	case Int8Range:
		var min, minHex, minBin = "", "", ""
		if v.MinInt8 != nil {
			min = fmt.Sprintf("%d", *v.MinInt8)
			minHex = fmt.Sprintf("%02x", *v.MinInt8)
			minBin = fmt.Sprintf("%08b", *v.MinInt8)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MinInt8 != nil {
			max = fmt.Sprintf("%d", *v.MaxInt8)
			maxHex = fmt.Sprintf("%02x", *v.MaxInt8)
			maxBin = fmt.Sprintf("%08b", *v.MaxInt8)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		offsetBit = currentBit + (currentBit % 8) + 8
		if len(raw) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int8(raw[currentBit/8 : offsetBit/8][0])
		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%02x", got)
		report.GotBin = fmt.Sprintf("%08b", got)

		report.Pass = !((v.MinInt8 != nil && *v.MinInt8 > got) || (v.MaxInt8 != nil && got > *v.MaxInt8))

	case Int16:
		report.Expected = fmt.Sprintf("%d", *v.Int16)
		report.ExpectedHex = fmt.Sprintf("%04x", *v.Int16)
		report.ExpectedBin = fmt.Sprintf("%016b", *v.Int16)

		offsetBit = currentBit + (currentBit % 8) + 16
		if len(raw) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int16(binary.BigEndian.Uint16(raw[currentBit/8 : offsetBit/8]))
		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%04x", got)
		report.GotBin = fmt.Sprintf("%016b", got)

		report.Pass = got != *v.Int16

	case Int16Range:

		var min, minHex, minBin = "", "", ""
		if v.MinInt16 != nil {
			min = fmt.Sprintf("%d", *v.MinInt16)
			minHex = fmt.Sprintf("%04x", *v.MinInt16)
			minBin = fmt.Sprintf("%016b", *v.MinInt16)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MinInt16 != nil {
			max = fmt.Sprintf("%d", *v.MaxInt16)
			maxHex = fmt.Sprintf("%04x", *v.MaxInt16)
			maxBin = fmt.Sprintf("%016b", *v.MaxInt16)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		offsetBit = currentBit + (currentBit % 8) + 16
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := int16(binary.BigEndian.Uint16(raw[currentBit/8 : offsetBit/8]))
		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%04x", got)
		report.GotBin = fmt.Sprintf("%016b", got)
		v.GotInt16 = &got

		report.Pass = !((v.MinInt16 != nil && *v.MinInt16 > got) || (v.MaxInt16 != nil && got > *v.MaxInt16))

	case Int32:

		offsetBit = currentBit + (currentBit % 8) + 32
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := int32(binary.BigEndian.Uint32(raw[currentBit/8 : offsetBit/8]))
		v.GotInt32 = &got

		if *v.GotInt32 != *v.Int32 {
			v.Pass = false
		}
	case Int32Range:

		offsetBit = currentBit + (currentBit % 8) + 32
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := int32(binary.BigEndian.Uint32(raw[currentBit/8 : offsetBit/8]))
		v.GotInt32 = &got

		if v.MinInt32 != nil && *v.MinInt32 > *v.GotInt32 {
			v.Pass = false
		}

		if v.MaxInt32 != nil && *v.GotInt32 > *v.MaxInt32 {
			v.Pass = false
		}
	case Int64:

		offsetBit = currentBit + (currentBit % 8) + 64
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := int64(binary.BigEndian.Uint64(raw[currentBit/8 : offsetBit/8]))
		v.GotInt64 = &got

		if *v.GotInt64 != *v.Int64 {
			v.Pass = false
		}
	case Int64Range:

		offsetBit = currentBit + (currentBit % 8) + 64
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := int64(binary.BigEndian.Uint64(raw[currentBit/8 : offsetBit/8]))
		v.GotInt64 = &got

		if v.MinInt64 != nil && *v.MinInt64 > *v.GotInt64 {
			v.Pass = false
		}

		if v.MaxInt64 != nil && *v.GotInt64 > *v.MaxInt64 {
			v.Pass = false
		}
	case Uint8:

		offsetBit = currentBit + (currentBit % 8) + 8
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := raw[currentBit/8 : offsetBit/8][0]
		v.GotUint8 = &got

		if *v.GotUint8 != *v.Uint8 {
			v.Pass = false
		}
	case Uint8Range:

		offsetBit = currentBit + (currentBit % 8) + 8
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := raw[currentBit/8 : offsetBit/8][0]
		v.GotUint8 = &got

		if v.MinUint8 != nil && *v.MinUint8 > *v.GotUint8 {
			v.Pass = false
		}

		if v.MaxUint8 != nil && *v.GotUint8 > *v.MaxUint8 {
			v.Pass = false
		}
	case Uint16:

		offsetBit = currentBit + (currentBit % 8) + 16
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := binary.BigEndian.Uint16(raw[currentBit/8 : offsetBit/8])
		v.GotUint16 = &got

		if *v.GotUint16 != *v.Uint16 {
			v.Pass = false
		}
	case Uint16Range:

		offsetBit = currentBit + (currentBit % 8) + 16
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := binary.BigEndian.Uint16(raw[currentBit/8 : offsetBit/8])
		v.GotUint16 = &got

		if v.MinUint16 != nil && *v.MinUint16 > *v.GotUint16 {
			v.Pass = false
		}

		if v.MaxUint16 != nil && *v.GotUint16 > *v.MaxUint16 {
			v.Pass = false
		}
	case Uint32:

		offsetBit = currentBit + (currentBit % 8) + 32
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := binary.BigEndian.Uint32(raw[currentBit/8 : offsetBit/8])
		v.GotUint32 = &got

		if *v.GotUint32 != *v.Uint32 {
			v.Pass = false
		}
	case Uint32Range:

		offsetBit = currentBit + (currentBit % 8) + 32
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := binary.BigEndian.Uint32(raw[currentBit/8 : offsetBit/8])
		v.GotUint32 = &got

		if v.MinUint32 != nil && *v.MinUint32 > *v.GotUint32 {
			v.Pass = false
		}

		if v.MaxUint32 != nil && *v.GotUint32 > *v.MaxUint32 {
			v.Pass = false
		}
	case Uint64:

		offsetBit = currentBit + (currentBit % 8) + 64
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := binary.BigEndian.Uint64(raw[currentBit/8 : offsetBit/8])
		v.GotUint64 = &got

		if *v.GotUint64 != *v.Uint64 {
			v.Pass = false
		}
	case Uint64Range:

		offsetBit = currentBit + (currentBit % 8) + 64
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := binary.BigEndian.Uint64(raw[currentBit/8 : offsetBit/8])
		v.GotUint64 = &got

		if v.MinUint64 != nil && *v.MinUint64 > *v.GotUint64 {
			v.Pass = false
		}

		if v.MaxUint64 != nil && *v.GotUint64 > *v.MaxUint64 {
			v.Pass = false
		}
	case Float32:

		offsetBit = currentBit + (currentBit % 8) + 32
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := math.Float32frombits(binary.BigEndian.Uint32(raw[currentBit/8 : offsetBit/8]))
		v.GotFloat32 = &got

		if *v.GotFloat32 != *v.Float32 {
			v.Pass = false
		}
	case Float32Range:

		offsetBit = currentBit + (currentBit % 8) + 32
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := math.Float32frombits(binary.BigEndian.Uint32(raw[currentBit/8 : offsetBit/8]))
		v.GotFloat32 = &got

		if v.MinFloat32 != nil && *v.MinFloat32 > *v.GotFloat32 {
			v.Pass = false
		}

		if v.MaxFloat32 != nil && *v.GotFloat32 > *v.MaxFloat32 {
			v.Pass = false
		}

	case Float64:

		offsetBit = currentBit + (currentBit % 8) + 64
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := math.Float64frombits(binary.BigEndian.Uint64(raw[currentBit/8 : offsetBit/8]))
		v.GotFloat64 = &got

		if *v.GotFloat64 != *v.Float64 {
			v.Pass = false
		}
	case Float64Range:

		offsetBit = currentBit + (currentBit % 8) + 64
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := math.Float64frombits(binary.BigEndian.Uint64(raw[currentBit/8 : offsetBit/8]))
		v.GotFloat64 = &got

		if v.MinFloat64 != nil && *v.MinFloat64 > *v.GotFloat64 {
			v.Pass = false
		}

		if v.MaxFloat64 != nil && *v.GotFloat64 > *v.MaxFloat64 {
			v.Pass = false
		}
	case Bool:

		offsetBit = currentBit + 1
		if len(raw)*8 < offsetBit {
			v.Pass = false
			return
		}

		got := raw[currentBit/8]&(1<<(currentBit%8)) != 0
		v.GotBool = &got

		if *v.GotBool != *v.Bool {
			v.Pass = false
		}
	case String:

		offsetBit = currentBit + (currentBit % 8) + (len(*v.String) * 8)
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		got := string(raw[currentBit/8 : offsetBit/8])
		v.GotString = &got

		if *v.GotString != *v.String {
			v.Pass = false
		}
	case Byte:
		expected, err := parseStringByte(*v.Byte)

		if err != nil {
			logrus.Fatal(err)
		}

		offsetBit = currentBit + (currentBit % 8) + (len(expected) * 8)
		if len(raw) < offsetBit/8 {
			v.Pass = false
			return
		}

		v.GotByte = raw[currentBit/8 : offsetBit/8]

		if len(v.GotByte) != len(expected) {
			v.Pass = false
			return
		}

		for i, _ := range expected {
			if v.GotByte[i] != expected[i] {
				v.Pass = false
				return
			}
		}
	}
	return
}

func (v *Value) Write() ([]byte, error) {
	buf := new(bytes.Buffer)
	switch {
	case v.Int8 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Int8); err != nil {
			return nil, err
		}
	case v.Int16 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Int16); err != nil {
			return nil, err
		}
	case v.Int32 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Int32); err != nil {
			return nil, err
		}
	case v.Int64 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Int64); err != nil {
			return nil, err
		}
	case v.Uint8 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Uint8); err != nil {
			return nil, err
		}
	case v.Uint16 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Uint16); err != nil {
			return nil, err
		}
	case v.Uint32 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Uint32); err != nil {
			return nil, err
		}
	case v.Uint64 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Uint64); err != nil {
			return nil, err
		}
	case v.Float32 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Float32); err != nil {
			return nil, err
		}
	case v.Float64 != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Float64); err != nil {
			return nil, err
		}
	case v.Bool != nil:
		if err := binary.Write(buf, binary.BigEndian, v.Bool); err != nil {
			return nil, err
		}
	case v.String != nil:
		buf.WriteString(*v.String)
	case v.Byte != nil:
		b, err := parseStringByte(*v.Byte)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	default:
		return nil, fmt.Errorf("empty value")
	}
	return buf.Bytes(), nil
}

func (v *Value) Type() TypeValue {
	switch {
	case v.Int8 != nil:
		return Int8
	case v.MinInt8 != nil || v.MaxInt8 != nil:
		return Int8Range
	case v.Int16 != nil:
		return Int16
	case v.MinInt16 != nil || v.MaxInt16 != nil:
		return Int16Range
	case v.Int32 != nil:
		return Int32
	case v.MinInt32 != nil || v.MaxInt32 != nil:
		return Int32Range
	case v.Int64 != nil:
		return Int64
	case v.MinInt64 != nil || v.MaxInt64 != nil:
		return Int64Range
	case v.Uint8 != nil:
		return Uint8
	case v.MinUint8 != nil || v.MaxUint8 != nil:
		return Uint8Range
	case v.Uint16 != nil:
		return Uint16
	case v.MinUint16 != nil || v.MaxUint16 != nil:
		return Uint16Range
	case v.Uint32 != nil:
		return Uint32
	case v.MinUint32 != nil || v.MaxUint32 != nil:
		return Uint32Range
	case v.Uint64 != nil:
		return Uint64
	case v.MinUint64 != nil || v.MaxUint64 != nil:
		return Uint64Range
	case v.Float32 != nil:
		return Float32
	case v.MinFloat32 != nil || v.MaxFloat32 != nil:
		return Float32Range
	case v.Float64 != nil:
		return Float64
	case v.MinFloat64 != nil || v.MaxFloat64 != nil:
		return Float64Range
	case v.Bool != nil:
		return Bool
	case v.String != nil:
		return String
	case v.Byte != nil:
		return Byte
	default:
		return Nil
	}
}
