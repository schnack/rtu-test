package e2e

import (
	"rtu-test/e2e/common"
	"rtu-test/e2e/display"
	"rtu-test/e2e/modbus/master"
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

func (m *Message) PrintReportMasterGroups(reports master.ReportGroups) {
	display.Console().Print(m, reports)
}
