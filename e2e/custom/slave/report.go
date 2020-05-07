package slave

import (
	"rtu-test/e2e/common"
)

type ReportCustomSlaveTest struct {
	// Название теста
	Name string
	// Флаг о прошествии
	Pass bool
	// Сообщение почему тест пропущен
	Skip string
	// Отчет о отвеченных данных
	Write []common.ReportWrite
	// Отчет о пропущенных данных
	Expected []common.ReportExpected
	// сырые полученные данные
	GotByte []byte
}
