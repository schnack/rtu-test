package master

import (
	"rtu-test/e2e/common"
	"time"
)

// Отчет о тестировании команд ModBus мастера
type ReportMasterTest struct {
	Name     string
	Pass     bool
	Pause    string
	Skip     string
	Write    []common.ReportWrite
	Expected []common.ReportExpected
	GotByte  []byte
	GotTime  time.Duration
	GotError string
}

type ReportGroup struct {
	Name  string
	Pause string
	Tests []ReportMasterTest
}

type ReportGroups struct {
	Name        string
	Description string
	Pause       string
	ReportGroup []ReportGroup
}
