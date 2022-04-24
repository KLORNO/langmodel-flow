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
	examples = append(examp