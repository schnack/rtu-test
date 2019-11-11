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

func (m *Message) Print(report ReportTest) {
	d := parseDuration(m.Pause)
	report.Pause = d.String()

	if m.Message != "" {
		message := render(m.Message, report)
		if d < 0 {
			message = fmt.Sprintf("%s %s", message, "[Enter]")
			logrus.Info(message)

			if Init().Log != LogStdout {
				fmt.Println(message)
			}

			var t string
			_, _ = fmt.Scanln(&t)
		}
	}

	if d > 0 {
		time.Sleep(d)
	}
}
