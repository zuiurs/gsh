package main

import (
	"bufio"
	"os"
)

func gshReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func gshParseLine(line string) []string {
	var args []string

	tmpstr := make([]rune, 0, 20)

	for _, c := range line {
		if c == ' ' {
			// remove forefront or continual spaces
			if len(tmpstr) == 0 {
				continue
			} else {
				args = append(args, string(tmpstr))
				tmpstr = make([]rune, 0, 20)
			}
		} else {
			tmpstr = append(tmpstr, c)
		}
	}
	if len(tmpstr) != 0 {
		args = append(args, string(tmpstr))
	}

	return args
}
