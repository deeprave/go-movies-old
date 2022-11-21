package app

import (
	"github.com/chigopher/pathlib"
	"github.com/spf13/afero"
	"go-movies/api/test"
	"testing"
)

func TestFindPath(t *testing.T) {
	home := "/home/testing"

	fs := afero.NewMemMapFs()

	// make a "home" path
	homePath := pathlib.NewPathAfero(home, fs)
	err := homePath.MkdirAll()
	test.ShouldBeNoError(t, err, "%v: %v", homePath, err)
	ok, err := homePath.IsDir()
	test.ShouldBeTrue(t, ok, "%v: %v", homePath, err)

	// create a file in the home path with some data
	configFile := pathlib.NewPathAfero("config.yml", fs)
	configPath := homePath.JoinPath(configFile)
	data := []byte("test:\n\t- first item\n\t- second item\n")
	err = configPath.WriteFile(data)
	test.ShouldBeNoError(t, err, "%v: %v", configFile, err)

	// how try to locate it
	foundPath, err := FindPath(configFile, ".", homePath)
	test.ShouldBeNoError(t, err, "%v", err)
	size, err := foundPath.Size()
	test.ShouldBeNoError(t, err, "%v (size): %v", configFile, err)
	test.ShouldBeEqual(t, len(data), int(size))
	t.Logf("successfully found: %v, %d bytes", foundPath, size)
}
