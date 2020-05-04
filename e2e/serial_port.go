package e2e

import (
	"go.bug.st/serial"
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

type SerialPort interface {
	Connect() (err error)
	Close() error
	Read(p []byte) (int, error)
	Write(p []byte) (n int, err error)
}

var NewSerialPort = func(config *SerialPortConfig) SerialPort {
	return &serialPort{
		Config: config,
		mu:     sync.Mutex{},
		port:   nil,
	}
}

type serialPort struct {
	Config *SerialPortConfig

	mu           sync.Mutex
	port         serial.Port
	lastActivity time.Time
	closeTimer   time.Timer
}

func (s *serialPort) Read(p []byte) (int, error) {
	if s.port == nil {
		if err := s.Connect(); err != nil {
			return 0, err
		}
	}
	return s.port.Read(p)
}

func (s *serialPort) Write(p []byte) (n int, err error) {
	if s.port == nil {
		if err := s.Connect(); err != nil {
			return 0, err
		}
	}
	return s.port.Write(p)
}

// Открываем порт
func (s *serialPort) Connect() (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.port == nil {
		mode := &serial.Mode{
			BaudRate: s.Config.BaudRate,
			DataBits: s.Config.DataBits,
		}
		switch s.Config.Parity {
		case "E":
			mode.Parity = serial.EvenParity
		case "O":
			mode.Parity = serial.OddParity
		default:
			mode.Parity = serial.NoParity
		}

		switch s.Config.StopBits {
		case 1:
			mode.StopBits = serial.OneStopBit
		case 15:
			mode.StopBits = serial.OnePointFiveStopBits
		default:
			mode.StopBits = serial.TwoStopBits
		}

		s.port, err = serial.Open(s.Config.Port, mode)
		if err != nil {
			s.port = nil
			return err
		}
	}
	return nil
}

// Закрываем порт
func (s *serialPort) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.port != nil {
		err := s.port.Close()
		s.port = nil
		if err != nil {
			return err
		}
	}
	return nil
}
