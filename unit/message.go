package unit

import (
	"fmt"
	"log"
	"time"
)

type Message struct {
	Message string `yaml:"message"`
	Pause   string `yaml:"pause"`
}

func (m *Message) Print() {
	log.Println(m.Message)
	d := parseDuration(m.Pause)
	if d < 0 {
		log.Println("Press ENTER to continue...")
		var t string
		_, _ = fmt.Scanln(&t)
	} else {
		log.Printf("Pause %s\n", d.String())
		time.Sleep(d)
	}
}

func (m *Message) parsePause() time.Duration {
	return parseDuration(m.Pause)
}
