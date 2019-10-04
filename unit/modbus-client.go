package unit

type ModbusClient struct {
	SlaveId uint8 `yaml:"slaveId"`
	Tests   Tests `yaml:"tests"`
}
