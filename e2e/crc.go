package e2e

import "github.com/sirupsen/logrus"

type Crc struct {
	Algorithm string   `yaml:"algorithm"`
	Staffing  bool     `yaml:"staffing"`
	Read      []string `yaml:"read"`
	Write     []string `yaml:"write"`
	Error     []string `yaml:"error"`
}

// CheckSum8 Modulo 256.
func CrcMod256(data []byte) uint8 {
	var sum uint8
	for _, b := range data {
		sum += b
	}
	return sum & 0xff
}

// Len - возвращает длину данных crc
func (c *Crc) Len() int {
	switch c.Algorithm {
	case "mod256":
		return 1
	default:
		logrus.Fatalf("crc algorithm no support %s", c.Algorithm)
	}
	return 0
}
