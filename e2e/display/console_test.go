package display

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestConsole(t *testing.T) {
	suite.Run(t, new(ConsoleTestSuite))
}

type ConsoleTestSuite struct {
	suite.Suite
}

func (s *ConsoleTestSuite) TestPrint() {
	buff := new(bytes.Buffer)
	console := &ConsoleRender{output: buff}
	currentTime := time.Now()
	console.Print(&testMessage{}, &testMessage{Test: "testValue"})
	s.True(time.Since(currentTime) > time.Millisecond)
	s.Equal("test testValue\n", buff.String())
}

// Helper struct
type testMessage struct {
	Test string
}

func (t *testMessage) GetMessage() string      { return "test {{.Test}}" }
func (t *testMessage) GetPause() time.Duration { return 2 * time.Millisecond }
