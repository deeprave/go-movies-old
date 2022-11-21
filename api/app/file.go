package app

import (
	"fmt"
	"github.com/chigopher/pathlib"
)

type PathTypes interface {
	string | *pathlib.Path
}

// FindFile
// Locate a file in the following places
// - current directory (default)
// - directory of <appname> executable
// - in the directory $HOME/.<appname>
func FindFile(filename string) (string, error) {
	path, err := FindPath(pathlib.NewPath(filename))
	if err != nil {
		return "", err
	}
	return path.String(), err
}

// FindPath
// Locate a file in the following places, supporting Afero FS types
// - current directory (default)
// - directory of <appname> executable
// - in the directory $HOME/.<appname>
func FindPath(filepath *pathlib.Path, search ...any) (*pathlib.Path, error) {
	// try the given file directly, caller may have provided the full path
	if ok, _ := filepath.Exists(); ok {
		if ok, _ := filepath.IsDir(); !ok {
			return filepath, nil
		}
	}
	fs := filepath.Fs()

	dirs := make([]*pathlib.Path, 0, len(search)+1)
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
			p := dir.JoinPath(filepath)
			if ok, _ = p.Exists(); ok {
				return p, nil
			}
		}
	}
	return nil, fmt.Errorf("no file found '%s'", filepath.String())
}
