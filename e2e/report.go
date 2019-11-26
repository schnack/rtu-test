package e2e

import "time"

type ReportWrite struct {
	Name    string
	Type    string
	Data    string
	DataHex string
	DataBin string
}

type ReportExpected struct {
	Name        string
	Pass        bool
	Type        string
	Expected    string
	ExpectedHex string
	ExpectedBin string
	Got         string
	GotHex      string
	GotBin      string
}

type ReportGroup struct {
	Name  string
	Pause string
	Tests []ReportTest
}

type ReportGroups struct {
	Name        string
	Description string
	Pause       string
	ReportGroup []ReportGroup
}

type ReportTest struct {
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
