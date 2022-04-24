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
		println("usage: examples 