package unit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseStringByte(sb string) ([]byte, error) {
	buf := new(bytes.Buffer)
	byteClear := strings.ReplaceAll(strings.ReplaceAll(sb, " ", ""), "0x", "")
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

func dataSingleCoil(data []byte) []byte {
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

func countBit(v []*Value, is16bit bool) (bits uint16, err error) {
	for _, w := range v {
		if w.Type() == Bool {
			bits++
		} else {
			if bits%8 != 0 {
				bits += 8 - bits%8
			}
			byteData, err := w.Write()
			if err != nil {
				return 0, fmt.Errorf("%s", err)
			}
			bits += 8 * uint16(len(byteData))
		}
	}
	if is16bit {
		if bits%16 != 0 {
			bits += 16 - bits%16
		}
		return bits / 16, nil
	} else {
		return bits, nil
	}
}

func valueToByte(v []*Value) (data []byte, err error) {
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
			b, err := w.Write()
			if err != nil {
				return data, err
			}
			data = append(data, b...)
		}
	}
	if i != 0 {
		data = append(data, vByte)
		vByte = 0
		i = 0
	}
	return
}

// byteToEq - byte-by-byte comparison
func byteToEq(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func parseDuration(d string) time.Duration {
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
