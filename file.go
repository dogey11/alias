package main

import (
	"fmt"
	"os"
	"strings"
)

func write(filename string, contents string, allow_args bool, bat bool, echo bool) {
	if !echo {
		contents = "@echo off\n" + contents
	}

	if allow_args {
		contents = contents + " %*"
	}

	if bat {
		filename = fmt.Sprintf("%s.bat", strings.ReplaceAll(filename, " ", ""))
	} else {
		filename = fmt.Sprintf("%s.cmd", strings.ReplaceAll(filename, " ", ""))
	}

	os.Chdir(location())
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(contents)
	if err != nil {
		fmt.Println(err)
		f.Close()
	}
}
