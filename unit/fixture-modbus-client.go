package unit

import "github.com/goburrow/modbus"

func NewFixtureModBusClient(Results []byte, err error) modbus.Client {
	return &FixtureModBusClient{Results: Results, Error: err}
}

type FixtureModBusClient struct {
	Address     uint16
	Quantity    uint16
	SingleValue uint16
	Value       []byte

	Results []byte
	Error   error
}

func (f *FixtureModBusClient) ReadCoils(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = address
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadDiscreteInputs(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = address
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteSingleCoil(address, value uint16) (results []byte, err error) {
	f.Address = address
	f.SingleValue = value
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteMultipleCoils(address, quantity uint16, value []byte) (results []byte, err error) {
	f.Address = address
	f.Quantity = address
	f.Value = value
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadInputRegisters(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = address
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = address
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = address
	f.SingleValue = value
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	f.Address = address
	f.Quantity = address
	f.Value = value
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadWriteMultipleRegisters(readAddress, readQuantity, writeAddress, writeQuantity uint16, value []byte) (results []byte, err error) {
	f.Address = readAddress
	f.Quantity = readQuantity
	f.Value = value
	return f.Results, f.Error
}

func (f *FixtureModBusClient) MaskWriteRegister(address, andMask, orMask uint16) (results []byte, err error) {
	f.Address = address
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadFIFOQueue(address uint16) (results []byte, err error) {
	f.Address = address
	return f.Results, f.Error
}
