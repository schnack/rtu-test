package reports

// Результат теста одного параметра
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
