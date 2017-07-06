package main

import (
	"fmt"
	"os/exec"
	//"github.com/zuiurs/gsh/builtin"
	"io"
	"io/ioutil"
)

func gshExec(args []string) (io.Reader, io.Reader, error) {
	if args == nil {
		return nil, nil, nil
	}

	// separate by pipe
	cmdSets := make([][]string, 0, 10)
	set := make([]string, 0, 10)
	for _, arg := range args {
		if arg != "|" {
			set = append(set, arg)
		} else {
			cmdSets = append(cmdSets, set)
			set = make([]string, 0, 10)
		}
	}
	cmdSets = append(cmdSets, set)

	var stdout, stderr io.ReadCloser
	for i, cmdSet := range cmdSets {
		cmd := exec.Command(cmdSet[0], cmdSet[1:]...)
		stdin, err := cmd.StdinPipe()

		if i != 0 {
			stdoutBytes, err := ioutil.ReadAll(stdout)
			if err != nil {
				return nil, nil, err
			}
			stdin.Write(stdoutBytes)
			stdin.Close()
		}

		stdout, err = cmd.StdoutPipe()
		if err != nil {
			return nil, nil, err
		}
		stderr, err = cmd.StderrPipe()
		if err != nil {
			return nil, nil, err
		}
		if err = cmd.Start(); err != nil {
			return nil, nil, err
		}
	}
	// search from builtin
	//for i, cmd := range GshBuiltinCmd {
	//	if args[0] == cmd {
	//		GshBuiltinFunc[i]()
	//		return "", ""
	//	}
	//}

	//	out, err := gshLaunch(args)

	// TODO:
	// builtinコマンドの実装
	// cmdをgoroutineで接続する
	// コマンド失敗時の対処
	// コマンドのファイルが見つからないときの対処

	return stdout, stderr, nil

}

//func gshLaunch(args []string) (string, string) {
//	// search from $PATH
//	path_ary := strings.Split(os.Getenv("PATH"), ":")
//	var filepath []string
//	for _, path := range path_ary {
//		filepath = strings.Join(path, args[0])
//		if isExist(filepath) {
//
//		}
//	}
//
//}

func debug(format string, a ...interface{}) (n int, err error) {
	if Debug {
		return fmt.Printf(format, a)
	}
	return 0, nil
}
