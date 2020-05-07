package display

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
	"text/template"
	"time"
)

type Message interface {
	// Сырое не обработанное сообщение
	GetMessage() string
	// Пауза после сообщения
	GetPause() time.Duration
}

var instanceConsole *ConsoleRender
var onceConsole sync.Once

// Создает объект консоль
func Console() *ConsoleRender {
	onceConsole.Do(func() {
		instanceConsole = new(ConsoleRender)
	})
	return instanceConsole
}

type ConsoleRender struct {
	output io.Writer
}

func (c *ConsoleRender) SetOutput(f io.Writer) {
	c.output = f
}

func (c *ConsoleRender) renderTxt(tmpl string, data interface{}) string {
	t := template.Must(template.New("message").Parse(tmpl))
	buff := new(bytes.Buffer)
	if err := t.Execute(buff, data); err != nil {
		logrus.Fatal(err)
	}
	return buff.String()
}

// Печатает на консоль сообщение с отчетом
func (c *ConsoleRender) Print(message Message, report interface{}) {
	if message.GetMessage() != "" {
		resultText := c.renderTxt(message.GetMessage(), report)
		if message.GetPause() < 0 {
			resultText = fmt.Sprintf("%s [Enter]", message)
			logrus.Info(resultText)

			if c.output != nil {
				_, _ = fmt.Fprintln(c.output, resultText)
			}
			// Ожидаем нажатие клавиши Enter
			var tmp string
			_, _ = fmt.Scanln(&tmp)
		} else {
			logrus.Info(resultText)
			if c.output != nil {
				_, _ = fmt.Fprintln(c.output, resultText)
			}
		}
	}
	// Ждем указанное место
	if message.GetPause() > 0 {
		time.Sleep(message.GetPause())
	}
}
