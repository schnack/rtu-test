package e2e

import (
	"bytes"
	"io"
)

type FixtureSerialPort struct {
	serialPort

	// Канал буфферизированный для записи
	BytesRead chan byte
	// Буфер записанной инфы
	BufferWrite bytes.Buffer

	srcFunc func(config *SerialPortConfig) SerialPort
}

func (f *FixtureSerialPort) Load() {
	f.srcFunc = NewSerialPort
	f.BytesRead = make(chan byte, 2550)

	NewSerialPort = func(config *SerialPortConfig) SerialPort {
		f.Config = config
		return f
	}
}

func (f *FixtureSerialPort) Unload() {
	NewSerialPort = f.srcFunc
}

func (f *FixtureSerialPort) Connect() error {
	return nil
}

func (f *FixtureSerialPort) Close() error {
	return nil
}

func (f *FixtureSerialPort) Read(p []byte) (int, error) {
	var ok bool
	for i := range p {
		if p[i], ok = <-f.BytesRead; ok {
			return len(p), io.EOF
		}
	}
	return len(p), nil
}

func (f *FixtureSerialPort) Write(p []byte) (int, error) {
	return f.BufferWrite.Write(p)
}
