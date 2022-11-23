package app

import (
	"fmt"
	"github.com/chigopher/pathlib"
	"github.com/spf13/afero"
)

func FindFs(v ...any) (afero.Fs, int) {
	var fs afero.Fs = afero.OsFs{}
	var index = 0
	if len(v) > 0 { // default fs
		var ok bool
		if fs, ok = v[0].(afero.Fs); ok {
			index++
		}
	}
	return fs, index
}

// FindFile
// Locate a file in the following places
// - current directory (default)
// - directory of <appname> executable
// - in the directory $HOME/.<appname>
func FindFile(filename string, v ...any) (string, error) {
	fs, index := FindFs(v...)
	filepath := pathlib.NewPathAfero(filename, fs)
	if len(v) > index {
		v = v[index:]
	}
	path, err := FindPath(filepath, v...)
	if err != nil {
		return "", err
	}
	return path.String(), err
}

func FindFileExists(path *pathlib.Path) bool {
	if ok, _ := path.Exists(); ok {
		if ok, _ := path.IsFile(); ok {
			return true
		}
	}
	return false
}

// FindPath
// Locate a file in the following places, supporting Afero FS types
// - current directory (default)
// - directory of <appname> executable
// - in the directory $HOME/.<appname>
func FindPath(path *pathlib.Path, search ...any) (*pathlib.Path, error) {
	// try the given file directly, caller may have provided the full path
	if FindFileExists(path) {
		return path, nil
	}

	fs := path.Fs()
	dirs := make([]*pathlib.Path, 0, len(search))
	if len(search) == 0 {
		dirs = append(dirs, pathlib.NewPathAfero(".", fs))
	} else {
		for _, d := range search {
			if s, ok := d.(*pathlib.Path); ok {
				dirs = append(dirs, s)
			} else if s, ok := d.(string); ok {
				dirs = append(dirs, pathlib.NewPathAfero(s, fs))
			}
		}
	}
	for _, dir := range dirs {
		if ok, _ := dir.IsDir(); ok {
			p := dir.JoinPath(path)
			if FindFileExists(p) {
				return p, nil
			}
		}
	}
	return nil, fmt.Errorf("no file found '%s'", path.String())
}
