package e2e

type MasterTest struct {
	Name       string   `yaml:"name"`
	Skip       string   `yaml:"skip"`
	Before     Message  `yaml:"before"`
	Write      []*Value `yaml:"write"`
	Expected   []*Value `yaml:"expected"`
	Success    Message  `yaml:"success"`
	Error      Message  `yaml:"error"`
	After      Message  `yaml:"after"`
	Fatal      string   `yaml:"fatal"`
	Disconnect bool     `yaml:"disconnect"`

	// Заменяет глобальные настройки
	Const       map[string][]string  `yaml:"const"`
	Staffing    *Staffing            `yaml:"staffing"`
	Len         map[string]*LenBytes `yaml:"len"`
	Crc         map[string]*Crc      `yaml:"crc"`
	WriteFormat []string             `yaml:"writeFormat"`
	ReadFormat  []string             `yaml:"readFormat"`
	ErrorFormat []string             `yaml:"errorFormat"`
}

//func (mt *MasterTest) Run(port serial.Port) ReportMasterTest {
//	// TODO Требует реализации
//
//	scanner := bufio.NewScanner(port)
//	scanner.Split(frameAmsSplit)
//	scanner.Scan()
//
//	return ReportMasterTest{Name: mt.Name, Pass: true, Skip: mt.Skip}
//}
//
//// Сканер пакетов AMS
//func frameAmsSplit(data []byte, atEOF bool) (int, []byte, error) {
//	start := -1
//	stop := -1
//
//	if len(data) < minSizePackage {
//		return 0, nil, nil
//	}
//
//	for i := 0; i < len(data); i++ {
//		switch {
//		case i == 0:
//			// Пропускаем. Для определения начала пакета нужно 2 байта
//			continue
//		case data[i] == startBit && data[i-1] == startBit:
//			// Поиск стартовых битов
//			start = i - 1
//		case data[i] == stopBit && data[i-1] == stopBit:
//			// Поиск стоповых битов
//			stop = i + 1
//		}
//
//		// Удалось собрать пакет
//		if start != -1 && stop != -1 {
//			return stop, data[start:stop], nil
//		}
//	}
//
//	// Если отсуствуют данные для чтения
//	if atEOF {
//		return 0, data[:0], bufio.ErrFinalToken
//	}
//
//	switch {
//	case start == -1:
//		// Данные - мусор очищаем буфер
//		return len(data), data[:0], nil
//	case start != -1 && stop == -1:
//		// Нашли стартовый но не нашли стоповый бит. Если собираемый пакет превышает 255 байт то удаляем
//		if (len(data) - start) > 255 {
//			return len(data), data[:0], nil
//		} else {
//			return 0, nil, nil
//		}
//	}
//
//	// Запрашиваем еще данных удаляя мусор с сохранением одного бита
//	return len(data) - 1, nil, nil
//}
