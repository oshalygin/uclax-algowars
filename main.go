package main

import (
	"strings"

	"github.com/fatih/color"
)

func main() {
	startMessage := "The start of something great"

	color.Green(startMessage)
	color.Green(strings.Repeat("=", len(startMessage)))

}
