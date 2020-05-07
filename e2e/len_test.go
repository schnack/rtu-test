package e2e

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestLenBytes(t *testing.T) {
	suite.Run(t, new(LenBytesTestSuite))
}

type LenBytesTestSuite struct {
	suite.Suite
}

func (s *LenBytesTestSuite) TestContains() {
	l := LenBytes{
		Staffing:   false,
		CountBytes: 0,
		Read:       []string{"data#", "crc#"},
		Write:      []string{"data#"},
		Error:      []string{"crc#"},
	}

	s.True(l.Contains(ActionRead, "data#"))
	s.True(l.Contains(ActionRead, "crc#"))
	s.True(l.Contains(ActionWrite, "data#"))
	s.False(l.Contains(ActionWrite, "crc#"))
	s.False(l.Contains(ActionError, "data#"))
	s.True(l.Contains(ActionError, "crc#"))
}
