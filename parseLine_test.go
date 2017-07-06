package main

import (
	"reflect"
	"testing"
)

func TestGshParseLine(t *testing.T) {
	testsNum := 6
	testStrings := make([]string, testsNum)
	testResults := make([][]string, testsNum)

	testStrings[0] = `echo hello`
	testResults[0] = []string{"echo", "hello"}

	testStrings[1] = `     echo  -n hello    world   `
	testResults[1] = []string{"echo", "-n", "hello", "world"}

	testStrings[2] = `sh -c "echo hello"`
	testResults[2] = []string{"sh", "-c", "echo hello"}

	testStrings[3] = `sh -c 'echo hello    world  '   `
	testResults[3] = []string{"sh", "-c", "echo hello    world  "}

	testStrings[4] = `sh -c "echo 'hello world'"`
	testResults[4] = []string{"sh", "-c", "echo 'hello world'"}

	testStrings[5] = `sh -c "echo \"hello\""`
	testResults[5] = []string{"sh", "-c", `echo "hello"`}

	for i, str := range testStrings {
		if !reflect.DeepEqual(gshParseLine(str), testResults[i]) {
			t.Errorf("parse result is not correct: %#v\n", gshParseLine(str))
		}
	}
}
