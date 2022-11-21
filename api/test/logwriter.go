package test

import "github.com/google/go-cmp/cmp"

// log.Logger used for testing (collecting log data)

type LogWriter struct {
	LogData [][]byte
}

func (tlw *LogWriter) Write(p []byte) (int, error) {
	tlw.LogData = append(tlw.LogData, p)
	return len(p), nil
}

func (tlw *LogWriter) Contains(v string) (int, bool) {
	p := []byte(v)
	pLen := len(p)
	for i, a := range tlw.LogData {
		aLen := len(a)
		// -1 here because log lines are automatically terminated with \n
		if cmp.Equal(a[aLen-pLen-1:aLen-1], p) {
			return i, true
		}
	}
	return -1, false
}

func NewTestLog() *LogWriter {
	return &LogWriter{
		LogData: make([][]byte, 0, 100),
	}
}
