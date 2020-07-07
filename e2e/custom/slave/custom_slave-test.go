package slave

import (
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"math"
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
func (s *CustomSlaveTest) Check(data []byte, previousTest string) bool {
	// Если время жизни теста истекло
	if s.LifeTime < 0 {
		return false
	}

	// Соблюдаем порядок выполнения тестов
	if len(s.Next) != 0 {
		previosCheck := false
		for i := range s.Next {
			if s.Next[i] == previousTest {
				previosCheck = true
				break
			}
		}
		if !previosCheck {
			return false
		}
	}

	result := true
	var report common.ReportExpected
	offsetBit := 0
	for i := range s.Pattern {
		if s.Pattern[i].Address != "" {
			// Делаем смещение согласно заданному адресу
			rawAddress, err := strconv.Atoi(s.Pattern[i].Address)
			if err != nil {
				logrus.Fatalf("parse address %s", err)
			}
			rawAddress = int(math.Abs(float64(rawAddress)))
			if rawAddress != 0 {
				offsetBit = (rawAddress - 1) * 8
			}
		}
		offsetBit, report = s.Pattern[i].Check(data, 0, "", offsetBit, 8, binary.LittleEndian)
		if !report.Pass {
			result = false
		}
	}

	// Если значение не установленно то тест выигрывает всегда
	if s.LifeTime == 0 {
		return result
	}

	s.LifeTime--
	if s.LifeTime == 0 {
		// Если значение по умолчанию = 0
		s.LifeTime--
	}

	return result
}

// Запускает тест и поверяет значение
func (s *CustomSlaveTest) Exec(data []byte, report *ReportCustomSlaveTest) {
	//TODO Выполнение самого теста
	logrus.Infof("Exec: %02x", data)
	return
}

// Возвращает данны для записи в компорт
func (s *CustomSlaveTest) ReturnData(order binary.ByteOrder) (out []byte) {
	// TODO подготовка данных для ответа устройству
	for i := range s.Write {
		out = append(out, s.Write[i].Write(order)...)
	}
	return
}

func (s *CustomSlaveTest) ReturnError(order binary.ByteOrder) (out []byte) {
	// TODO подготовка ошибки для устройства
	for i := range s.WriteError {
		out = append(out, s.WriteError[i].Write(order)...)
	}
	return
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
