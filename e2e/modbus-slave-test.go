package e2e

type ModbusSlaveTest struct {
	Name     string              `yaml:"name"`
	Skip     string              `yaml:"skip"`
	Before   Message             `yaml:"before"`
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
	Success  Message             `yaml:"success"`
	Error    Message             `yaml:"error"`
	After    Message             `yaml:"after"`
}
