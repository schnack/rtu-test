package e2e

import (
	"bufio"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"rtu-test/e2e/common"
	common2 "rtu-test/e2e/custom/common"
	"rtu-test/e2e/transport"
	"strings"
)

const (
	ActionRead  = "read"
	ActionWrite = "write"
	ActionError = "Error"
)

type CustomSlave struct {
	Port           string              `yaml:"port"`
	BoundRate      int                 `yaml:"boundRate"`
	DataBits       int                 `yaml:"dataBits"`
	Parity         string              `yaml:"parity"`
	StopBits       int                 `yaml:"stopBits"`
	SilentInterval string              `yaml:"silentInterval"`
	ByteOrder      string              `yaml:"byteOrder"`
	Const          map[string][]string `yaml:"const"`
	Staffing       *common2.Staffing   `yaml:"staffing"`
	MaxLen         int                 `yaml:"maxLen"`
	Len            *LenBytes           `yaml:"len"`
	Crc            *common2.Crc        `yaml:"crc"`
	WriteFormat    []string            `yaml:"writeFormat"`
	ReadFormat     []string            `yaml:"readFormat"`
	ErrorFormat    []string            `yaml:"errorFormat"`
	SlaveTest      []CustomSlaveTest   `yaml:"test"`
}

// ParseReadFormat создает сплиттер для поиска фреймов в потоке данных rs
func (s *CustomSlave) ParseReadFormat() (start []byte, lenPosition, suffixLen int, end []byte) {
	prefixLen := 0
	// позволяет собирать стартовые байты
	findStart := true
	// Если суффикс сработал перед len то фаталка
	suffixTrigger := false
	for _, data := range s.ReadFormat {
		// Если нет специальной вставки то определяем всю строку как стартовые байты
		if strings.Contains(data, "#") {
			// =======  Собирается хедер с фиксированной длиной  ===========
			// Должен прибавлять к переменной prefixLen
			// ============================

			// Длина данные позволяет определить длину фрейма без стоповых бит
			// #len должен быть первым после констант или типов с фиксированной длиной
			if strings.HasPrefix(data, "len#") {
				if suffixTrigger {
					logrus.Fatal("the suffix was used before len")
				}
				if s.Len == nil {
					logrus.Fatal("Data len not found in config")
				}
				lenPosition += len(start) + prefixLen
			}

			// ======== Собирается суфикс ============
			if strings.HasPrefix(data, "data#") {
				suffixTrigger = true
				if !s.Len.Contains(ActionRead, "data#") {
					suffixLen += 0
				}
			}

			if strings.HasPrefix(data, "crc#") {
				suffixTrigger = true
				if !s.Len.Contains(ActionRead, "crc#") {
					suffixLen += s.Crc.Len()
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
		if constanta, ok := s.Const[data]; ok {
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
	if lenPosition == 0 || len(end) == 0 {
		logrus.Fatal("end byte or len not found")
	}
	return
}

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
	start, lenPosition, suffixLen, end := s.ParseReadFormat()
	listen.Split(s.GetSplit(start, lenPosition, suffixLen, end))

	// Включаем прослушку ком порта
	for listen.Scan() {
		for i := range s.SlaveTest {
			if s.SlaveTest[i].Check(listen.Bytes()) {
				//s.CustomSlaveTest[i].Exec(listen.Bytes(), )
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
			if strings.HasPrefix(name, "len#") {
				_, l := s.CalcLen(action, data)
				tmpData = append(tmpData, l...)
			}
			if strings.HasPrefix(name, "data#") {
				tmpData = append(tmpData, data...)
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

	return s.Crc.Calc(tmpData)
}

// CalcLen - Подсчитывает длину согласно шаблону
// action - read, write, error
// data - чистые данные из теста (writeError, expected, write) без staffing byte
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
				if s.Len.Staffing {
					countByte += len(s.AddStaffing(data))
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

// AddStaffing -  Добавляет staffing byte к data
func (s *CustomSlave) AddStaffing(data []byte) (out []byte) {

	if s.Staffing == nil || len(s.Staffing.Byte) == 0 || len(s.Staffing.Pattern) == 0 {
		return data
	}

	staffingByte, err := common.ParseStringByte(s.Staffing.Byte)
	if err != nil || len(staffingByte) == 0 {
		logrus.Fatalf("Staffing byte error %s", err)
	}

	// Собираем шаблоны которые надо экранировать
	var staffingPatterns []byte
	for _, name := range s.Staffing.Pattern {
		if constanta, ok := s.Const[name]; ok {
			for _, stringBytes := range constanta {
				dataConst, err := common.ParseStringByte(stringBytes)
				if err != nil {
					logrus.Fatal(err)
				}
				staffingPatterns = append(staffingPatterns, dataConst...)
			}
		} else {
			logrus.Fatalf("Constant not found %s", constanta)
		}
	}

	// Добавляем стаффинг байт если попался символ из шаблона
	for _, d := range data {
		isAddStuffing := false
		for _, p := range staffingPatterns {
			if d == p {
				isAddStuffing = true
				break
			}
		}
		if isAddStuffing {
			out = append(out, d, staffingByte[0])
		} else {
			out = append(out, d)
		}
	}
	return
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

// GetSplit - Функция пытается собрать фрейм на основе данных из ReadFormat
// Start - фиксированные байты
// lenPosition - позиция байтов длины
// suffixLen - Длина окончания которое не входит в len
// end - филированная концовка
func (s *CustomSlave) GetSplit(start []byte, lenPosition, suffixLen int, end []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (int, []byte, error) {

		positionEnd := 0
		// Если мы можем определить длину пакета
		if lenPosition > 0 {
			if len(data) < lenPosition+s.Len.CountBytes {
				// Если отсуствуют данные для чтения
				if atEOF {
					return 0, data[:0], bufio.ErrFinalToken
				}
				return 0, nil, nil
			}
			// определяем порядок байт
			var order binary.ByteOrder = binary.BigEndian
			if s.ByteOrder == "little" {
				order = binary.LittleEndian
			}
			// TODO надо тестировать возможно ошибся в размерах
			// Определяем длину
			length := 0

			if s.Len.Staffing {
				// TODO тут нужно убить stafing byte перед определением длины
			}
			switch s.Len.CountBytes {
			case 2:
				length = int(order.Uint16(data[lenPosition : lenPosition+s.Len.CountBytes]))
			case 4:
				length = int(order.Uint32(data[lenPosition : lenPosition+s.Len.CountBytes]))
			case 8:
				length = int(order.Uint64(data[lenPosition : lenPosition+s.Len.CountBytes]))
			default:
				length = int(uint8(data[lenPosition]))
			}

			if len(data) < length+suffixLen+len(end) {
				// Если отсуствуют данные для чтения
				if atEOF {
					return 0, data[:0], bufio.ErrFinalToken
				}
				return 0, nil, nil
			}
			positionEnd = lenPosition + length + suffixLen + 1
		}

		// Проверяем что в буфере достаточно данных для начала пакета
		if len(data) < len(start)+len(end) {
			// Если отсуствуют данные для чтения
			if atEOF {
				return 0, data[:0], bufio.ErrFinalToken
			}
			return 0, nil, nil
		}

		// Проверяем начало пакета
		for i := range start {
			if data[i] != start[i] {
				// Если отсуствуют данные для чтения
				if atEOF {
					return 0, data[:0], bufio.ErrFinalToken
				}
				return i + 1, nil, nil
			}
		}

		// Мы незнаем где точно должна быть финальная константа
		if positionEnd == 0 {
			// курсор end
			pos := 0
			for i := len(start); i < len(data) && i < s.MaxLen; i++ {
				if data[i] == end[pos] {
					// если есть совпадения то плюсуем
					pos++
				} else {
					// если не совпало то возможно это данные -> двигаемся дальше
					pos = 0
				}
				// нашли последний финальный хвост -> возвращаем фрейм
				if pos == len(end) {
					// Фрейм полностью собран
					if atEOF {
						return i, data[:i], bufio.ErrFinalToken
					}
					return i, data[:i], nil
				}
			}
			// Если данные превысили верхнюю планку пакета
			if len(data) > s.MaxLen {
				// Если отсуствуют данные для чтения
				if atEOF {
					return 0, data[:0], bufio.ErrFinalToken
				}
				return 1, nil, nil
			}
			// Если отсуствуют данные для чтения
			if atEOF {
				return 0, data[:0], bufio.ErrFinalToken
			}
			// Пока не нашли конец пакета... ждем продолжения
			return 0, nil, nil
		} else {
			// Знаем точное наступление end остается только сверить корректность end
			for i := range end {
				// Ошибка в окончании пытаемся сдвинуться на 1 байт дальше
				if data[positionEnd+i] != end[i] {
					// Если отсуствуют данные для чтения
					if atEOF {
						return 0, data[:0], bufio.ErrFinalToken
					}
					return 1, nil, nil
				}
			}

			// Фрейм полностью собран
			if atEOF {
				return positionEnd + len(end), data[:positionEnd+len(end)], bufio.ErrFinalToken
			}
			return positionEnd + len(end), data[:positionEnd+len(end)], nil
		}

	}
}
