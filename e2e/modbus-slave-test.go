package e2e

type ModbusSlaveTest struct {
	Name     string              `yaml:"name"`
	Skip     string              `yaml:"skip"`
	Before   string              `yaml:"before"`
	Next     []string            `yaml:"next"`
	Lifetime int                 `yaml:"lifetime"`
	TimeOut  string              `yaml:"timeout"`
	AutoRun  string              `yaml:"autorun"`
	Function string              `yaml:"function"`
	Address  *uint16             `yaml:"address"`
	Quantity *uint16             `yaml:"quantity"`
	Data     []*Value            `yaml:"data"`
	Expected map[string][]*Value `yaml:"expected"`
	Write    map[string][]*Value `yaml:"write"`
	Success  string              `yaml:"success"`
	Error    string              `yaml:"error"`
	After    string              `yaml:"after"`
}
