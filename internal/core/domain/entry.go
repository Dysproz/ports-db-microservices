package domain

// Entry is a base JSON stream data entry
type Entry struct {
	Error error
	Key   string
	Port  Port
}
