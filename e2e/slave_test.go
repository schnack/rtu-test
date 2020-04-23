package e2e

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSlave(t *testing.T) {
	suite.Run(t, new(SlaveTestSuit))
}

type SlaveTestSuit struct {
	suite.Suite
}

func (s *SlaveTestSuit) TestParseReadFormat() {
	v := Slave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0x01", "0x02"},
			"end":   {"0x03", "0x04"},
		},
		Len: &Len{
			Staffing:   true,
			CountBytes: 1,
			Read:       []string{"test_#"},
		},
		Crc: &Crc{
			Algorithm: "mod256",
		},
		ReadFormat: []string{
			"start", "_len_#", "crc_#", "end",
		},
	}

	start, lenPosition, suffixLen, end := v.ParseReadFormat()

	s.Equal([]byte{0x01, 0x02}, start)
	s.Equal(2, lenPosition)
	s.Equal(1, suffixLen)
	s.Equal([]byte{0x03, 0x04}, end)
}

func (s *SlaveTestSuit) TestGetSpkit() {
	v := Slave{
		ByteOrder: "big",
		MaxLen:    255,
		Len: &Len{
			Staffing:   true,
			CountBytes: 1,
			Read:       []string{"test_#"},
		},
	}
	split := v.GetSplit([]byte{1, 2}, 2, 1, []byte{3, 4})

	// Один полный пакет
	offset, data, err := split([]byte{1, 2, 0, 1, 3, 4}, true)
	s.Equal(6, offset)
	s.Equal([]byte{0x1, 0x2, 0x0, 0x1, 0x3, 0x4}, data)
	s.EqualError(err, "final token")

	// Один полный пакет с данными
	offset, data, err = split([]byte{1, 2, 2, 0xDD, 0xDD, 1, 3, 4}, false)
	s.Equal(8, offset)
	s.Equal([]byte{0x1, 0x2, 0x2, 0xDD, 0xDD, 0x1, 0x3, 0x4}, data)
	s.NoError(err)

	// получен пакет с мусорным началом
	offset, data, err = split([]byte{56, 1, 2, 0, 1, 3, 4}, false)
	s.Equal(1, offset)
	s.Nil(data)
	s.NoError(err)

	// получен пакет с мусорным концом
	offset, data, err = split([]byte{1, 2, 0, 1, 3, 5}, false)
	s.Equal(1, offset)
	s.Nil(data)
	s.NoError(err)
}
