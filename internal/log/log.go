package log

// Log
type Log struct {
	writer func(string)
}

// Write
func (f *Log) Write(p []byte) (n int, err error) {
	f.writer(string(p))
	return 0, nil
}

// NewLog
func NewLog(writer func(string)) *Log {
	return &Log{
		writer: writer,
	}
}
