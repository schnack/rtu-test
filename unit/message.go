package unit

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"text/template"
	"time"
)

type Message struct {
	Message string `yaml:"message"`
	Pause   string `yaml:"pause"`
}

func (m *Message) Print(group string, test *ModbusTest) {
	d := parseDuration(m.Pause)

	type display struct {
		Group  string
		Test   string
		Pause  string
		Params []struct {
			Name   string
			Expect string
			Got    string
		}
	}

	a := display{Group: group, Test: test.Name, Pause: d.String()}

	for _, t := range test.Expected {
		a.Params = append(a.Params, struct {
			Name   string
			Expect string
			Got    string
		}{Name: t.Name, Expect: t.StringExpected(), Got: t.StringGot()})
	}

	if m.Message != "" {
		t := template.Must(template.New("message").Parse(m.Message))
		buff := new(bytes.Buffer)
		if err := t.Execute(buff, a); err != nil {
			logrus.Fatal(err)
		}
		Init().Display(buff.String())
	}

	if d < 0 {
		Init().Display("Press ENTER to continue...")
		var t string
		_, _ = fmt.Scanln(&t)
	} else if d > 0 {
		time.Sleep(d)
	}
}

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
