package unit

type Report struct {
	Pass     bool
	Name     string
	Type     TypeValue
	Expected []byte
	Got      []byte
}
