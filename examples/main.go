package main

import (
	"fmt"
	"os"
	"strconv"
)

type example struct {
	name        string
	description string
	fn          func()
}

var examples []example

func registerExample(name, description string, fn func()) {
	examples = append(examples, example{name: name, description: description, fn: fn})
}

func main() {
	if len(os.Args) == 1 {
		println("usage: examples <example_number>")
		println("available examples:")
		for i, e := range examples {
			fmt.Printf(" %d) %-20s - %s\n