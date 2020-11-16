package cache

import (
	"bytes"
	"testing"
	"time"
)

func parse(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}

func assertContentEquals(t *testing.T, content []byte, expected []byte) {
	if !bytes.Equal(content, expected) {
		t.Errorf("coutent should %s, but was '%s'", expected, content)
	}
}

func TestGetEmpty(t *testing.T) {
	storage := NewStorage()
	content := storage.Get("My_KEY")

	assertContentEquals(t, content, []byte(""))
}

func TestGetValue(t *testing.T) {
	storage := NewStorage()
	storage.Set("My_Key", []byte("123"), parse("5s"))
	content := storage.Get("My_Key")

	assertContentEquals(t, content, []byte("123"))
}

func TestGetExpiredValue(t *testing.T) {
	storage := NewStorage()
	storage.Set("My_Key", []byte("12345"), parse("1s"))
	time.Sleep(parse("1s200ms"))
	content := storage.Get("My_Key")

	assertContentEquals(t, content, []byte(""))
}
