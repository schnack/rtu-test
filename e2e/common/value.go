package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"rtu-test/e2e/reports"
	"time"
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
	case Error:
		return "error"
	case Time:
		return "time"
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
	Error
	Time
)

type Value struct {
	Name string `yaml:"name"`
	// Используется для сервера
	Address string `yaml:"address"`

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

	// Особый Максимальное время выполнения
	Time *string `yaml:"time"`
	// Проверяем ошибку
	Error *string `yaml:"error"`
}

const FormatRange = "%s..%s"

// LengthBit - Длина значения в байтах
func (v *Value) LengthBit() int {
	switch v.Type() {
	case Bool:
		return 1
	case Int8, Int8Range, Uint8, Uint8Range:
		return 8
	case Int16, Int16Range, Uint16, Uint16Range:
		return 16
	case Int32, Int32Range, Uint32, Uint32Range, Float32, Float32Range:
		return 32
	case Int64, Int64Range, Uint64, Uint64Range, Float64, Float64Range:
		return 64
	case String:
		return len(v.Write(binary.BigEndian)) * 8
	case Byte:
		return len(v.Write(binary.BigEndian)) * 8
	default:
		return 0
	}
}

// cursor - Используется для отслеживания положения бита при проверки булевых значений
// currentBit - текущий бит в массиве байтов
// bitSize - размер значения
// basicBitSize - размер исходного значения
// startBit - индекс массива байтов
// endBit - конечный индекс массива байтов
// offsetBit - смещение в битах
func (v *Value) cursor(currentBit, bitSize, basicBitSize int, byteOrder binary.ByteOrder) (startBit, endBit, offsetBit int) {
	if currentBit%bitSize != 0 {
		currentBit += bitSize - (currentBit % bitSize)
	}

	offsetBit = currentBit + bitSize

	size := basicBitSize
	if bitSize > basicBitSize {
		size = bitSize
	}

	// корректировка подсчета данных для типа bool
	if bitSize == 1 {
		bitSize = 8
	}

	index := (currentBit % size) / bitSize
	if byteOrder == binary.BigEndian {
		index = ((size / bitSize) - 1) - index
	}
	endIndex := bitSize / 8

	startBit = (currentBit-(currentBit%size))/8 + index
	endBit = startBit + endIndex
	return
}

// Check - Осуществляет проверку данных
// rawBite - []byte полученные от устройства
// rawTime - время выполнения команды
// rawError - текст ошибки
// currentBit - курсор бита
// minBitSize - размер данных хранимых в табличке модбас 8 до 64 (дополняет нулями если тип например bool)
// orderByte - порядок байт
func (v *Value) Check(rawBite []byte, rawTime time.Duration, rawError string, currentBit int, minBitSize int, byteOrder binary.ByteOrder) (offsetBit int, report reports.ReportExpected) {
	report.Name = v.Name
	report.Pass = true
	report.Type = v.Type().String()
	switch v.Type() {
	case Nil:
		offsetBit = 0

	case Int8:
		report.Expected = fmt.Sprintf("%d", *v.Int8)
		report.ExpectedHex = fmt.Sprintf("%02x", *v.Int8)
		report.ExpectedBin = fmt.Sprintf("[%08b]", *v.Int8)

		start, _, offset := v.cursor(currentBit, 8, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int8(rawBite[start])

		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%02x", got)
		report.GotBin = fmt.Sprintf("[%08b]", got)

		report.Pass = got == *v.Int8

	case Int8Range:
		var min, minHex, minBin = "", "", ""
		if v.MinInt8 != nil {
			min = fmt.Sprintf("%d", *v.MinInt8)
			minHex = fmt.Sprintf("%02x", *v.MinInt8)
			minBin = fmt.Sprintf("[%08b]", *v.MinInt8)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxInt8 != nil {
			max = fmt.Sprintf("%d", *v.MaxInt8)
			maxHex = fmt.Sprintf("%02x", *v.MaxInt8)
			maxBin = fmt.Sprintf("[%08b]", *v.MaxInt8)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, _, offset := v.cursor(currentBit, 8, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int8(rawBite[start])

		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%02x", got)
		report.GotBin = fmt.Sprintf("[%08b]", got)

		report.Pass = !((v.MinInt8 != nil && *v.MinInt8 > got) || (v.MaxInt8 != nil && got > *v.MaxInt8))

	case Int16:
		report.Expected = fmt.Sprintf("%d", *v.Int16)
		b := make([]byte, 2)
		byteOrder.PutUint16(b, uint16(*v.Int16))
		report.ExpectedHex = fmt.Sprintf("%04x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 16, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int16(byteOrder.Uint16(rawBite[start:end]))
		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint16(b, uint16(got))
		report.GotHex = fmt.Sprintf("%04x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Int16

	case Int16Range:

		b := make([]byte, 2)
		var min, minHex, minBin = "", "", ""
		if v.MinInt16 != nil {
			min = fmt.Sprintf("%d", *v.MinInt16)
			byteOrder.PutUint16(b, uint16(*v.MinInt16))
			minHex = fmt.Sprintf("%04x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxInt16 != nil {
			max = fmt.Sprintf("%d", *v.MaxInt16)
			byteOrder.PutUint16(b, uint16(*v.MaxInt16))
			maxHex = fmt.Sprintf("%04x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 16, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int16(byteOrder.Uint16(rawBite[start:end]))

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint16(b, uint16(got))
		report.GotHex = fmt.Sprintf("%04x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinInt16 != nil && *v.MinInt16 > got) || (v.MaxInt16 != nil && got > *v.MaxInt16))

	case Int32:
		report.Expected = fmt.Sprintf("%d", *v.Int32)
		b := make([]byte, 4)
		byteOrder.PutUint32(b, uint32(*v.Int32))
		report.ExpectedHex = fmt.Sprintf("%08x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 32, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int32(byteOrder.Uint32(rawBite[start:end]))

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint32(b, uint32(got))
		report.GotHex = fmt.Sprintf("%08x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Int32

	case Int32Range:
		b := make([]byte, 4)
		var min, minHex, minBin = "", "", ""
		if v.MinInt32 != nil {
			min = fmt.Sprintf("%d", *v.MinInt32)
			byteOrder.PutUint32(b, uint32(*v.MinInt32))
			minHex = fmt.Sprintf("%08x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxInt32 != nil {
			max = fmt.Sprintf("%d", *v.MaxInt32)
			byteOrder.PutUint32(b, uint32(*v.MaxInt32))
			maxHex = fmt.Sprintf("%08x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 32, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int32(byteOrder.Uint32(rawBite[start:end]))
		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint32(b, uint32(got))
		report.GotHex = fmt.Sprintf("%08x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinInt32 != nil && *v.MinInt32 > got) || (v.MaxInt32 != nil && got > *v.MaxInt32))

	case Int64:
		report.Expected = fmt.Sprintf("%d", *v.Int64)
		b := make([]byte, 8)
		byteOrder.PutUint64(b, uint64(*v.Int64))
		report.ExpectedHex = fmt.Sprintf("%016x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 64, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int64(byteOrder.Uint64(rawBite[start:end]))

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint64(b, uint64(got))
		report.GotHex = fmt.Sprintf("%016x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Int64

	case Int64Range:
		b := make([]byte, 8)
		var min, minHex, minBin = "", "", ""
		if v.MinInt64 != nil {
			min = fmt.Sprintf("%d", *v.MinInt64)
			byteOrder.PutUint64(b, uint64(*v.MinInt64))
			minHex = fmt.Sprintf("%016x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxInt64 != nil {
			max = fmt.Sprintf("%d", *v.MaxInt64)
			byteOrder.PutUint64(b, uint64(*v.MaxInt64))
			maxHex = fmt.Sprintf("%016x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 64, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := int64(byteOrder.Uint64(rawBite[start:end]))
		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint64(b, uint64(got))
		report.GotHex = fmt.Sprintf("%016x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinInt64 != nil && *v.MinInt64 > got) || (v.MaxInt64 != nil && got > *v.MaxInt64))

	case Uint8:

		report.Expected = fmt.Sprintf("%d", *v.Uint8)
		report.ExpectedHex = fmt.Sprintf("%02x", *v.Uint8)
		report.ExpectedBin = fmt.Sprintf("[%08b]", *v.Uint8)

		start, _, offset := v.cursor(currentBit, 8, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := rawBite[start]

		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%02x", got)
		report.GotBin = fmt.Sprintf("[%08b]", got)

		report.Pass = got == *v.Uint8

	case Uint8Range:
		var min, minHex, minBin = "", "", ""
		if v.MinUint8 != nil {
			min = fmt.Sprintf("%d", *v.MinUint8)
			minHex = fmt.Sprintf("%02x", *v.MinUint8)
			minBin = fmt.Sprintf("[%08b]", *v.MinUint8)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxUint8 != nil {
			max = fmt.Sprintf("%d", *v.MaxUint8)
			maxHex = fmt.Sprintf("%02x", *v.MaxUint8)
			maxBin = fmt.Sprintf("[%08b]", *v.MaxUint8)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, _, offset := v.cursor(currentBit, 8, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := rawBite[start]
		report.Got = fmt.Sprintf("%d", got)
		report.GotHex = fmt.Sprintf("%02x", got)
		report.GotBin = fmt.Sprintf("[%08b]", got)

		report.Pass = !((v.MinUint8 != nil && *v.MinUint8 > got) || (v.MaxUint8 != nil && got > *v.MaxUint8))

	case Uint16:
		report.Expected = fmt.Sprintf("%d", *v.Uint16)
		b := make([]byte, 2)
		byteOrder.PutUint16(b, *v.Uint16)
		report.ExpectedHex = fmt.Sprintf("%04x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 16, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := byteOrder.Uint16(rawBite[start:end])

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint16(b, got)
		report.GotHex = fmt.Sprintf("%04x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Uint16

	case Uint16Range:
		b := make([]byte, 2)
		var min, minHex, minBin = "", "", ""
		if v.MinUint16 != nil {
			min = fmt.Sprintf("%d", *v.MinUint16)
			byteOrder.PutUint16(b, *v.MinUint16)
			minHex = fmt.Sprintf("%04x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxUint16 != nil {
			max = fmt.Sprintf("%d", *v.MaxUint16)
			byteOrder.PutUint16(b, *v.MaxUint16)
			maxHex = fmt.Sprintf("%04x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 16, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := byteOrder.Uint16(rawBite[start:end])
		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint16(b, got)
		report.GotHex = fmt.Sprintf("%04x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinUint16 != nil && *v.MinUint16 > got) || (v.MaxUint16 != nil && got > *v.MaxUint16))

	case Uint32:
		b := make([]byte, 4)
		byteOrder.PutUint32(b, *v.Uint32)
		report.Expected = fmt.Sprintf("%d", *v.Uint32)
		report.ExpectedHex = fmt.Sprintf("%08x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 32, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := byteOrder.Uint32(rawBite[start:end])

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint32(b, got)
		report.GotHex = fmt.Sprintf("%08x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Uint32

	case Uint32Range:
		b := make([]byte, 4)
		var min, minHex, minBin = "", "", ""
		if v.MinUint32 != nil {
			min = fmt.Sprintf("%d", *v.MinUint32)
			byteOrder.PutUint32(b, *v.MinUint32)
			minHex = fmt.Sprintf("%08x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxUint32 != nil {
			max = fmt.Sprintf("%d", *v.MaxUint32)
			byteOrder.PutUint32(b, *v.MaxUint32)
			maxHex = fmt.Sprintf("%08x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 32, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := byteOrder.Uint32(rawBite[start:end])

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint32(b, got)
		report.GotHex = fmt.Sprintf("%08x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinUint32 != nil && *v.MinUint32 > got) || (v.MaxUint32 != nil && got > *v.MaxUint32))

	case Uint64:
		report.Expected = fmt.Sprintf("%d", *v.Uint64)
		b := make([]byte, 8)
		byteOrder.PutUint64(b, *v.Uint64)
		report.ExpectedHex = fmt.Sprintf("%016x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 64, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := byteOrder.Uint64(rawBite[start:end])

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint64(b, got)
		report.GotHex = fmt.Sprintf("%016x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Uint64

	case Uint64Range:
		b := make([]byte, 8)
		var min, minHex, minBin = "", "", ""
		if v.MinUint64 != nil {
			min = fmt.Sprintf("%d", *v.MinUint64)
			byteOrder.PutUint64(b, *v.MinUint64)
			minHex = fmt.Sprintf("%016x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxUint64 != nil {
			max = fmt.Sprintf("%d", *v.MaxUint64)
			byteOrder.PutUint64(b, uint64(*v.MaxUint64))
			maxHex = fmt.Sprintf("%016x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 64, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := byteOrder.Uint64(rawBite[start:end])

		report.Got = fmt.Sprintf("%d", got)
		byteOrder.PutUint64(b, got)
		report.GotHex = fmt.Sprintf("%016x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinUint64 != nil && *v.MinUint64 > got) || (v.MaxUint64 != nil && got > *v.MaxUint64))

	case Float32:
		b := make([]byte, 4)
		byteOrder.PutUint32(b, math.Float32bits(*v.Float32))
		report.Expected = fmt.Sprintf("%f", *v.Float32)
		report.ExpectedHex = fmt.Sprintf("%08x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 32, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		gotbit := byteOrder.Uint32(rawBite[start:end])
		got := math.Float32frombits(gotbit)

		report.Got = fmt.Sprintf("%f", got)
		byteOrder.PutUint32(b, gotbit)
		report.GotHex = fmt.Sprintf("%08x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Float32

	case Float32Range:
		b := make([]byte, 4)
		var min, minHex, minBin = "", "", ""
		if v.MinFloat32 != nil {
			min = fmt.Sprintf("%f", *v.MinFloat32)
			byteOrder.PutUint32(b, math.Float32bits(*v.MinFloat32))
			minHex = fmt.Sprintf("%08x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxFloat32 != nil {
			max = fmt.Sprintf("%f", *v.MaxFloat32)
			byteOrder.PutUint32(b, math.Float32bits(*v.MaxFloat32))
			maxHex = fmt.Sprintf("%08x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 32, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		gotbit := byteOrder.Uint32(rawBite[start:end])
		got := math.Float32frombits(gotbit)

		report.Got = fmt.Sprintf("%f", got)
		byteOrder.PutUint32(b, gotbit)
		report.GotHex = fmt.Sprintf("%08x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinFloat32 != nil && *v.MinFloat32 > got) || (v.MaxFloat32 != nil && got > *v.MaxFloat32))

	case Float64:
		b := make([]byte, 8)
		byteOrder.PutUint64(b, math.Float64bits(*v.Float64))
		report.Expected = fmt.Sprintf("%f", *v.Float64)
		report.ExpectedHex = fmt.Sprintf("%016x", b)
		report.ExpectedBin = fmt.Sprintf("%08b", b)

		start, end, offset := v.cursor(currentBit, 64, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		gotbit := byteOrder.Uint64(rawBite[start:end])
		got := math.Float64frombits(gotbit)

		report.Got = fmt.Sprintf("%f", got)
		byteOrder.PutUint64(b, gotbit)
		report.GotHex = fmt.Sprintf("%016x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = got == *v.Float64

	case Float64Range:
		b := make([]byte, 8)
		var min, minHex, minBin = "", "", ""
		if v.MinFloat64 != nil {
			min = fmt.Sprintf("%f", *v.MinFloat64)
			byteOrder.PutUint64(b, math.Float64bits(*v.MinFloat64))
			minHex = fmt.Sprintf("%016x", b)
			minBin = fmt.Sprintf("%08b", b)
		}
		var max, maxHex, maxBin = "", "", ""
		if v.MaxFloat64 != nil {
			max = fmt.Sprintf("%f", *v.MaxFloat64)
			byteOrder.PutUint64(b, math.Float64bits(*v.MaxFloat64))
			maxHex = fmt.Sprintf("%016x", b)
			maxBin = fmt.Sprintf("%08b", b)
		}
		report.Expected = fmt.Sprintf(FormatRange, min, max)
		report.ExpectedHex = fmt.Sprintf(FormatRange, minHex, maxHex)
		report.ExpectedBin = fmt.Sprintf(FormatRange, minBin, maxBin)

		start, end, offset := v.cursor(currentBit, 64, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		gotbit := byteOrder.Uint64(rawBite[start:end])
		got := math.Float64frombits(gotbit)

		report.Got = fmt.Sprintf("%f", got)
		byteOrder.PutUint64(b, gotbit)
		report.GotHex = fmt.Sprintf("%016x", b)
		report.GotBin = fmt.Sprintf("%08b", b)

		report.Pass = !((v.MinFloat64 != nil && *v.MinFloat64 > got) || (v.MaxFloat64 != nil && got > *v.MaxFloat64))

	case Bool:

		report.Expected = fmt.Sprintf("%t", *v.Bool)
		if *v.Bool {
			report.ExpectedHex = fmt.Sprint("1")
			report.ExpectedBin = fmt.Sprint("1")
		} else {
			report.ExpectedHex = fmt.Sprint("0")
			report.ExpectedBin = fmt.Sprint("0")
		}

		start, _, offset := v.cursor(currentBit, 1, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := rawBite[start]&(1<<(currentBit%8)) != 0

		report.Got = fmt.Sprintf("%t", got)
		if got {
			report.GotHex = fmt.Sprint("1")
			report.GotBin = fmt.Sprint("1")
		} else {
			report.GotHex = fmt.Sprint("0")
			report.GotBin = fmt.Sprint("0")
		}

		report.Pass = got == *v.Bool

	case String:

		report.Expected = fmt.Sprintf("%s", *v.String)
		report.ExpectedHex = fmt.Sprintf("%x", []byte(*v.String))
		report.ExpectedBin = fmt.Sprintf("%b", []byte(*v.String))

		start, end, offset := v.cursor(currentBit, len(*v.String)*8, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := string(rawBite[start:end])

		report.Got = fmt.Sprintf("%s", got)
		report.GotHex = fmt.Sprintf("%x", rawBite[currentBit/8:offsetBit/8])
		report.GotBin = fmt.Sprintf("%b", rawBite[currentBit/8:offsetBit/8])

		report.Pass = got == *v.String

	case Byte:
		expected, err := ParseStringByte(*v.Byte)
		if err != nil {
			logrus.Fatal(err)
		}
		report.Expected = fmt.Sprintf("% x", expected)
		report.ExpectedHex = fmt.Sprintf("%02x", expected)
		report.ExpectedBin = fmt.Sprintf("%08b", expected)

		start, end, offset := v.cursor(currentBit, len(expected)*8, minBitSize, byteOrder)
		offsetBit = offset
		if len(rawBite) < offsetBit/8 {
			report.Pass = false
			return
		}

		got := rawBite[start:end]

		report.Got = fmt.Sprintf("% x", got)
		report.GotHex = fmt.Sprintf("%02x", got)
		report.GotBin = fmt.Sprintf("%08b", got)

		if len(got) != len(expected) {
			report.Pass = false
			return
		}

		for i := range expected {
			if got[i] != expected[i] {
				report.Pass = false
				return
			}
		}

	case Time:
		d := ParseDuration(*v.Time)
		if rawTime > d {
			report.Pass = false
		}
		report.Expected = d.String()
		report.ExpectedHex = fmt.Sprintf("%016x", d.Nanoseconds())
		report.ExpectedBin = fmt.Sprintf("[%064b]", d.Nanoseconds())
		report.Got = rawTime.String()
		report.GotHex = fmt.Sprintf("%016x", rawTime.Nanoseconds())
		report.GotBin = fmt.Sprintf("[%064b]", rawTime.Nanoseconds())

	case Error:
		if rawError != *v.Error {
			report.Pass = false
		}
		report.Expected = *v.Error
		report.ExpectedHex = fmt.Sprintf("%02x", []byte(*v.Error))
		report.ExpectedBin = fmt.Sprintf("%08b", []byte(*v.Error))
		report.Got = rawError
		report.GotHex = fmt.Sprintf("%02x", []byte(rawError))
		report.GotBin = fmt.Sprintf("%08b", []byte(rawError))
	}
	return
}

// ReportWrite - возвращает отчет о записанных данных
func (v *Value) ReportWrite(byteOrder binary.ByteOrder) reports.ReportWrite {
	report := reports.ReportWrite{Name: v.Name, Type: v.Type().String()}
	b := v.Write(byteOrder)
	switch v.Type() {
	case Int8:
		report.Data = fmt.Sprintf("%d", *v.Int8)
		report.DataHex = fmt.Sprintf("%02x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Int16:
		report.Data = fmt.Sprintf("%d", *v.Int16)
		report.DataHex = fmt.Sprintf("%04x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Int32:
		report.Data = fmt.Sprintf("%d", *v.Int32)
		report.DataHex = fmt.Sprintf("%08x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Int64:
		report.Data = fmt.Sprintf("%d", *v.Int64)
		report.DataHex = fmt.Sprintf("%016x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Uint8:
		report.Data = fmt.Sprintf("%d", *v.Uint8)
		report.DataHex = fmt.Sprintf("%02x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Uint16:
		report.Data = fmt.Sprintf("%d", *v.Uint16)
		report.DataHex = fmt.Sprintf("%04x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Uint32:
		report.Data = fmt.Sprintf("%d", *v.Uint32)
		report.DataHex = fmt.Sprintf("%08x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Uint64:
		report.Data = fmt.Sprintf("%d", *v.Uint64)
		report.DataHex = fmt.Sprintf("%016x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Float32:
		report.Data = fmt.Sprintf("%f", *v.Float32)
		report.DataHex = fmt.Sprintf("%08x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Float64:
		report.Data = fmt.Sprintf("%f", *v.Float64)
		report.DataHex = fmt.Sprintf("%016x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Bool:
		report.Data = fmt.Sprintf("%t", *v.Bool)
		report.DataHex = fmt.Sprintf("%02x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case String:
		report.Data = fmt.Sprintf("%s", *v.String)
		report.DataHex = fmt.Sprintf("%02x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	case Byte:
		b, err := ParseStringByte(*v.Byte)
		if err != nil {
			logrus.Fatal(err)
		}
		report.Data = fmt.Sprintf("% x", b)
		report.DataHex = fmt.Sprintf("%02x", b)
		report.DataBin = fmt.Sprintf("%08b", b)

	default:
		logrus.Fatal("empty value")
	}
	return report
}

// Write - Возвращает значение в байтах
func (v *Value) Write(byteOrder binary.ByteOrder) (b []byte) {

	buf := new(bytes.Buffer)
	switch v.Type() {
	case Int8:
		if err := binary.Write(buf, byteOrder, v.Int8); err != nil {
			logrus.Fatal(err)
		}
	case Int16:
		if err := binary.Write(buf, byteOrder, v.Int16); err != nil {
			logrus.Fatal(err)
		}
	case Int32:
		if err := binary.Write(buf, byteOrder, v.Int32); err != nil {
			logrus.Fatal(err)
		}
	case Int64:
		if err := binary.Write(buf, byteOrder, v.Int64); err != nil {
			logrus.Fatal(err)
		}
	case Uint8:
		if err := binary.Write(buf, byteOrder, v.Uint8); err != nil {
			logrus.Fatal(err)
		}

	case Uint16:
		if err := binary.Write(buf, byteOrder, v.Uint16); err != nil {
			logrus.Fatal(err)
		}

	case Uint32:
		if err := binary.Write(buf, byteOrder, v.Uint32); err != nil {
			logrus.Fatal(err)
		}

	case Uint64:
		if err := binary.Write(buf, byteOrder, v.Uint64); err != nil {
			logrus.Fatal(err)
		}

	case Float32:
		if err := binary.Write(buf, byteOrder, v.Float32); err != nil {
			logrus.Fatal(err)
		}

	case Float64:
		if err := binary.Write(buf, byteOrder, v.Float64); err != nil {
			logrus.Fatal(err)
		}

	case Bool:
		if err := binary.Write(buf, byteOrder, v.Bool); err != nil {
			logrus.Fatal(err)
		}

	case String:
		buf.WriteString(*v.String)
	case Byte:
		b, err := ParseStringByte(*v.Byte)
		if err != nil {
			logrus.Fatal(err)
		}
		buf.Write(b)

	}
	return buf.Bytes()
}

/// Type - возвращает тип текущего значения
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
	case v.Error != nil:
		return Error
	case v.Time != nil:
		return Time
	default:
		return Nil
	}
}
