package unit

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"text/template"
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
		t := template.Must(template.New("message").Parse(m.Message))
		buff := new(bytes.Buffer)
		if err := t.Execute(buff, report); err != nil {
			logrus.Fatal(err)
		}
		if d < 0 {
			buff.WriteString(" [Enter]")
			var t string
			_, _ = fmt.Scanln(&t)
		}
		Init().Display(buff.String())
	}

	if d > 0 {
		time.Sleep(d)
	}
}
