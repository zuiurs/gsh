package util

import (
	"fmt"
	"os"
	"strings"
)

const (
	Absolute = iota
	None
)

func Getwd(mode int) string {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if mode != Absolute {
		dir_ary := strings.Split(directory, "/")
		return dir_ary[len(dir_ary)-1]
	}

	return directory
}

func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
