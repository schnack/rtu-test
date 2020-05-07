package reports

import "time"

// Отчет о тестировании команд ModBus мастера
type ReportMasterTest struct {
	Name     string
	Pass     bool
	Pause    string
	Skip     string
	Write    []ReportWrite
	Expected []ReportExpected
	GotByte  []byte
	GotTime  time.Duration
	GotError string
}
