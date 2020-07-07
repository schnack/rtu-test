package slave

import (
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"rtu-test/e2e/common"
	"strconv"
)

type CustomSlaveTest struct {
	Name   string   `yaml:"name"`
	Skip   string   `yaml:"skip"`
	Before Message  `yaml:"before"`
	Next   []string `yaml:"next"`
	Fatal  string   `yaml:"fatal"`
	// Количество срабатываний теста
	LifeTime   int            `yaml:"lifetime"`
	Timeout    string         `yaml:"timeout"`
	Pattern    []common.Value `yaml:"pattern"`
	Write      []common.Value `yaml:"write"`
	WriteError []common.Value `yaml:"writeError"`
	Expected   []common.Value `yaml:"expected"`
	Success    Message        `yaml:"success"`
	Error      Message        `yaml:"error"`
	After      Message        `yaml:"after"`
}

// Проверяем пакет принадлежит этому тесту или нет с использованием Pattern
func (s *CustomSlaveTest) Check(data []byte) bool {
	// TODO проверим принадлежит ли пакет этому тесту
	logrus.Infof("Check: % 02x", data)
	offsetBit := 0
	for i := range s.Pattern {
		if s.Pattern[i].Address != "" {
			rawAddress, err := strconv.Atoi(s.Pattern[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			offsetBit = (rawAddress - 1) * 8
		}

		var report common.ReportExpected
		offsetBit, report = s.Pattern[i].Check(data, 0, "", offsetBit, 8, binary.LittleEndian)
		logrus.Warn(offsetBit)
		logrus.Warn(report)
	}
	return false
}

// Запускает тест и поверяет значение
func (s *CustomSlaveTest) Exec(data []byte, report *ReportCustomSlaveTest) {
	//TODO Выполнение самого теста
	logrus.Infof("Exec: %02x", data)
	return
}

// Возвращает данны для записи в компорт
func (s *CustomSlaveTest) ReturnData() []byte {
	// TODO подготовка данных для ответа устройству
	return []byte{0xff, 0xfe, 0xfd}
}

func (s *CustomSlaveTest) ReturnError() []byte {
	// TODO подготовка ошибки для устройства
	return []byte{0x01, 0x02, 0x03, 0x04, 0x05}
}

// Генерирует объект отчета для этого теста. Начальные данные можно использовать в сообщении before
func (s *CustomSlaveTest) GetReport() *ReportCustomSlaveTest {
	return &ReportCustomSlaveTest{
		Name:     s.Name,
		Pass:     false,
		Skip:     s.Skip,
		Write:    nil,
		Expected: nil,
		GotByte:  nil,
	}
}
