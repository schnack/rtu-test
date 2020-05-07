package master

import "time"

func NewFixtureModBusClient(Results []byte, err error) *FixtureModBusClient {
	return &FixtureModBusClient{Results: Results, Error: err, Sleep: 1}
}

type FixtureModBusClient struct {
	Address     uint16
	Quantity    uint16
	SingleValue uint16
	Value       []byte
	Sleep       time.Duration

	Results []byte
	Error   error
}

func (f *FixtureModBusClient) ReadCoils(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = quantity
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadDiscreteInputs(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = quantity
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteSingleCoil(address, value uint16) (results []byte, err error) {
	f.Address = address
	f.SingleValue = value
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteMultipleCoils(address, quantity uint16, value []byte) (results []byte, err error) {
	f.Address = address
	f.Quantity = quantity
	f.Value = value
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadInputRegisters(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = quantity
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	f.Address = address
	f.Quantity = quantity
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	f.Address = address
	f.SingleValue = value
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	f.Address = address
	f.Quantity = quantity
	f.Value = value
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadWriteMultipleRegisters(readAddress, readQuantity, writeAddress, writeQuantity uint16, value []byte) (results []byte, err error) {
	f.Address = readAddress
	f.Quantity = readQuantity
	f.Value = value
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) MaskWriteRegister(address, andMask, orMask uint16) (results []byte, err error) {
	f.Address = address
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}

func (f *FixtureModBusClient) ReadFIFOQueue(address uint16) (results []byte, err error) {
	f.Address = address
	time.Sleep(f.Sleep)
	return f.Results, f.Error
}
