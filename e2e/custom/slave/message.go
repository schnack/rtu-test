package slave

import (
	"time"
)

type Message string

func (m *Message) GetMessage() string {
	return string(*m)
}

func (m *Message) GetPause() time.Duration {
	return 0
}
