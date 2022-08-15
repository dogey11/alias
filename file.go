package main

import (
	"fmt"
	"os"
)

func write(filename string, contents string, allow_args bool, echo bool) {
	if !echo {
		contents = "@echo off\n" + contents
	}

	if allow_args {
		contents = contents + " %*"
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

func delete(filename string) {
	err := os.Remove(filename + ".cmd")
	if err != nil {
		panic(err)
	}
}
