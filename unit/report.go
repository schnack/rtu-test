package unit

type Report struct {
	Pass        bool
	Name        string
	Type        TypeValue
	Expected    []byte
	ExpectedMin []byte
	ExpectedMax []byte
	Got         []byte
}
