package helpers

import "github.com/google/go-cmp/cmp"

// Testing logger

type Writer interface {
	Write(p []byte) (n int, err error)
}

type TestLogWriter struct {
	LogData [][]byte
}

func (tlw *TestLogWriter) Write(p []byte) (int, error) {
	tlw.LogData = append(tlw.LogData, p)
	return len(p), nil
}

func (tlw *TestLogWriter) Contains(v string) (int, bool) {
	p := []byte(v)
	plen := len(p)
	for i, a := range tlw.LogData {
		alen := len(a)
		// -1 here because log lines are automatically terminated with \n
		if cmp.Equal(a[alen-plen-1:alen-1], p) {
			return i, true
		}
	}
	return -1, false
}

func NewTestLog() *TestLogWriter {
	return &TestLogWriter{
		LogData: make([][]byte, 0, 100),
	}
}
