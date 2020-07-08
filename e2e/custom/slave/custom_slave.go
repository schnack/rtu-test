package slave

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"rtu-test/e2e/common"
	"rtu-test/e2e/custom/module"
	"rtu-test/e2e/display"
	"rtu-test/e2e/template"
	"rtu-test/e2e/transport"
	"strings"
	"time"
)

const (
	ActionRead  = "read"
	ActionWrite = "write"
	ActionError = "Error"
)

type CustomSlave struct {
	Port            string              `yaml:"port"`
	BoundRate       int                 `yaml:"boundRate"`
	DataBits        int                 `yaml:"dataBits"`
	Parity          string              `yaml:"parity"`
	StopBits        int                 `yaml:"stopBits"`
	SilentInterval  string              `yaml:"silentInterval"`
	ByteOrder       string              `yaml:"byteOrder"`
	Const           map[string][]string `yaml:"const"`
	Staffing        *module.Staffing    `yaml:"staffing"`
	MaxLen          int                 `yaml:"maxLen"`
	Len             *module.LenBytes    `yaml:"len"`
	Crc             *module.Crc         `yaml:"crc"`
	WriteFormat     []string            `yaml:"writeFormat"`
	ReadFormat      []string            `yaml:"readFormat"`
	ErrorFormat     []string            `yaml:"errorFormat"`
	CustomSlaveTest []CustomSlaveTest   `yaml:"test"`
}

// TODO Сделать проверку контрольной суммы

// Запускает тест на выполнение
// TODO тесты
func (s *CustomSlave) Run() error {
	port := transport.NewSerialPort(&transport.SerialPortConfig{
		Port:     s.Port,
		BaudRate: s.BoundRate,
		DataBits: s.DataBits,
		Parity:   s.Parity,
		StopBits: s.StopBits,
	})
	listen := bufio.NewScanner(port)
	// Собираем сканер пакетов,
	start, lenPosition, suffix, end := s.ParseReadFormat()
	// Определяем парсер пакетов по длине или по стартовым и стоповым байтам
	if lenPosition == 0 {
		listen.Split(s.GetSplitStartEnd(start, end))
	} else {
		listen.Split(s.GetSplitLen(start, lenPosition, suffix))
	}

	previousTest := ""
	// Включаем прослушку ком порта
	for listen.Scan() {
		for i := range s.CustomSlaveTest {
			// Достаем только данные
			data := s.ParseReadData(listen.Bytes())

			if s.CustomSlaveTest[i].Check(data, previousTest) {
				// Запоминаем текущий тест
				previousTest = s.CustomSlaveTest[i].Name

				// Получаем отчет для использования в сообщениях
				report := s.CustomSlaveTest[i].GetReport()
				report.GotByte = listen.Bytes()

				logrus.Warn(common.Render(template.TestSlaveCustomRUN, report))

				if report.Skip != "" {
					logrus.Warn(common.Render(template.TestSlaveCustomSKIP, report))
					continue
				}

				// Проверяем результат
				s.CustomSlaveTest[i].Exec(data, report)
				//

				// Сообщение перед тестом
				display.Console().Print(&s.CustomSlaveTest[i].Before, report)

				// определяем порядок байт
				var order binary.ByteOrder = binary.BigEndian
				if s.ByteOrder == "little" {
					order = binary.LittleEndian
				}

				// Задержка перед ответом
				duration := common.ParseDuration(s.CustomSlaveTest[i].Timeout)

				// Готовим ответ для устройства. Ошибка в приоритете
				if len(s.CustomSlaveTest[i].WriteError) > 0 {
					// Отвечаем тестируемому устройству
					out := make([]byte, 0)
					out, report.Write = s.CustomSlaveTest[i].ReturnError(order)
					out = s.GenerateAnswer(ActionError, out)
					if duration > 0 {
						logrus.Debugf("Timeout %s", duration)
						time.Sleep(duration)
					}
					logrus.Debugf("Send error: % 02x", out)
					if _, err := port.Write(out); err != nil {
						logrus.Fatalf("write answer error: %s", err.Error())
					}
				} else if len(s.CustomSlaveTest[i].Write) > 0 {
					// Отвечаем тестируемому устройству
					out := make([]byte, 0)
					out, report.Write = s.CustomSlaveTest[i].ReturnData(order)
					out = s.GenerateAnswer(ActionWrite, out)
					if duration > 0 {
						logrus.Debugf("Timeout %s", duration)
						time.Sleep(duration)
					}
					logrus.Debugf("Send answer: % 02x", out)
					if _, err := port.Write(out); err != nil {
						logrus.Fatalf("write answer error: %s", err.Error())
					}
				}

				// отчет о проделанном тесте
				if report.Pass {
					logrus.Warn(common.Render(template.TestSlaveCustomPASS, report))
					display.Console().Print(&s.CustomSlaveTest[i].Success, report)

				} else {
					logrus.Warn(common.Render(template.TestSlaveCustomFAIL, report))
					display.Console().Print(&s.CustomSlaveTest[i].Error, report)
					if s.CustomSlaveTest[i].Fatal != "" {
						logrus.Error(common.Render(template.TestSlaveCustomFATAL, report))
						return nil
					}
				}

				// Сообщение после теста
				display.Console().Print(&s.CustomSlaveTest[i].After, report)
			}
		}
	}
	return nil
}

// CalcCrc - Подсчитывает контрольную сумму согласно шаблону.
// action - read, write, error
// data - чистые данные из теста (writeError, expected, write) без staffing byte
func (s *CustomSlave) CalcCrc(action string, data []byte) []byte {
	if s.Crc == nil {
		logrus.Fatal("Crc is not specified in the configuration")
	}

	var tmpData []byte
	var format []string

	switch action {
	case ActionRead:
		format = s.ReadFormat
	case ActionWrite:
		format = s.WriteFormat
	case ActionError:
		format = s.ErrorFormat
	default:
		logrus.Fatalf("Action not found %s", action)
	}

	for _, name := range format {
		if strings.Contains(name, "#") {
			if strings.HasPrefix(name, "len#") {
				_, l := s.CalcLen(action, data)
				if s.Crc.Staffing {
					l = s.StaffingProcessing(true, l)
				}
				tmpData = append(tmpData, l...)
			}
			if strings.HasPrefix(name, "data#") {
				tmpData = append(tmpData, s.StaffingProcessing(true, data)...)
			}
			continue
		}
		if constanta, ok := s.Const[name]; ok {
			for _, stringBytes := range constanta {
				dataConst, err := common.ParseStringByte(stringBytes)
				if err != nil {
					logrus.Fatal(err)
				}
				tmpData = append(tmpData, dataConst...)
			}
		} else {
			logrus.Fatalf("Constant not found %s", constanta)
		}
	}

	// определяем порядок байт
	var order binary.ByteOrder = binary.BigEndian
	if s.ByteOrder == "little" {
		order = binary.LittleEndian
	}

	return s.Crc.Calc(order, tmpData)
}

func (s *CustomSlave) CheckCrc(action string, adu []byte) bool {
	// TODO
	return true
}

// CalcLen - Подсчитывает длину согласно шаблону
// action - read, write, error
// data - Длина в byte
func (s *CustomSlave) CalcLen(action string, data []byte) (int, []byte) {
	if s.Len == nil {
		logrus.Fatal("Length is not specified in the configuration")
	}

	countByte := 0
	var format []string

	switch action {
	case ActionRead:
		format = s.Len.Read
	case ActionWrite:
		format = s.Len.Write
	case ActionError:
		format = s.Len.Error
	default:
		logrus.Fatalf("Action not found %s", action)
	}

	for _, name := range format {
		if strings.Contains(name, "#") {
			// Подсчитываем шаблоны
			if strings.HasPrefix(name, "data#") {
				if s.Len.CountStaffing {
					countByte += len(s.StaffingProcessing(true, data))
				} else {
					countByte += len(data)
				}
			}
			continue
		}

		if constanta, ok := s.Const[name]; ok {
			for _, stringBytes := range constanta {
				dataConst, err := common.ParseStringByte(stringBytes)
				if err != nil {
					logrus.Fatal(err)
				}
				countByte += len(dataConst)
			}
		} else {
			logrus.Fatalf("Constant not found %s", constanta)
		}
	}

	// определяем порядок байт
	var order binary.ByteOrder = binary.BigEndian
	if s.ByteOrder == "little" {
		order = binary.LittleEndian
	}

	b := make([]byte, s.Len.CountBytes)

	switch s.Len.CountBytes {
	case 1:
		b[0] = uint8(countByte)
	case 2:
		order.PutUint16(b, uint16(countByte))
	case 4:
		order.PutUint32(b, uint32(countByte))
	case 8:
		order.PutUint64(b, uint64(countByte))
	default:
		logrus.Fatalf("error countByte to len %d", s.Len.CountBytes)
	}

	return countByte, b
}

// data - чистая без стаффинг байтов
// TODO тесты
func (s *CustomSlave) GenerateAnswer(action string, data []byte) (out []byte) {
	var format []string
	switch action {
	case ActionRead:
		format = s.ReadFormat
	case ActionWrite:
		format = s.WriteFormat
	case ActionError:
		format = s.ErrorFormat
	default:
		logrus.Fatalf("Action not found %s", action)
	}
	for _, templ := range format {
		if strings.Contains(templ, "#") {
			if strings.HasPrefix(templ, "len#") {
				_, l := s.CalcLen(action, data)
				if s.Len.Staffing {
					out = append(out, s.StaffingProcessing(true, l)...)
				} else {
					out = append(out, l...)
				}
				continue
			}

			// ======== Собирается суфикс ============
			if strings.HasPrefix(templ, "data#") {
				out = append(out, s.StaffingProcessing(true, data)...)
				continue
			}

			if strings.HasPrefix(templ, "crc#") {
				out = append(out, s.StaffingProcessing(true, s.CalcCrc(action, data))...)
				continue
			}
		}

		// Ищем стартовые байты в константах
		if constanta, ok := s.Const[templ]; ok {
			for _, stringBytes := range constanta {
				data, err := common.ParseStringByte(stringBytes)
				if err != nil {
					logrus.Fatal(err)
				}
				out = append(out, data...)
			}
		} else {
			logrus.Fatalf("Constant not found %s", constanta)
		}
	}
	return
}

// Возвращает чистую дату без staffing
// TODO тесты
func (s *CustomSlave) ParseReadData(adu []byte) []byte {
	adu = s.StaffingProcessing(false, adu)
	prefix := 0
	suffix := 0
	header := true
	for _, templ := range s.ReadFormat {
		if strings.Contains(templ, "#") {
			if strings.HasPrefix(templ, "len#") {
				if header {
					prefix += s.Len.CountBytes
				} else {
					suffix += s.Len.CountBytes
				}
				continue
			}

			// ======== Собирается суфикс ============
			if strings.HasPrefix(templ, "data#") {
				header = false
				continue
			}

			if strings.HasPrefix(templ, "crc#") {
				if header {
					prefix += s.Crc.Len()
				} else {
					suffix += s.Crc.Len()
				}
				continue
			}
		}
		// Ищем стартовые байты в константах
		if constanta, ok := s.Const[templ]; ok {
			for _, stringBytes := range constanta {
				data, err := common.ParseStringByte(stringBytes)
				if err != nil {
					logrus.Fatal(err)
				}
				// считаем количество байт в начале или конце пакета
				if header {
					prefix += len(data)
				} else {
					suffix += len(data)
				}
			}
		} else {
			logrus.Fatalf("Constant not found %s", constanta)
		}
	}

	if len(adu) == 0 || (len(adu)-suffix) < 0 {
		return []byte{}
	}
	if suffix == 0 {
		return adu[prefix:]
	}
	return adu[prefix:(len(adu) - suffix)]
}

// ParseReadFormat создает сплиттер для поиска фреймов в потоке данных rs
func (s *CustomSlave) ParseReadFormat() (start []byte, lenPosition int, suffix []string, end []byte) {
	prefixLen := 0
	// позволяет собирать стартовые байты
	findStart := true
	// Если суффикс сработал перед len то фаталка
	suffixTrigger := false
	for _, templ := range s.ReadFormat {
		// Если нет специальной вставки то определяем всю строку как стартовые байты
		if strings.Contains(templ, "#") {
			// =======  Собирается хедер с фиксированной длиной  ===========
			// Должен прибавлять к переменной prefixLen
			// ============================

			// Длина данные позволяет определить длину фрейма без стоповых бит
			// #len должен быть первым после констант или типов с фиксированной длиной
			if strings.HasPrefix(templ, "len#") {
				if suffixTrigger {
					logrus.Fatal("the suffix was used before len")
				}
				if s.Len == nil {
					logrus.Fatal("Data len not found in config")
				}
				lenPosition += len(start) + prefixLen
			}

			// ======== Собирается суфикс ============
			if strings.HasPrefix(templ, "data#") {
				suffixTrigger = true
				if s.Len != nil && !s.Len.Contains(ActionRead, "data#") {
					suffix = append(suffix, "data#")
				}
			}

			if strings.HasPrefix(templ, "crc#") {
				suffixTrigger = true
				if s.Len != nil && !s.Len.Contains(ActionRead, "crc#") {
					suffix = append(suffix, "crc#")
				}
			}
			// ===========================

			// Отменяем дальнейшую сборку стартовых битов
			findStart = false
			// Если конец пакета не константа то пытаемся ориентироваться по длине пакета
			end = []byte{}

			continue
		}
		// Ищем стартовые байты в константах
		if constanta, ok := s.Const[templ]; ok {
			for _, stringBytes := range constanta {
				data, err := common.ParseStringByte(stringBytes)
				if err != nil {
					logrus.Fatal(err)
				}
				if findStart {
					start = append(start, data...)
				} else {
					end = append(end, data...)
				}
			}
		} else {
			logrus.Fatalf("Constant not found %s", constanta)
		}

	}

	// Если не надйен шаблон начала пакета
	if len(start) == 0 {
		logrus.Fatal("Start byte not found")
	}
	// Если не найден шаблон конца пакета или хотябы длина
	if lenPosition == 0 && len(end) == 0 {
		logrus.Fatal("end byte or len not found")
	}
	return
}

// StaffingProcessing - Добавляет staffing byte к data
// TODO тесты
func (s *CustomSlave) StaffingProcessing(isInsert bool, data []byte) []byte {

	if s.Staffing == nil || len(s.Staffing.Byte) == 0 || len(s.Staffing.Pattern) == 0 {
		return data
	}

	staffingByte, err := common.ParseStringByte(s.Staffing.Byte)
	if err != nil || len(staffingByte) == 0 {
		logrus.Fatalf("StaffingProcessing byte error %s", err)
	}

	// Собираем шаблоны которые надо экранировать
	staffingPatterns := make(map[byte]struct{})
	for _, name := range s.Staffing.Pattern {
		if constanta, ok := s.Const[name]; ok {
			for _, stringBytes := range constanta {
				dataConst, err := common.ParseStringByte(stringBytes)
				if err != nil {
					logrus.Fatalf("StaffingProcessing parse const %s", err)
				}
				for _, c := range dataConst {
					staffingPatterns[c] = struct{}{}
				}
			}
		} else {
			logrus.Fatalf("StaffingProcessing Constant not found %s", constanta)
		}
	}

	out := make([]byte, len(data))
	copy(out, data)
	for p := range staffingPatterns {
		b := []byte{p}
		if isInsert {
			out = bytes.ReplaceAll(out, b, append(b, staffingByte...))
		} else {
			out = bytes.ReplaceAll(out, append(b, staffingByte...), b)
		}
	}
	return out
}

//func (rt *RtuTransport) SilentInterval() (frameDelay time.Duration) {
//	if rt.Config.SilentInterval.Nanoseconds() != 0 {
//		frameDelay = rt.Config.SilentInterval
//	} else if rt.BaudRate <= 0 || rt.BaudRate > 19200 {
//		frameDelay = 1750 * time.Microsecond
//	} else {
//		frameDelay = time.Duration(35000000/rt.BaudRate) * time.Microsecond
//	}
//	return
//}

// GetSplitLen - parses packets with a fixed length
// TODO Отладка
func (s *CustomSlave) GetSplitLen(start []byte, lenPosition int, suffix []string) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (int, []byte, error) {
		lenLen := 1
		if s.Len != nil && s.Len.CountBytes != 0 {
			lenLen = s.Len.CountBytes
		}
		// Если отсутствуют данные для чтения
		var err error
		if atEOF {
			err = bufio.ErrFinalToken
		}

		// Поиск стартовый байтов пакетов
		startIndex := bytes.Index(data, start)
		if startIndex < 0 {
			// Если начало пакета не найдено
			return len(data) - len(start), nil, err
		}

		// offset
		lenPosition += startIndex

		tail := lenPosition + lenLen
		// Waiting for the position length and size
		if len(data) < tail {
			return 0, nil, err
		}

		// Учитываем Staffing байт в длине
		lenCountStaffing := 0
		if s.Len != nil && s.Len.Staffing {
			lenCountStaffing = lenLen - len(s.StaffingProcessing(false, data[lenPosition:lenPosition+lenLen]))
		}
		tail += lenCountStaffing
		if len(data) < tail {
			return 0, nil, err
		}

		// Defining the byte order
		var order binary.ByteOrder = binary.BigEndian
		if s.ByteOrder == "little" {
			order = binary.LittleEndian
		}

		// Парсим длину пакета
		lengthData := 0
		switch lenLen {
		case 2:
			if s.Len != nil && s.Len.Staffing {
				lengthData = int(order.Uint16(s.StaffingProcessing(false, data[lenPosition:])[:lenLen]))
			} else {
				lengthData = int(order.Uint16(data[lenPosition : lenPosition+lenLen]))
			}
		case 4:
			if s.Len != nil && s.Len.Staffing {
				lengthData = int(order.Uint32(s.StaffingProcessing(false, data[lenPosition:])[:lenLen]))
			} else {
				lengthData = int(order.Uint32(data[lenPosition : lenPosition+lenLen]))
			}
		case 8:
			if s.Len != nil && s.Len.Staffing {
				lengthData = int(order.Uint64(s.StaffingProcessing(false, data[lenPosition:])[:lenLen]))
			} else {
				lengthData = int(order.Uint64(data[lenPosition : lenPosition+lenLen]))
			}
		default:
			lengthData = int(uint8(data[lenPosition]))
		}

		// Учитываем стаффинг байт в данных
		dataCountStaffing := 0
		if s.Len != nil && s.Len.CountStaffing {
			dataCountStaffing = lengthData - len(s.StaffingProcessing(false, data[lenPosition+lenLen+lenCountStaffing:lenPosition+lenLen+lenCountStaffing+lengthData]))
		}
		tail += lengthData + dataCountStaffing
		if len(data) < tail {
			return 0, nil, err
		}

		// Учитываем стаффинг байт Crc и конечных констант
		for _, suff := range suffix {
			switch suff {
			case "crc#":
				lenCrc := s.Crc.Len()
				tail += lenCrc
				if len(data) < tail {
					return 0, nil, err
				}
				tail += lenCrc - len(s.StaffingProcessing(false, data[lenPosition+lenLen+lenCountStaffing+lengthData+dataCountStaffing:lenPosition+lenLen+lenCountStaffing+lengthData+dataCountStaffing+lenCrc]))
			default:
				if constanta, ok := s.Const[suff]; ok {
					for _, stringBytes := range constanta {
						dataConst, err := common.ParseStringByte(stringBytes)
						if err != nil {
							logrus.Fatalf("StaffingProcessing parse const %s", err)
						}
						tail += len(dataConst)
					}
				} else {
					logrus.Fatalf("StaffingProcessing Constant not found %s", constanta)
				}
			}
		}

		// Возвращаем результат
		if len(data) < tail {
			// if this maximum packet length is exceeded
			if len(data) > s.MaxLen {
				return len(start), nil, err
			}
			return 0, nil, err
		} else {
			return tail, data[startIndex:tail], err
		}
	}
}

// Сплиттер пакета по стартовым и конечным байтам
func (s *CustomSlave) GetSplitStartEnd(start []byte, end []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (int, []byte, error) {
		// Если отсутствуют данные для чтения
		var err error
		if atEOF {
			err = bufio.ErrFinalToken
		}

		// Поиск стартовый байтов пакетов
		startIndex := bytes.Index(data, start)
		if startIndex < 0 {
			s := len(data) - len(start)
			if s < 0 {
				return 0, nil, err
			}
			logrus.Debugf("Search for start. Drop the trash: % 02x", data[:len(data)-len(start)])
			// Если начало пакета не найдено
			return len(data) - len(start), nil, err
		}

		// Поиск финальных байтов пакета
		endIndex := bytes.Index(data[len(start):], end)
		if endIndex < 0 {
			// Если данные превысили верхнюю планку пакета
			if len(data) > s.MaxLen {
				logrus.Debugf("Buffer is full. Drop the trash: % 02x", data[:len(start)])
				return len(start), nil, err
			}
			// Ждем конца пакета
			return 0, nil, err
		} else {
			tail := (len(start) + endIndex) + len(end)
			// Отбрасываем мусор перед стартовыми байтами
			if startIndex != 0 {
				logrus.Debugf("Drop the trash: % 02x", data[:startIndex])
			}

			// Если crc не прошел проверку
			if !s.CheckCrc(ActionRead, data[startIndex:tail]) {
				logrus.Debugf("Crc fail. Drop the trash: % 02x", data[:len(start)])
				return len(start), nil, err
			}

			logrus.Debugf("Found a new package: % 02x", data[startIndex:tail])
			return tail, data[startIndex:tail], err
		}
	}
}
