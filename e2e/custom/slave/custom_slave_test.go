package slave

import (
	"github.com/stretchr/testify/suite"
	"rtu-test/e2e/custom/module"
	"testing"
)

func TestCustomSlave(t *testing.T) {
	suite.Run(t, new(CustomSlaveTestSuit))
}

type CustomSlaveTestSuit struct {
	suite.Suite
}

func (s *CustomSlaveTestSuit) TestParseReadFormat() {
	v := CustomSlave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0x01", "0x02"},
			"end":   {"0x03", "0x04"},
		},
		Len: &module.LenBytes{
			Staffing:   true,
			CountBytes: 1,
			Read:       []string{"data#"},
		},
		Crc: &module.Crc{
			Algorithm: "mod256",
		},
		ReadFormat: []string{
			"start", "len#", "crc#", "end",
		},
	}

	start, lenPosition, suffix, end := v.ParseReadFormat()

	s.Equal([]byte{0x01, 0x02}, start)
	s.Equal(2, lenPosition)
	s.Equal([]string{"crc#"}, suffix)
	s.Equal([]byte{0x03, 0x04}, end)
}

func (s *CustomSlaveTestSuit) TestGetSplitStartEnd() {
	v := CustomSlave{
		ByteOrder: "big",
		MaxLen:    255,
	}

	split := v.GetSplitStartEnd([]byte{1, 2}, []byte{3, 4})
	// Один полный пакет
	offset, data, err := split([]byte{1, 2, 0, 1, 3, 4}, true)
	s.Equal(6, offset)
	s.Equal([]byte{0x1, 0x2, 0x0, 0x1, 0x3, 0x4}, data)
	s.EqualError(err, "final token")

	offset, data, err = split([]byte{0, 1, 0, 2, 1, 2, 0, 1, 3, 4, 4, 3, 4}, true)
	s.Equal(10, offset)
	s.Equal([]byte{0x1, 0x2, 0x0, 0x1, 0x3, 0x4}, data)
	s.EqualError(err, "final token")

}

func (s *CustomSlaveTestSuit) TestGetSplitLen() {
	v := CustomSlave{
		ByteOrder: "big",
		MaxLen:    255,
		Len: &module.LenBytes{
			CountBytes: 2,
		},
		Const: map[string][]string{
			"end": {
				"0x01", "0x02",
			},
		},
	}

	split := v.GetSplitLen([]byte{1, 2}, 2, []string{"end"})
	// Один полный пакет
	offset, data, err := split([]byte{1, 2, 0, 1, 6, 3, 4}, true)
	s.Equal(7, offset)
	s.Equal([]byte{0x1, 0x2, 0x0, 0x1, 0x6, 0x3, 0x4}, data)
	s.EqualError(err, "final token")

	offset, data, err = split([]byte{0, 1, 0, 2, 1, 2, 0, 1, 6, 3, 4, 4, 3, 4}, true)
	s.Equal(11, offset)
	s.Equal([]byte{0x1, 0x2, 0x0, 0x1, 0x6, 0x3, 0x4}, data)
	s.EqualError(err, "final token")

	v = CustomSlave{
		ByteOrder: "big",
		MaxLen:    255,
		Crc: &module.Crc{
			Algorithm: module.Mod256,
			Staffing:  false,
		},
	}

	split = v.GetSplitLen([]byte{1, 2}, 2, []string{"crc#"})
	offset, data, err = split([]byte{1, 2, 1, 6, 3, 4}, true)
	s.Equal(5, offset)
	s.Equal([]byte{0x1, 0x2, 0x1, 0x6, 0x3}, data)
	s.EqualError(err, "final token")
}

func (s *CustomSlaveTestSuit) TestStaffingProcessing() {
	v := CustomSlave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0xcf", "0xbf"},
			"end":   {"0xff", "0xef"},
		},
		Staffing: &module.Staffing{
			Byte:    "0x00",
			Pattern: []string{"start", "end"},
		},
	}

	s.Equal([]byte{0x1, 0xcf, 0x0, 0x2, 0xbf, 0x0, 0x3, 0xff, 0x0, 0x4, 0xef, 0x0, 0x5}, v.StaffingProcessing(true, []byte{0x01, 0xcf, 0x02, 0xbf, 0x03, 0xff, 0x04, 0xef, 0x05}))
	s.Equal([]byte{0x01, 0xcf, 0x02, 0xbf, 0x03, 0xff, 0x04, 0xef, 0x05}, v.StaffingProcessing(false, []byte{0x1, 0xcf, 0x0, 0x2, 0xbf, 0x0, 0x3, 0xff, 0x0, 0x4, 0xef, 0x0, 0x5}))
}

func (s *CustomSlaveTestSuit) TestCalcLen() {
	v := CustomSlave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0xcf", "0xbf"},
			"end":   {"0xff", "0xef"},
		},
		Staffing: &module.Staffing{
			Byte:    "0x00",
			Pattern: []string{"start", "end"},
		},
		Len: &module.LenBytes{
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

func (s *CustomSlaveTestSuit) TestCalcCrc() {
	v := CustomSlave{
		ByteOrder: "big",
		Const: map[string][]string{
			"start": {"0xcf", "0xbf"},
			"end":   {"0xff", "0xef"},
		},
		Staffing: &module.Staffing{
			Byte:    "0x00",
			Pattern: []string{"start", "end"},
		},
		Len: &module.LenBytes{
			Staffing:   true,
			CountBytes: 2,
			Read:       []string{"data#", "end"},
			Write:      []string{"data#"},
			Error:      []string{"end"},
		},
		Crc: &module.Crc{
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
