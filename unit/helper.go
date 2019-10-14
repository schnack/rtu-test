package unit

import (
	"strconv"
	"strings"
	"time"
)

func countBit(v []Value, is16bit bool) (bits uint16) {
	for _, w := range v {
		if w.Type() == Bool {
			bits++
		} else {
			if bits%8 != 0 {
				bits += 8 - bits%8
			}
			byteData, err := w.Write()
			if err != nil {
				return 0
			}
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

func parseDuration(d string) (time.Duration, string) {
	switch {
	case strings.HasSuffix(d, "ns"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "ns"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t), "ns"
	case strings.HasSuffix(d, "us"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "us"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Microsecond, "us"
	case strings.HasSuffix(d, "ms"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "ms"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Millisecond, "ms"
	case strings.HasSuffix(d, "s"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "s"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Second, "s"
	case strings.HasSuffix(d, "m"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "m"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Minute, "m"
	case strings.HasSuffix(d, "h"):
		s := strings.TrimSpace(strings.TrimSuffix(d, "h"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Hour, "h"
	default:
		s := strings.TrimSpace(d)
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Second, "s"
	}
}
