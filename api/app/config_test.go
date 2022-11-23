package app

import (
	"go-movies/api/test"
	"log"
	"testing"
)

func TestNewConfigFromFile(t *testing.T) {
	cfg, err := NewConfigFromFile("config.yml", "1")
	if err != nil {
		log.Fatalf("%v", err)
	}
	test.ShouldBeTrue(t, cfg.VersionOk("1"), "incorrect version")
	test.ShouldBeEqual(t, cfg.Env, "dev")
	test.ShouldBeEqual(t, cfg.Host, "localhost")
	test.ShouldBeEqual(t, cfg.Port, 4000)
}
