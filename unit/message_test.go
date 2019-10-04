package unit

import (
	"github.com/schnack/gotest"
	"testing"
	"time"
)

func TestMessage_parsePauseNs(t *testing.T) {
	d, tp := (&Message{Pause: "1 ns"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(1)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("ns"); err != nil {
		t.Error(err)
	}
}

func TestMessage_parsePauseUs(t *testing.T) {
	d, tp := (&Message{Pause: "1 us"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Microsecond); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("us"); err != nil {
		t.Error(err)
	}
}

func TestMessage_parsePauseMs(t *testing.T) {
	d, tp := (&Message{Pause: "1 ms"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Millisecond); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("ms"); err != nil {
		t.Error(err)
	}
}

func TestMessage_parsePauseS(t *testing.T) {
	d, tp := (&Message{Pause: "1 s"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Second); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("s"); err != nil {
		t.Error(err)
	}
}

func TestMessage_parsePauseM(t *testing.T) {
	d, tp := (&Message{Pause: "1 m"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Minute); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("m"); err != nil {
		t.Error(err)
	}
}

func TestMessage_parsePauseH(t *testing.T) {
	d, tp := (&Message{Pause: "1 h"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Hour); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("h"); err != nil {
		t.Error(err)
	}
}

func TestMessage_parsePause(t *testing.T) {
	d, tp := (&Message{Pause: "1"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(1) * time.Second); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq("s"); err != nil {
		t.Error(err)
	}
}

func TestMessage_parsePauseEnter(t *testing.T) {
	d, tp := (&Message{Pause: "enter"}).parsePause()
	if err := gotest.Expect(d).Eq(time.Duration(-1)); err != nil {
		t.Error(err)
	}

	if err := gotest.Expect(tp).Eq(""); err != nil {
		t.Error(err)
	}
}
