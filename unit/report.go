package unit

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

type Report struct {
	Group         string
	Test          string
	Pause         string
	Pass          bool
	Write         []ReportWrite
	Expected      []ReportExpected
	ExpectedTime  string
	ExpectedError string
	GotTime       string
	GotError      string
}
