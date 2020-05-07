package reports

// Отчет тестирования Modbus Slave
type ReportSlaveTest struct {
	Name                     string
	Pass                     bool
	Skip                     string
	ExpectedCoils            []ReportExpected
	ExpectedDiscreteInput    []ReportExpected
	ExpectedHoldingRegisters []ReportExpected
	ExpectedInputRegisters   []ReportExpected
}
