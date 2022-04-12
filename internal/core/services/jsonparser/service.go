package jsonparser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)

type Entry struct {
	Error error
	Key string
	Port domain.Port
}

type stream struct {
	stream chan domain.Entry
}

// NewStream return new JSON stream
func NewStream() *stream {
	return &stream{
		stream: make(chan domain.Entry),
	}
}

// Load reads JSON file in stream
func (s stream) Load(path string) {
	file, err := os.Open(path)
	if err != nil {
		s.stream <- domain.Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	t, err := decoder.Token()
	if err != nil {
		s.stream <- domain.Entry{Error: fmt.Errorf("JSON token: %w", err)}
		return
	}
	if t != json.Delim('{') {
		s.stream <- domain.Entry{Error: fmt.Errorf("expected {, got %v", t)}
		return
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			s.stream <- domain.Entry{Error: fmt.Errorf("JSON token: %w", err)}
			return
		}
		key := t.(string)

		var port domain.Port
		if err := decoder.Decode(&port); err != nil {
			s.stream <- domain.Entry{Error: fmt.Errorf("decode: %w", err)}
			return
		}

		s.stream <- domain.Entry{Key: key, Port: port}
	}
}

// Watch watches for stream entries
func (s stream) Watch() <-chan domain.Entry {
	return s.stream
}
