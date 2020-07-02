package module

type LenBytes struct {
	Staffing      bool `yaml:"staffing"`
	CountStaffing bool `yaml:"countStaffing"`
	// от 1 до 8
	CountBytes int      `yaml:"coundBytes"`
	Read       []string `yaml:"read"`
	Write      []string `yaml:"write"`
	Error      []string `yaml:"error"`
}

// Проверяет есть ли текущий параметр в массиве
func (l *LenBytes) Contains(action, param string) bool {
	var data []string
	switch action {
	case ActionRead:
		data = l.Read
	case ActionWrite:
		data = l.Write
	case ActionError:
		data = l.Error
	}

	for _, read := range data {
		if read == param {
			return true
		}
	}
	return false
}
