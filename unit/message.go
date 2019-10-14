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
	d, t := parseDuration(m.Pause)
	if d < 0 {
		log.Println("Press ENTER to continue...")
		_, _ = fmt.Scanln(&t)
	} else {
		log.Printf("Pause %d%s\n", d, t)
		time.Sleep(d)
	}
}

func (m *Message) parsePause() (time.Duration, string) {
	return parseDuration(m.Pause)
}
