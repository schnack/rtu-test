package slave

import (
	"rtu-test/e2e/common"
	"rtu-test/e2e/modbus/slave"
)

type CustomSlaveTest struct {
	Name   string   `yaml:"name"`
	Skip   string   `yaml:"skip"`
	Before Message  `yaml:"before"`
	Next   []string `yaml:"next"`
	Fatal  string   `yaml:"fatal"`
	// Количество срабатываний теста
	LifeTime   int            `yaml:"lifetime"`
	Timeout    string         `yaml:"timeout"`
	Pattern    []common.Value `yaml:"pattern"`
	Write      []common.Value `yaml:"write"`
	WriteError []common.Value `yaml:"writeError"`
	Expected   []common.Value `yaml:"expected"`
	Success    Message        `yaml:"success"`
	Error      Message        `yaml:"error"`
	After      Message        `yaml:"after"`
}

// Проверяем пакет принадлежит этому тесту или нет с использованием Pattern
func (s *CustomSlaveTest) Check(data []byte) bool {

	return false
}

// Запускает тест и поверяет значение
func (s *CustomSlaveTest) Exec(data []byte, report *slave.ReportSlaveTest) {

	return
}

// Возвращает данны для записи в компорт
func (s *CustomSlaveTest) WriteData() []byte {
	return nil
}
