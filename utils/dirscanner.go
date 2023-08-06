package utils

import "os"

func DirScanner(path string) ([]string, error) {
	var res []string = make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		return res, err
	}
	files, err := file.ReadDir(-1)
	for _, value := range files {
		res = append(res, value.Name())
	}
	return res, err
}
