package module

import (
	"encoding/binary"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCrc(t *testing.T) {
	suite.Run(t, new(CrcTestSuit))
}

type CrcTestSuit struct {
	suite.Suite
}

func (s *CrcTestSuit) TestParseReadFormat() {
	crc := Crc{
		Algorithm: "modBus",
		ByteOrder: "",
	}

	s.Equal([]byte{0xc0, 0x40}, crc.Calc(binary.LittleEndian, []byte{0xfe, 0xfe, 0x00, 0x06, 0x04, 0x09, 0x00, 0x00, 0x21, 0x00, 0x00}))
}
