package e2e

import (
	"go.bug.st/serial"
	"io"
	"sync"
	"time"
)

type SerialPortConfig struct {
	Port     string
	BaudRate int
	DataBits int
	Parity   string
	StopBits int
	// Интервал между adu
	SilentInterval time.Duration
	Timeout        time.Duration
}

type SerialPort struct {
	Config *SerialPortConfig

	IdleTimeout  time.Duration
	mu           sync.Mutex
	port         io.ReadWriteCloser
	lastActivity time.Time
	closeTimer   *time.Timer
}

func (s *SerialPort) Connect() (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.port == nil {
		s.port, err = openSerialPort(s.Config.Port, s.Config.Parity, s.Config.BaudRate, s.Config.DataBits, s.Config.StopBits)
		if err != nil {
			s.port = nil
		}
	}
	return
}

func (s *SerialPort) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.port != nil {
		_ = s.port.Close()
	}
	s.port = nil
}

var openSerialPort = func(port string, parity string, baundRate, dataBits, stopBits int) (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: baundRate,
		DataBits: dataBits,
	}
	switch parity {
	case "E":
		mode.Parity = serial.EvenParity
	case "O":
		mode.Parity = serial.OddParity
	default:
		mode.Parity = serial.NoParity
	}

	switch stopBits {
	case 1:
		mode.StopBits = serial.OneStopBit
	case 15:
		mode.StopBits = serial.OnePointFiveStopBits
	default:
		mode.StopBits = serial.TwoStopBits
	}

	return serial.Open(port, mode)
}
