package slave

import (
	"rtu-test/e2e/display"
	"time"
)

type Message string

func (m *Message) GetMessage() string {
	return string(*m)
}

func (m *Message) GetPause() time.Duration {
	return 0
}

func (m *Message) PrintReportSlaveTest(report ReportSlaveTest) {
	display.Console().Print(m, report)
}
