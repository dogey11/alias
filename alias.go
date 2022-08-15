package main

import (
	"fmt"
	"os"
	"strings"
)

const VERSION = "1.1.0"

func parseArgs(arg string) bool {
	for _, s := range os.Args[1:] {
		if s == arg {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 0:
		os.Exit(0)
	case 1:
		switch strings.ToLower(args[0]) {
		case "-v":
			fmt.Println(VERSION)
			os.Exit(0)
		case "-p":
			fmt.Printf("Add this to your system PATH:\n%s\n", location())
			os.Exit(0)
		default:
			fmt.Printf("Unknown parameter \"%s\"", strings.ToLower(args[0]))
			os.Exit(1)
		}
	case 2:
		switch strings.ToLower(args[0]) {
		case "-d":
			delete(args[1])
			os.Exit(0)
		default:
			break
		}
	default:
		for i := 2; i < len(args); i++ {
			if strings.ToLower(args[i]) != "-a" && strings.ToLower(args[i]) != "-e" {
				fmt.Printf("Unknown parameter \"%s\"", strings.ToLower(args[i]))
				os.Exit(1)
			}
		}
	}

	write(args[0]+".cmd", args[1], parseArgs("-a"), parseArgs("-e"))
}
