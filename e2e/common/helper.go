package common

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Render(tmpl string, data interface{}) string {
	t := template.Must(template.New("message").Parse(tmpl))
	buff := new(bytes.Buffer)
	if err := t.Execute(buff, data); err != nil {
		logrus.Fatal(err)
	}
	return buff.String()
}

// ParseStringByte - Превращает текстовое представления байт в настоящие байты
func ParseStringByte(sb string) ([]byte, error) {
	buf := new(bytes.Buffer)
	byteClear := strings.ReplaceAll(strings.ReplaceAll(sb, " ", ""), "0x", "")
	for i := range byteClear {
		if i%2 != 0 {
			b, err := strconv.ParseUint(byteClear[i-1:i+1], 16, 8)
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

func DataSingleCoil(data []byte) []byte {
	if len(data) > 1 && (data[0] == 0xff || data[0] == 0x00) && data[1] == 0x00 {
		return data[:2]
	} else if len(data) > 0 && data[0] == 0x01 || data[0] == 0x00 {
		if data[0] == 0x01 {
			return []byte{0xff, 0x00}
		} else {
			return []byte{0x00, 0x00}
		}
	} else {
		return nil
	}
}

func CountBit(v []*Value, is16bit bool) (bits uint16) {
	for _, w := range v {
		if w.Type() == Bool {
			bits++
		} else {
			if bits%8 != 0 {
				bits += 8 - bits%8
			}
			byteData := w.Write(binary.BigEndian)
			bits += 8 * uint16(len(byteData))
		}
	}
	if is16bit {
		if bits%16 != 0 {
			bits += 16 - bits%16
		}
		return bits / 16
	} else {
		return bits
	}
}

func ValueToByte(v []*Value) (data []byte) {
	var i int
	var vByte uint8
	for _, w := range v {
		switch w.Type() {
		case Bool:
			if *w.Bool {
				vByte = vByte | 1<<i
			}
			i++
			if i > 7 {
				data = append(data, vByte)
				vByte = 0
				i = 0
			}
		default:
			if i != 0 {
				data = append(data, vByte)
				vByte = 0
				i = 0
			}
			data = append(data, w.Write(binary.BigEndian)...)
		}
	}
	if i != 0 {
		data = append(data, vByte)
		vByte = 0
		i = 0
	}
	return
}

// Есть особенность если указан в конфигурации
//  int8 будет дополнен []{byte{int8, 0}} LittleEndian
//  int16 #TODO Доделат!!!
func ValueToByte16(v []*Value) (data []byte) {
	var i int
	var vByte uint8
	for _, w := range v {
		switch w.Type() {
		case Bool:
			if *w.Bool {
				vByte = vByte | 1<<i
			}
			i++
			if i > 7 {
				data = append(data, vByte)
				vByte = 0
				i = 0
			}
		case Int8, Uint8:
			if i != 0 {
				data = append(data, vByte)
				vByte = 0
				i = 0
			}
			data = append(data, w.Write(binary.BigEndian)...)
		default:
			if i != 0 {
				data = append(data, vByte)
				vByte = 0
				i = 0
			}
			if len(data)%2 != 0 {
				data = append(data, 0)
			}
			data = append(data, w.Write(binary.BigEndian)...)
		}
	}
	if i != 0 {
		data = append(data, vByte)
		vByte = 0
		i = 0
	}
	if len(data)%2 != 0 {
		data = append(data, 0)
	}
	return
}

// Получает таймер строкой и возвращает объект таймера
func ParseDuration(d string) time.Duration {
	switch {
	case d == "":
		return time.Duration(0)
	case strings.HasSuffix(d, "ns"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "ns"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1)
		}
		return time.Duration(t)

	case strings.HasSuffix(d, "us"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "us"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1)
		}
		return time.Duration(t) * time.Microsecond

	case strings.HasSuffix(d, "ms"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "ms"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1)
		}
		return time.Duration(t) * time.Millisecond
	case strings.HasSuffix(d, "s"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "s"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1)
		}
		return time.Duration(t) * time.Second
	case strings.HasSuffix(d, "m"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "m"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1)
		}
		return time.Duration(t) * time.Minute
	case strings.HasSuffix(d, "h"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "h"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1)
		}
		return time.Duration(t) * time.Hour
	default:
		s := strings.TrimSpace(d)
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1)
		}
		return time.Duration(t) * time.Second
	}
}
