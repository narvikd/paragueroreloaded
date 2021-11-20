package osutils

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func GetExecutableDir() (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", errors.Wrap(err, "couldn't find the path of the current executable on getExecutableDir")
	}
	return filepath.Dir(executable), nil
}

// GetExecAbsPathFile gets the absolute path + the filename passed as a parameter from the current directory where the executable is
func GetExecAbsPathFile(fileName string) (string, error) {
	basePath, err := GetExecutableDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(basePath, fileName), nil
}
