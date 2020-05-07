package e2e

type SlaveTest struct {
	Name   string   `yaml:"name"`
	Skip   string   `yaml:"skip"`
	Before Message  `yaml:"before"`
	Next   []string `yaml:"next"`
	Fatal  string   `yaml:"fatal"`
	// Количество срабатываний теста
	LifeTime   int     `yaml:"lifetime"`
	Timeout    string  `yaml:"timeout"`
	Pattern    []Value `yaml:"pattern"`
	Write      []Value `yaml:"write"`
	WriteError []Value `yaml:"writeError"`
	Expected   []Value `yaml:"expected"`
	Success    Message `yaml:"success"`
	Error      Message `yaml:"error"`
	After      Message `yaml:"after"`
}

// Проверяем пакет принадлежит этому тесту или нет с использованием Pattern
func (s *SlaveTest) Check(data []byte) bool {

	return false
}

// Запускает тест и поверяет значение
func (s *SlaveTest) Exec(data []byte, report *ReportSlaveTest) {

	return
}

// Возвращает данны для записи в компорт
func (s *SlaveTest) WriteData() []byte {
	return nil
}
