package coco

import "os"

func open(filename string) (data []byte, err error) {
	return os.ReadFile(filename)
}

func save(filename string, data []byte) (err error) {
	return os.WriteFile(filename, data, os.ModePerm)
}
