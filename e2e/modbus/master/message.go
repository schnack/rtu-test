package master

import (
	"rtu-test/e2e/common"
	"rtu-test/e2e/display"
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

func (m *Message) PrintReportMasterTest(report ReportMasterTest) {
	d := common.ParseDuration(m.Pause)
	report.Pause = d.String()
	display.Console().Print(m, report)
}
