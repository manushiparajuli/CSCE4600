package builtins

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var (
	ErrInvalidArgCnt = errors.New("invalid argument count")
)

func ListDirectory(args ...string) error {
	var path string
	switch len(args) {
	case 0:
		path = "."
	case 1:
		path = args[0]
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCnt)
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(filepath.Join(path, file.Name()))
	}
	return nil
}
