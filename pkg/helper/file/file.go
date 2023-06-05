package file

import (
	"hitss/pkg/helper/logger"
	"io/ioutil"
)

func GetNames(path string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		logger.Write(err)
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
