package unit

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
)

type Device struct {
	Version      string       `yaml:"version"`
	Name         string       `yaml:"name"`
	Description  string       `yaml:"description"`
	Port         string       `yaml:"port"`
	BoundRate    string       `yaml:"boundRate"`
	DataBits     string       `yaml:"dataBits"`
	Parity       string       `yaml:"parity"`
	StopBits     string       `yaml:"stopBits"`
	Timeout      string       `yaml:"timeout"`
	Logs         string       `yaml:"logs"`
	ModbusClient ModbusClient `yaml:"modbusClient"`
}

func (d *Device) Load() error {
	file, err := os.Open("example.yml")
	if err != nil {
		return fmt.Errorf("device load config: %s", err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(d); err != nil {
		return fmt.Errorf("parse yaml error: %s", err)
	}
	return nil
}
