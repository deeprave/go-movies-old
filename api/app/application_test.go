package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-movies/api/test"
	"io"
	"net/http"
	"testing"
)

func MakeTestApplication(testLog io.Writer) *Application {
	cfg := &Config{
		Port:    9000,
		Host:    "localhost",
		Env:     "dev",
		Version: "1.0",
	}
	return NewApplication(cfg, testLog, "test ")
}

// test app creation
func TestNewApplication(t *testing.T) {
	testLog := test.NewTestLog()
	app := MakeTestApplication(testLog)
	logText := "this is a log line"
	app.Log(logText)
	if _, result := testLog.Contains(logText); !result {
		t.Error("expected log entry is not present")
	}
}

type Record struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func TestModelToJson(t *testing.T) {
	app := MakeTestApplication(test.NewTestLog())
	// func (app *Application) ModelToJson(w http.ResponseWriter, status int, data interface{}, wrap string) (int, error)
	records := []Record{
		{Id: 14, Name: "Alpha", Description: "This is Alpha"},
		{Id: 22, Name: "Beta", Description: "This is Beta"},
		{Id: 39, Name: "Gamma", Description: "This is Gamma"},
	}
	var (
		status       int
		err          error
		buffer       = new(bytes.Buffer)
		responseText string
		responseBody []byte
		rw           = test.NewResponseWriter()
	)
	status, err = app.ModelToJson(rw, http.StatusOK, records, "data")
	test.ShouldBeNoError(t, err, "Unexpected error in ModelToJson: %v", err)
	test.ShouldBeEqual(t, status, http.StatusOK)
	rw.Dump(buffer)
	responseText = buffer.String()
	test.ShouldBeFalse(t, len(responseText) <= 0, "response headers + body has invalid length")
	test.ShouldBeSubstring(t, responseText, "Content-Type")
	test.ShouldBeSubstring(t, responseText, "Content-Length")

	responseBody = rw.BodyAsBytes()
	data := make(map[string][]Record)
	err = json.Unmarshal(responseBody, &data)
	test.ShouldBeNoError(t, err, "error unmarshalling response: %v", err)
	test.ShouldBeEqual(t, records, data["data"])
}

func TestErrorToJson(t *testing.T) {
	app := MakeTestApplication(test.NewTestLog())
	errText := "this is an error message"
	customError := errors.New(errText)
	rw := test.NewResponseWriter()
	status, err := app.ErrorToJson(rw, http.StatusInternalServerError, customError)
	test.ShouldBeNoError(t, err, "error sending error response: %v", err)
	test.ShouldBeEqual(t, http.StatusOK, status)
	responseBody := []byte(rw.BodyAsString())
	data := make(map[string]map[string]string)
	err = json.Unmarshal(responseBody, &data)
	test.ShouldBeNoError(t, err, "error unmarshalling response: %v", err)
	test.ShouldBeEqual(t, errText, data["error"]["message"])
}
