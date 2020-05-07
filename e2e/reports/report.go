package reports

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
