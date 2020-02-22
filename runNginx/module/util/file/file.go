package file

import (
	"io/ioutil"
	"os"
	"strings"
)

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func IsFile(f string) bool {
	fi, err := os.Stat(f)
	return err == nil || !fi.IsDir()
}

func FindTag(dir string, tag string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, onefile := range files {
		if strings.Contains(onefile.Name(), tag) {
			return true
		}
	}
	return false
}
