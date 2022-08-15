package main

import (
	"fmt"
	"os"
	"strings"
)

const VERSION = "1.1.0"

func parseArgs(arg string, arg_alt string) bool {
	for _, s := range os.Args[1:] {
		if s == arg || s == arg_alt {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	allow_args := false
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
		default:
			os.Exit(0)
		}
	case 2:
		switch strings.ToLower(args[0]) {
		case "-d", "--delete":
			delete(args[1])
			os.Exit(0)
		default:
			break
		}
	default:
		for i := 2; i < len(args); i++ {
			switch strings.ToLower(args[i]) {
			case "-a", "--args":
				allow_args = true
			case "-e", "--echo":
				echo = true
			default:
				fmt.Printf("Unknown parameter \"%s\"", strings.ToLower(args[i]))
				os.Exit(1)
			}
		}
	}

	write(args[0], args[1], parseArgs("-a", ""), echo)
}
