package unit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TypeValue int

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
}

func (v *Value) Check(raw []byte, currentBit int) (report Report, offsetBit int) {
	report.Name = v.Name
	report.Type = v.Type()
	report.Pass = true
	switch report.Type {
	case Nil:
		offsetBit = 0
	case Int8:
		report.Expected = []byte{uint8(*v.Int8)}
		offsetBit = currentBit + 8
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int8(report.Got[0])
		if got != *v.Int8 {
			report.Pass = false
		}
	case Int8Range:
		report.ExpectedMin = []byte{uint8(*v.MinInt8)}
		report.ExpectedMax = []byte{uint8(*v.MaxInt8)}
		offsetBit = currentBit + 8
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int8(report.Got[0])
		if *v.MinInt8 > got || got > *v.MaxInt8 {
			report.Pass = false
		}
	case Int16:
		report.Expected = make([]byte, 2)
		binary.BigEndian.PutUint16(report.Expected, uint16(*v.Int16))
		offsetBit = currentBit + 16
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int16(binary.BigEndian.Uint16(report.Got))
		if got != *v.Int16 {
			report.Pass = false
		}
	case Int16Range:
		report.ExpectedMin = make([]byte, 2)
		binary.BigEndian.PutUint16(report.ExpectedMin, uint16(*v.MinInt16))
		report.ExpectedMax = make([]byte, 2)
		binary.BigEndian.PutUint16(report.ExpectedMax, uint16(*v.MaxInt16))
		offsetBit = currentBit + 16
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int16(binary.BigEndian.Uint16(report.Got))
		if *v.MinInt16 > got || got > *v.MaxInt16 {
			report.Pass = false
		}
	case Int32:
		report.Expected = make([]byte, 4)
		binary.BigEndian.PutUint32(report.Expected, uint32(*v.Int32))
		offsetBit = currentBit + 32
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int32(binary.BigEndian.Uint32(report.Got))
		if got != *v.Int32 {
			report.Pass = false
		}
	case Int32Range:
		report.ExpectedMin = make([]byte, 4)
		binary.BigEndian.PutUint32(report.ExpectedMin, uint32(*v.MinInt32))
		report.ExpectedMax = make([]byte, 4)
		binary.BigEndian.PutUint32(report.ExpectedMax, uint32(*v.MaxInt32))
		offsetBit = currentBit + 32
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int32(binary.BigEndian.Uint32(report.Got))
		if *v.MinInt32 > got || got > *v.MaxInt32 {
			report.Pass = false
		}
	case Int64:
		report.Expected = make([]byte, 8)
		binary.BigEndian.PutUint64(report.Expected, uint64(*v.Int64))
		offsetBit = currentBit + 64
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int64(binary.BigEndian.Uint64(report.Got))
		if got != *v.Int64 {
			report.Pass = false
		}
	case Int64Range:
		report.ExpectedMin = make([]byte, 8)
		binary.BigEndian.PutUint64(report.ExpectedMin, uint64(*v.MinInt64))
		report.ExpectedMax = make([]byte, 8)
		binary.BigEndian.PutUint64(report.ExpectedMax, uint64(*v.MaxInt64))
		offsetBit = currentBit + 64
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := int64(binary.BigEndian.Uint64(report.Got))
		if *v.MinInt64 > got || got > *v.MaxInt64 {
			report.Pass = false
		}
	case Uint8:
		report.Expected = []byte{*v.Uint8}
		offsetBit = currentBit + 8
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := report.Got[0]
		if got != *v.Uint8 {
			report.Pass = false
		}
	case Uint8Range:
		report.ExpectedMin = []byte{*v.MinUint8}
		report.ExpectedMax = []byte{*v.MaxUint8}
		offsetBit = currentBit + 8
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := report.Got[0]
		if *v.MinUint8 > got || got > *v.MaxUint8 {
			report.Pass = false
		}
	case Uint16:
		report.Expected = make([]byte, 2)
		binary.BigEndian.PutUint16(report.Expected, *v.Uint16)
		offsetBit = currentBit + 16
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := binary.BigEndian.Uint16(report.Got)
		if got != *v.Uint16 {
			report.Pass = false
		}
	case Uint16Range:
		report.ExpectedMin = make([]byte, 2)
		binary.BigEndian.PutUint16(report.ExpectedMin, *v.MinUint16)
		report.ExpectedMax = make([]byte, 2)
		binary.BigEndian.PutUint16(report.ExpectedMax, *v.MaxUint16)
		offsetBit = currentBit + 16
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := binary.BigEndian.Uint16(report.Got)
		if *v.MinUint16 > got || got > *v.MaxUint16 {
			report.Pass = false
		}
	case Uint32:
		report.Expected = make([]byte, 4)
		binary.BigEndian.PutUint32(report.Expected, *v.Uint32)
		offsetBit = currentBit + 32
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := binary.BigEndian.Uint32(report.Got)
		if got != *v.Uint32 {
			report.Pass = false
		}
	case Uint32Range:
		report.ExpectedMin = make([]byte, 4)
		binary.BigEndian.PutUint32(report.ExpectedMin, *v.MinUint32)
		report.ExpectedMax = make([]byte, 4)
		binary.BigEndian.PutUint32(report.ExpectedMax, *v.MaxUint32)
		offsetBit = currentBit + 32
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := binary.BigEndian.Uint32(report.Got)
		if *v.MinUint32 > got || got > *v.MaxUint32 {
			report.Pass = false
		}
	case Uint64:
		report.Expected = make([]byte, 8)
		binary.BigEndian.PutUint64(report.Expected, *v.Uint64)
		offsetBit = currentBit + 64
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := binary.BigEndian.Uint64(report.Got)
		if got != *v.Uint64 {
			report.Pass = false
		}
	case Uint64Range:
		report.ExpectedMin = make([]byte, 8)
		binary.BigEndian.PutUint64(report.ExpectedMin, *v.MinUint64)
		report.ExpectedMax = make([]byte, 8)
		binary.BigEndian.PutUint64(report.ExpectedMax, *v.MaxUint64)
		offsetBit = currentBit + 64
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := binary.BigEndian.Uint64(report.Got)
		if *v.MinUint64 > got || got > *v.MaxUint64 {
			report.Pass = false
		}
	case Float32:
		report.Expected = make([]byte, 4)
		binary.BigEndian.PutUint32(report.Expected, math.Float32bits(*v.Float32))
		offsetBit = currentBit + 32
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := math.Float32frombits(binary.BigEndian.Uint32(report.Got))
		if got != *v.Float32 {
			report.Pass = false
		}
	case Float32Range:
		report.ExpectedMin = make([]byte, 4)
		binary.BigEndian.PutUint32(report.Expected, math.Float32bits(*v.MinFloat32))
		report.ExpectedMax = make([]byte, 4)
		binary.BigEndian.PutUint32(report.Expected, math.Float32bits(*v.MinFloat32))
		offsetBit = currentBit + 32
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := math.Float32frombits(binary.BigEndian.Uint32(report.Got))
		if *v.MinFloat32 > got || got > *v.MaxFloat32 {
			report.Pass = false
		}
	case Float64:
		report.Expected = make([]byte, 8)
		binary.BigEndian.PutUint64(report.Expected, math.Float64bits(*v.Float64))
		offsetBit = currentBit + 64
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := math.Float64frombits(binary.BigEndian.Uint64(report.Got))
		if got != *v.Float64 {
			report.Pass = false
		}
	case Float64Range:
		report.ExpectedMin = make([]byte, 8)
		binary.BigEndian.PutUint64(report.Expected, math.Float64bits(*v.MinFloat64))
		report.ExpectedMax = make([]byte, 8)
		binary.BigEndian.PutUint64(report.Expected, math.Float64bits(*v.MinFloat64))
		offsetBit = currentBit + 64
		report.Got = raw[currentBit/8 : offsetBit/8]
		got := math.Float64frombits(binary.BigEndian.Uint64(report.Got))
		if *v.MinFloat64 > got || got > *v.MaxFloat64 {
			report.Pass = false
		}
	case Bool:
		// TODO Если не полный байт то на других типах дополняем нулями
		if *v.Bool {
			report.Expected = []byte{1}
		} else {
			report.Expected = []byte{0}
		}
		got := raw[currentBit/8]&1<<currentBit%8 != 0
		offsetBit++

		if got {
			report.Got = []byte{1}
		} else {
			report.Got = []byte{0}
		}

		if got != *v.Bool {
			report.Pass = false
		}
	case String:
		// TODO
	case Byte:
		// TODO
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
	case v.MinInt8 != nil && v.MaxInt8 != nil:
		return Int8Range
	case v.Int16 != nil:
		return Int16
	case v.MinInt16 != nil && v.MaxInt16 != nil:
		return Int16Range
	case v.Int32 != nil:
		return Int32
	case v.MinInt32 != nil && v.MaxInt32 != nil:
		return Int32Range
	case v.Int64 != nil:
		return Int64
	case v.MinInt64 != nil && v.MaxInt64 != nil:
		return Int64Range
	case v.Uint8 != nil:
		return Uint8
	case v.MinUint8 != nil && v.MaxUint8 != nil:
		return Uint8Range
	case v.Uint16 != nil:
		return Uint16
	case v.MinUint16 != nil && v.MaxUint16 != nil:
		return Uint16Range
	case v.Uint32 != nil:
		return Uint32
	case v.MinUint32 != nil && v.MaxUint32 != nil:
		return Uint32Range
	case v.Uint64 != nil:
		return Uint64
	case v.MinUint64 != nil && v.MaxUint64 != nil:
		return Uint64Range
	case v.Float32 != nil:
		return Float32
	case v.MinFloat32 != nil && v.MaxFloat32 != nil:
		return Float32Range
	case v.Float64 != nil:
		return Float64
	case v.MinFloat64 != nil && v.MaxFloat64 != nil:
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

func (v *Value) GetByteWrite() ([]byte, error) {
	buf := new(bytes.Buffer)
	byteClear := strings.ReplaceAll(strings.ReplaceAll(*v.Byte, " ", ""), "0x", "")
	for i, _ := range byteClear {
		if i%2 != 0 {
			b, err := strconv.ParseUint(fmt.Sprintf("%c%c", byteClear[i-1], byteClear[i]), 16, 8)
			if err != nil {
				return nil, err
			}
			if err := binary.Write(buf, binary.BigEndian, uint8(b)); err != nil {
				return nil, err
			}
		}
	}
	return buf.Bytes(), nil
}
