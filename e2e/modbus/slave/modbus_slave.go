package slave

import "rtu-test/e2e/common"

// Отчет тестирования Modbus Slave
type ReportSlaveTest struct {
	Name                     string
	Pass                     bool
	Skip                     string
	ExpectedCoils            []common.ReportExpected
	ExpectedDiscreteInput    []common.ReportExpected
	ExpectedHoldingRegisters []common.ReportExpected
	ExpectedInputRegisters   []common.ReportExpected
}
