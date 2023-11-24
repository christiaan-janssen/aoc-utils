package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileToString(fileName string) {
	data, err := os.ReadFile(fileName)
	check(err)

	fmt.Print(string(data))
}
