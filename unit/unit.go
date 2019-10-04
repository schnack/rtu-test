package unit

import (
	"github.com/goburrow/modbus"
	"log"
)

type MudbusTest struct {
	Name          string        `yaml:"name"`
	Before        Message       `yaml:"before"`
	Function      Function      `yaml:"function"`
	Address       Address       `yaml:"address"`
	Quantity      Quantity      `yaml:"quantity"`
	Write         []Value       `yaml:"write"`
	Expected      []Value       `yaml:"expected"`
	ExpectedError ExpectedError `yaml:"expectedError"`
	Success       Message       `yaml:"success"`
	Error         Message       `yaml:"error"`
	After         Message       `yaml:"after"`
}

func (mt *MudbusTest) Exec(client modbus.Client) error {
	log.Printf("Run %s", mt.Name)
	mt.Before.Print()

	mt.Error.Print()

	mt.Success.Print()

	mt.After.Print()
	return nil
}
