package e2e

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"rtu-test/e2e/common"
	"rtu-test/e2e/reports"
	"time"
)

type Message struct {
	Message string `yaml:"message"`
	Pause   string `yaml:"pause"`
}

func (m *Message) PrintReportMasterTest(report reports.ReportMasterTest) {
	d := common.ParseDuration(m.Pause)
	report.Pause = d.String()

	if m.Message != "" {
		message := common.Render(m.Message, report)
		if d < 0 {
			message = fmt.Sprintf("%s %s", message, "[Enter]")
			logrus.Info(message)

			if Init().Log != LogStdout {
				fmt.Println(message)
			}

			var t string
			_, _ = fmt.Scanln(&t)
		} else {
			logrus.Info(message)
			if Init().Log != LogStdout {
				fmt.Println(message)
			}
		}
	}

	if d > 0 {
		time.Sleep(d)
	}
}

func (m *Message) PrintReportSlaveTest(report reports.ReportSlaveTest) {
	if m.Message != "" {
		message := common.Render(m.Message, report)
		logrus.Info(message)
		if Init().Log != LogStdout {
			fmt.Println(message)
		}
	}
}

func (m *Message) PrintReportMasterGroup(report reports.ReportGroup) {
	d := common.ParseDuration(m.Pause)
	report.Pause = d.String()

	if m.Message != "" {
		message := common.Render(m.Message, report)
		if d < 0 {
			message = fmt.Sprintf("%s %s", message, "[Enter]")
			logrus.Info(message)

			if Init().Log != LogStdout {
				fmt.Println(message)
			}

			var t string
			_, _ = fmt.Scanln(&t)
		} else {
			logrus.Info(message)
			if Init().Log != LogStdout {
				fmt.Println(message)
			}
		}
	}

	if d > 0 {
		time.Sleep(d)
	}
}

func (m *Message) PrintReportMasterGroups(reports reports.ReportGroups) {
	d := common.ParseDuration(m.Pause)

	if m.Message != "" {
		if d < 0 {
			reports.Pause = m.Pause
			message := common.Render(m.Message, reports)
			logrus.Info(message)

			if Init().Log != LogStdout {
				fmt.Println(message)
			}

			var t string
			_, _ = fmt.Scanln(&t)
		} else {
			reports.Pause = d.String()
			message := common.Render(m.Message, reports)

			logrus.Info(message)
			if Init().Log != LogStdout {
				fmt.Println(message)
			}
		}
	}

	if d > 0 {
		time.Sleep(d)
	}
}
