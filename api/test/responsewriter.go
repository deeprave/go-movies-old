package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// http.ResponseWriter for testing

type ResponseWriter struct {
	headers    http.Header
	body       *bytes.Buffer
	error      error
	status     int
	statusText string
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		headers:    make(map[string][]string),
		body:       new(bytes.Buffer),
		error:      nil,
		status:     http.StatusOK,
		statusText: http.StatusText(http.StatusOK),
	}
}

func (rw ResponseWriter) Header() http.Header {
	return rw.headers
}

func (rw ResponseWriter) SetError(err error) {
	rw.error = err
}

func (rw ResponseWriter) SetStatus(status int) {
	rw.status = status
}

func (rw ResponseWriter) Write(bytes []byte) (int, error) {
	if rw.error == nil {
		rw.body.Write(bytes)
		rw.headers.Set("Content-Length", strconv.Itoa(rw.body.Len()))
		return len(bytes), nil
	}
	return -1, rw.error
}

func (rw ResponseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.statusText = http.StatusText(statusCode)
}

func (rw ResponseWriter) BodyAsBytes() []byte {
	return rw.body.Bytes()
}

func (rw ResponseWriter) BodyAsString() string {
	return rw.body.String()
}

//goland:noinspection GoUnhandledErrorResult
func (rw ResponseWriter) Dump(out io.Writer) {
	fmt.Fprintf(out, "%d %s\n", rw.status, rw.statusText)
	rw.headers.Write(out)
	fmt.Fprintln(out)
	fmt.Fprint(out, rw.BodyAsString())
	fmt.Fprintln(out)
	if rw.error != nil {
		fmt.Fprintf(out, "Response Error: %v", rw.error)
	}
}
