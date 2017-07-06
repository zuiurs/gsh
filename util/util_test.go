package util

import (
	"os"
	"testing"
)

func TestGetwd(t *testing.T) {
	absolutePath := Getwd(Absolute)
	relativePath := Getwd(None) // relative?

	if wd, _ := os.Getwd(); absolutePath != wd {
		t.Errorf("Incorrect absolute path: %s\n", absolutePath)
	}

	if relativePath != "util" {
		t.Errorf("Incorrect relative path: %s\n", relativePath)
	}
}

func TestIsExist(t *testing.T) {
	if !IsExist("util_test.go") || IsExist("not_exist_file") {
		t.Error("not work")
	}
}
