package slave

import (
	"rtu-test/e2e/common"
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
	// TODO проверим принадлежит ли пакет этому тесту
	return false
}

// Запускает тест и поверяет значение
func (s *CustomSlaveTest) Exec(data []byte, report *ReportCustomSlaveTest) {
	//TODO Выполнение самого теста
	return
}

// Возвращает данны для записи в компорт
func (s *CustomSlaveTest) ReturnData() []byte {
	// TODO подготовка данных для ответа устройству
	return nil
}

func (s *CustomSlaveTest) ReturnError() []byte {
	// TODO подготовка ошибки для устройства
	return nil
}

// Генерирует объект отчета для этого теста. Начальные данные можно использовать в сообщении before
func (s *CustomSlaveTest) GetReport() *ReportCustomSlaveTest {
	return &ReportCustomSlaveTest{
		Name:     s.Name,
		Pass:     false,
		Skip:     s.Skip,
		Write:    nil,
		Expected: nil,
		GotByte:  nil,
	}
}
