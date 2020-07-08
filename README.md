# rtu-test

### Custom slave

        ---
        version: 1.0.0
        
        name: New Device
        description: "My new Device"
        log: stdout        # "off", stdout, stderr, /path/to/file
        logLvl: debug        # trace | debug | info | warn | error | fatal | panic
        
        # Сообщение выводиться при закрытии программы
        exitMessage:
          message: Для выхода нажмите
          pause: Enter
        
        slave:
          port: com5
          boundRate: 115200
          dataBits: 8
          parity: N         # Parity: N - None, E - Even, O - Odd (default E)
          stopBits: 2
          # TODO 
          # Интервал между байтами
          #silentInterval: 50ms
        
          # Порядок байт
          byteOrder: "little" # little
        
          # Не изменяемые параметры можно обращаться просто start
          const:
            start:
              - 0xFE
              - 0xFE
            addressMaster:
              - 0x00
            addressSlave:
              - 0x06
            end:
              - 0xFC
              - 0xFC
        
          # Параметры staffing байта и константы которые подлежат экранированию
          staffing:
            byte: 0x00
            # Какие данные экранируются стаффингом
            pattern:
              - start
              - end
        
          # Максимальная длина сообщения
          maxLen: 255
          
          # Если пакет ограничен размером
          len:
            # Экранировать длину staffing байтом
            staffing: true
            # считает длину с установленным staffing
            countStaffing: true
            # Длина длины в байтах 1 2 4 8
            coundBytes: 1
            read:
              - data#
            write:
              - data#
            error:
              - data#
        
          # Обращение к значениям происходит с помощью crc#write
          crc:
            # Алгоритм crc mod256, modBus
            algorithm: modBus
            # Экранировать данные staffing байтом перед подсчетом не влияет на длину
            staffing: true
            # Что входит в подсчет контрольной суммы
            read:
              - start
              - addressSlave
              - addressMaster
              - data#
            write:
              - start
              - addressSlave
              - addressMaster
              - data#
            error:
              - start
              - addressSlave
              - addressMaster
              - data#
        
          # Тут происходит не явное обработка staffing. Поля что входят в pattern не экранируются
          writeFormat:
            - start
            - addressMaster
            - addressSlave
            - data#
            - crc#
            - end
        
          # Тут происходит не явное удаление staffing. Поля что входят в pattern не преобразуются
          readFormat:
            - start
            - addressSlave
            - addressMaster
            - data#
            - crc#
            - end
        
          # Тут происходит не явное удаление staffing. Поля что входят в pattern не преобразуются
          errorFormat:
            - start
            - addressMaster
            - addressSlave
            - data#
            - crc#
            - end
        
          test:
              # название теста
            - name: Start 
            
              # Если тест пропускается
              skip: пока пропустить
        
              # Сообщение перед тестом
              before: "the message before the test"
        
              # Строги порядок выполнения команд
              next:
        
              # Если тест провален программа тестирования завершается
              fatal: exit
        
              # Количество запусков 0 - не ограниченно
              lifetime: 0
        
              # Задержка перед ответом
              timeout: 2s
        
              # Определяет что запрос пришел для этого теста. Оценка происходит только поля Data
              pattern:
                - name: "func"
                  # начиная с какого бита это значение по умолчанию 0
                  address: 0
                  uint8: 0x03
                - name: "Reg"
                  uint16: 0x0002
        
              # Параметры для записи в ответ
              write:
                - name: "param1"
                  uint16: 1
                - name: "param2"
                  uint32: 2
        
              # Возвращаем ошибку
              writeError:
                - name: "address"
                  uint8: 0x0A
                - name: "error"
                  uint16: 0x04
        
              # Проверяем поле data#
              expected:
                - name: func
                  address: 0
                  uint8: 0x03
                - name: addr
                  uint16: 0x0003
        
              success: "the message on successful completion of the test"
              error: "the message about the failed test execution"
              #after: "the message after the test"
        
            - name: Test2  # Test Name
              skip:
              before: "the message before the test"
        
              # Строги порядок выполнения команд
              next:
                - Start
        
              fatal:
        
              lifetime: 1
        
              # Задержка перед ответом
              timeout:
        
              pattern:
                - name: "func"
                  # начиная с какого бита это значение по умолчанию 0
                  address: 0
                  uint8: 0x03
                - name: "Reg"
                  uint16: 0x0002
        
              write:
                - name: "param1"
                  uint16: 1
                - name: "param2"
                  uint32: 2
        
              writeError:
                - name: "address"
                  uint8: 0x0A
                - name: "error"
                  uint16: 0x04
        
              expected:
                - name: func
                  address: 0
                  uint8: 0x03
                - name: addr
                  uint16: 0x0002
        
              success: "the message on successful completion of the test"
              error: "the message about the failed test execution"
              after: "the message after the test"