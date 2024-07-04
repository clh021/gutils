package exec

import (
	"os"
	"path/filepath"
)


func GetProgramPath() string {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		panic(err)
	}
	// fmt.Println("exReal: ", exReal)
	return filepath.Dir(exReal)
}