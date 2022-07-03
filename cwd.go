package main

import "os"

func location() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return ex
}
