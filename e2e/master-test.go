package e2e

type MasterTest struct {
	Name       string   `yaml:"name"`
	Skip       string   `yaml:"skip"`
	Before     Message  `yaml:"before"`
	Write      []*Value `yaml:"write"`
	Expected   []*Value `yaml:"expected"`
	Success    Message  `yaml:"success"`
	Error      Message  `yaml:"error"`
	After      Message  `yaml:"after"`
	Fatal      string   `yaml:"fatal"`
	Disconnect bool     `yaml:"disconnect"`

	// Заменяет глобальные настройки
	Const       map[string][]string `yaml:"const"`
	Staffing    *Staffing           `yaml:"staffing"`
	Len         map[string]*Len     `yaml:"len"`
	Crc         map[string]*Crc     `yaml:"crc"`
	WriteFormat []string            `yaml:"writeFormat"`
	ReadFormat  []string            `yaml:"readFormat"`
	ErrorFormat []string            `yaml:"errorFormat"`
}
