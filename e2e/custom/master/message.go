package master

import (
	"rtu-test/e2e/common"
	"time"
)

type Message struct {
	Message string `yaml:"message"`
	Pause   string `yaml:"pause"`
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) GetPause() time.Duration {
	return common.ParseDuration(m.Pause)
}
