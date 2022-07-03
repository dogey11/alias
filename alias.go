package main

import (
	"fmt"
	"os"
	"strings"
)

const VERSION = "alias 1.0.0"

func main() {
	args := os.Args[1:]
	allow_args := false
	bat := false
	echo := false

	switch len(args) {
	case 0:
		os.Exit(0)
	case 1:
		switch strings.ToLower(args[0]) {
		case "-v", "--version":
			fmt.Println(VERSION)
			os.Exit(0)
		case "-p", "--path":
			fmt.Printf("Add this to your system PATH:\n%s\n", location())
			os.Exit(0)
		}
	case 2:
		break
	default:
		for i := 2; i < len(args); i++ {
			switch strings.ToLower(args[i]) {
			case "-a", "--args":
				allow_args = true
			case "-b", "--bat":
				bat = true
			case "-e", "--echo":
				echo = true
			default:
				fmt.Printf("Unknown parameter \"%s\"", strings.ToLower(args[i]))
				os.Exit(1)
			}
		}
	}

	write(args[0], args[1], allow_args, bat, echo)
}
