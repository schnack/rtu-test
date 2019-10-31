package unit

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Message struct {
	Message string `yaml:"message"`
	Pause   string `yaml:"pause"`
}

func (m *Message) Print() {
	logrus.Info(m.Message)
	d := parseDuration(m.Pause)
	if d < 0 {
		logrus.Info("Press ENTER to continue...")
		var t string
		_, _ = fmt.Scanln(&t)
	} else {
		logrus.Infof("Pause %s ...", d.String())
		time.Sleep(d)
	}
}

func (m *Message) parsePause() time.Duration {
	return parseDuration(m.Pause)
}
