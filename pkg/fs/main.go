package fs

import (
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func ReadDir(path string) ([]string, error) {
	info, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var size = len(info)
	var fileNames = []string{}
	for i := 0; i < size; i++ {
		fileNames = append(fileNames, info[i].Name())
	}

	return fileNames, nil
}
