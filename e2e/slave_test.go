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
			Read:       []string{"data#"},
		},
		Crc: &Crc{
			Algorithm: "mod256",
		},
		ReadFormat: []string{
			"start", "len#", "crc#", "end",
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
			Read:       []string{"data#"},
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

func (s *SlaveTestSuit) TestAddStaffing() {
	v := Slave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0xcf", "0xbf"},
			"end":   {"0xff", "0xef"},
		},
		Staffing: &Staffing{
			Byte:    "0x00",
			Pattern: []string{"start", "end"},
		},
	}

	s.Equal([]byte{0x1, 0xcf, 0x0, 0x2, 0xbf, 0x0, 0x3, 0xff, 0x0, 0x4, 0xef, 0x0, 0x5}, v.AddStaffing([]byte{0x01, 0xcf, 0x02, 0xbf, 0x03, 0xff, 0x04, 0xef, 0x05}))
}

func (s *SlaveTestSuit) TestCalcLen() {
	v := Slave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0xcf", "0xbf"},
			"end":   {"0xff", "0xef"},
		},
		Staffing: &Staffing{
			Byte:    "0x00",
			Pattern: []string{"start", "end"},
		},
		Len: &Len{
			Staffing:   true,
			CountBytes: 2,
			Read:       []string{"data#", "end"},
			Write:      []string{"data#"},
			Error:      []string{"end"},
		},
	}

	count, b := v.CalcLen(ActionRead, []byte{1, 2, 3, 4})
	s.Equal(6, count)
	s.Equal([]byte{0x0, 0x6}, b)

	// с добавлением стаффинга
	count, b = v.CalcLen(ActionRead, []byte{1, 2, 3, 0xef, 4})
	s.Equal(8, count)
	s.Equal([]byte{0x0, 0x8}, b)

	count, b = v.CalcLen(ActionWrite, []byte{1, 2, 3, 4})
	s.Equal(4, count)
	s.Equal([]byte{0x0, 0x4}, b)

	count, b = v.CalcLen(ActionError, []byte{1, 2, 3, 4})
	s.Equal(2, count)
	s.Equal([]byte{0x0, 0x2}, b)
}

func (s *SlaveTestSuit) TestCalcCrc() {
	v := Slave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0xcf", "0xbf"},
			"end":   {"0xff", "0xef"},
		},
		Staffing: &Staffing{
			Byte:    "0x00",
			Pattern: []string{"start", "end"},
		},
		Len: &Len{
			Staffing:   true,
			CountBytes: 2,
			Read:       []string{"data#", "end"},
			Write:      []string{"data#"},
			Error:      []string{"end"},
		},
		Crc: &Crc{
			Algorithm: "mod256",
			Staffing:  false,
			Read:      []string{"data#", "end"},
			Write:     []string{"data#"},
			Error:     []string{"end"},
		},
	}

	b := v.CalcCrc(ActionRead, []byte{1, 2, 3, 4})
	s.Equal([]byte{0xf8}, b)

	// с добавлением стаффинга
	b = v.CalcCrc(ActionRead, []byte{1, 2, 3, 0xef, 4})
	s.Equal([]byte{0xe7}, b)

	b = v.CalcCrc(ActionWrite, []byte{1, 2, 3, 4})
	s.Equal([]byte{0x0a}, b)

	b = v.CalcCrc(ActionError, []byte{1, 2, 3, 4})
	s.Equal([]byte{0xee}, b)
}
