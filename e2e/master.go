package e2e

type Master struct {
	Port      string `yaml:"port"`
	BoundRate int    `yaml:"boundRate"`
	DataBits  int    `yaml:"dataBits"`
	Parity    string `yaml:"parity"`
	StopBits  int    `yaml:"stopBits"`
	Timeout   string `yaml:"timeout"`
	Filter    string `yaml:"filter"`

	Const       map[string][]string `yaml:"const"`
	Staffing    *Staffing           `yaml:"staffing"`
	Len         map[string]*Len     `yaml:"len"`
	Crc         map[string]*Crc     `yaml:"crc"`
	WriteFormat []string            `yaml:"writeFormat"`
	ReadFormat  []string            `yaml:"readFormat"`
	ErrorFormat []string            `yaml:"errorFormat"`

	Tests map[string][]*MasterTest `yaml:"tests"`
}
