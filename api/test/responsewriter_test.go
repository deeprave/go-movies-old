package test

import (
	"testing"
)

func TestResponseWriter_Header(t *testing.T) {
	rw := NewResponseWriter()
	header := "some-header"
	expectedValue := "some value"
	rw.Header().Set(header, expectedValue)
	value := rw.Header().Get("Some-Header") // Capitalised
	ShouldBeEqual(t, expectedValue, value)
	anotherExpectedValue := "some other value"
	rw.Header().Set(header, anotherExpectedValue)
	value = rw.Header().Get(header)
	ShouldBeEqual(t, anotherExpectedValue, value)
	bothExpectedValues := []string{expectedValue, anotherExpectedValue}
	rw.Header().Add(header, expectedValue)
	values := rw.Header().Values(header)
	ShouldBeEqual(t, 2, len(values))
	ShouldBeInArray(t, bothExpectedValues, expectedValue)
	ShouldBeInArray(t, bothExpectedValues, anotherExpectedValue)
}

func TestResponseWriter_Write_and_Body(t *testing.T) {
	rw := NewResponseWriter()
	expectedBody := "This is the response body"
	expectedBodyBytes := []byte(expectedBody)
	l, err := rw.Write(expectedBodyBytes)
	ShouldBeNoError(t, err, "unexpected Write error: %v", err)
	ShouldBeEqual(t, len(expectedBody), l)
	ShouldBeEqual(t, expectedBody, rw.BodyAsString())
	ShouldBeEqual(t, expectedBodyBytes, rw.BodyAsBytes())
}

func TestResponseWriter_Dump(t *testing.T) {
	// rw := NewResponseWriter()
}

////goland:noinspection GoUnhandledErrorResult
//func main() {
//	rw.Header().Set("Content-Type", "application/json")
//	buf := new(bytes.Buffer)
//	rw.Write([]byte("Hello there\n"))
//	rw.Write([]byte("This is the body of this message, let's see how it goes...\n"))
//	rw.Dump(buf)
//	fmt.Print(buf.String())
//}
