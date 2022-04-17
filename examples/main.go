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

var