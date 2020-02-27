package e2e

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.bug.st/serial"
	"strings"
)

type Master struct {
	Port      string `yaml:"port"`
	BoundRate int    `yaml:"boundRate"`
	DataBits  int    `yaml:"dataBits"`
	Parity    string `yaml:"parity"`
	StopBits  int    `yaml:"stopBits"`
	Timeout   string `yaml:"timeout"`
	Filter    string `yaml:"filter"`

	Const       map[string][]string `yaml:"const"`
	Staffing    *Staffing           `yaml:"staffing"`
	Len         map[string]*Len     `yaml:"len"`
	Crc         map[string]*Crc     `yaml:"crc"`
	WriteFormat []string            `yaml:"writeFormat"`
	ReadFormat  []string            `yaml:"readFormat"`
	ErrorFormat []string            `yaml:"errorFormat"`

	Tests map[string][]*MasterTest `yaml:"tests"`
}

func (m *Master) getPort() (serial.Port, error) {
	parity := serial.NoParity
	switch m.Parity {
	case "N":
		parity = serial.NoParity
	case "E":
		parity = serial.EvenParity
	case "O":
		parity = serial.OddParity

	}

	stopBits := serial.TwoStopBits

	switch m.StopBits {
	case 1:
		stopBits = serial.OneStopBit
	case 15:
		stopBits = serial.OnePointFiveStopBits
	case 2:
		stopBits = serial.TwoStopBits
	}

	return serial.Open(m.Port, &serial.Mode{
		BaudRate: m.BoundRate,
		DataBits: m.DataBits,
		Parity:   parity,
		StopBits: stopBits,
	})
}

func (m *Master) Run(reports *ReportGroups) error {
	// TODO Доработать этот функционал
	port, err := m.getPort()
	if err != nil {
		return fmt.Errorf("port error %s", err)
	}
	defer port.Close()

	filterGroup := ""
	filterTest := ""
	filter := strings.Split(m.Filter, ":")
	if len(filter) > 1 {
		filterGroup = filter[0]
		filterTest = filter[1]
	} else {
		filterGroup = filter[0]
	}

	for group, tests := range m.Tests {
		if filterGroup != "" && filterGroup != "all" && filterGroup != group {
			continue
		}
		report := ReportGroup{Name: group}
		logrus.Warnf(render(TestGROUP, report))
		for _, test := range tests {
			if filterTest != "" && filterTest != "all" && filterTest != test.Name {
				continue
			}
			report.Tests = append(report.Tests, test.Run(port))
			// При необходимости закрываем порт
			if test.Disconnect {
				port.Close()
			}
		}
		reports.ReportGroup = append(reports.ReportGroup, report)
	}

	return nil
}
