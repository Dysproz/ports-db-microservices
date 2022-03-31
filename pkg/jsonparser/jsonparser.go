package jsonparser

import (
	"encoding/json"
	"fmt"
	"os"

	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

// Entry contains data parsed in JSON stream - key, port data and error
type Entry struct {
	Error error
	Key   string
	Port  pb.Port
}

// Stream is a type that contains stream channel used for JSON parsing
type Stream struct {
	stream chan Entry
}

// NewJSONStream returns new JSON stream
func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

// Watch watches for stream entries
func (s Stream) Watch() <-chan Entry {
	return s.stream
}

// Start starts stream parsing JSON file passed as an argument
func (s Stream) Start(path string) {
	defer close(s.stream)

	file, err := os.Open(path)
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	t, err := decoder.Token()
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("JSON token: %w", err)}
		return
	}
	if t != json.Delim('{') {
		s.stream <- Entry{Error: fmt.Errorf("expected {, got %v", t)}
		return
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			s.stream <- Entry{Error: fmt.Errorf("JSON token: %w", err)}
			return
		}
		key := t.(string)

		var port pb.Port
		if err := decoder.Decode(&port); err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode: %w", err)}
			return
		}

		s.stream <- Entry{Key: key, Port: port}
	}
	return
}
