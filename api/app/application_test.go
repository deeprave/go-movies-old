package app

import (
	"github.com/google/go-cmp/cmp"
	"go-movies/api/helpers"
	"testing"
)

// test app creation
func TestNewApplication(t *testing.T) {
	testLogger := helpers.NewTestLog()
	cfg := helpers.NewAppConfig()
	app := NewApplication(cfg, testLogger, "test ")
	if !cmp.Equal(app.Config, cfg) {
		t.Error("app config is not the same as passed config")
	}
	s := "this is a log line"
	app.Logger.Print(s)
	if _, result := testLogger.Contains(s); !result {
		t.Error("expected log entry is not created")
	}
}

func TestModelToJson(t *testing.T) {
	app := NewApplication(helpers.NewAppConfig(), helpers.NewTestLog(), "test ")
	app = app
}

func TestErrorToJson(t *testing.T) {
	t.Fail()
}
