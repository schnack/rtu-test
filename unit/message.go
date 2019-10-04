package unit

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Message string `yaml:"message"`
	Pause   string `yaml:"pause"`
}

func (m *Message) Print() {
	log.Println(m.Message)
	d, t := m.parsePause()
	if d < 0 {
		log.Println("Press ENTER to continue...")
		_, _ = fmt.Scanln(&t)
	} else {
		log.Printf("Pause %d%s\n", d, t)
		time.Sleep(d)
	}
}

func (m *Message) parsePause() (time.Duration, string) {
	switch {
	case strings.HasSuffix(m.Pause, "ns"):
		s := strings.TrimSpace(strings.TrimSuffix(m.Pause, "ns"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t), "ns"
	case strings.HasSuffix(m.Pause, "us"):
		s := strings.TrimSpace(strings.TrimSuffix(m.Pause, "us"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Microsecond, "us"
	case strings.HasSuffix(m.Pause, "ms"):
		s := strings.TrimSpace(strings.TrimSuffix(m.Pause, "ms"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Millisecond, "ms"
	case strings.HasSuffix(m.Pause, "s"):
		s := strings.TrimSpace(strings.TrimSuffix(m.Pause, "s"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Second, "s"
	case strings.HasSuffix(m.Pause, "m"):
		s := strings.TrimSpace(strings.TrimSuffix(m.Pause, "m"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Minute, "m"
	case strings.HasSuffix(m.Pause, "h"):
		s := strings.TrimSpace(strings.TrimSuffix(m.Pause, "h"))
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Hour, "h"
	default:
		s := strings.TrimSpace(m.Pause)
		t, err := strconv.Atoi(s)
		if err != nil {
			return time.Duration(-1), ""
		}
		return time.Duration(t) * time.Second, "s"
	}
}
