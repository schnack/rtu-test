package e2e

type Len struct {
	Staffing bool `yaml:"staffing"`
	// от 1 до 8
	CountBytes int      `yaml:"coundBytes"`
	Read       []string `yaml:"read"`
	Write      []string `yaml:"write"`
	Error      []string `yaml:"error"`
}
