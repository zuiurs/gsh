package main

import (
	"fmt"
	"github.com/zuiurs/gsh/util"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := gshLoop(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

const (
	IsFullPath int  = util.Absolute
	Debug      bool = true
)

func gshLoop() error {
	var line string
	var args []string

	for {
		fmt.Print(util.Getwd(IsFullPath), "> ")

		line = gshReadLine()
		args = gshParseLine(line)
		stdout, stderr, err := gshExec(args)
		if err != nil {
			return err
		}
		if stdout != nil {
			stdoutBytes, err := ioutil.ReadAll(stdout)
			if err != nil {
				return err
			}
			fmt.Println(string(stdoutBytes))
		}
		if stderr != nil {
			stderrBytes, err := ioutil.ReadAll(stderr)
			if err != nil {
				return err
			}
			fmt.Println(string(stderrBytes))
		}

		line = ""
		args = nil
	}
}
