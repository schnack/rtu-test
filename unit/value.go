package unit

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
		b, err := v.parseStringByte(*v.Byte)
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

func (v *Value) parseStringByte(s string) ([]byte, error) {
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
