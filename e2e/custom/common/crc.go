package common

import "github.com/sirupsen/logrus"

const (
	Mod256 = "mod256"
)

type Crc struct {
	Algorithm string   `yaml:"algorithm"`
	Staffing  bool     `yaml:"staffing"`
	Read      []string `yaml:"read"`
	Write     []string `yaml:"write"`
	Error     []string `yaml:"error"`
}

// Подсчет контрольной суммы согласно алгоритму
func (c *Crc) Calc(data []byte) []byte {
	switch c.Algorithm {
	case Mod256:
		return []byte{c.CrcMod256(data)}
	}
	return nil
}

// CheckSum8 Modulo 256.
func (c *Crc) CrcMod256(data []byte) uint8 {
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
