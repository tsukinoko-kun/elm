package elm

import (
	"os"
	"os/exec"
	"path/filepath"
)

var (
	f *os.File
)

const Version = "0.19.1"

func Init() error {
	p, err := os.MkdirTemp("", "elm-*")
	if err != nil {
		return err
	}

	f, err = os.Create(filepath.Join(p, name))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(file)
	if err != nil {
		return err
	}

	return os.Chmod(f.Name(), 0755)
}

func Dispose() error {
	return os.Remove(f.Name())
}

func Command(args ...string) *exec.Cmd {
	return exec.Command(f.Name(), args...)
}
