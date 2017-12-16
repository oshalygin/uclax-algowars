package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/oshalygin/uclax-algowars/solutions"
)

func main() {

	startMessage := "The start of something great"

	color.Green(startMessage)
	color.White(strings.Repeat("=", len(startMessage)))

	io := bufio.NewReader(os.Stdin)

	bob := make([]int, 3)
	alice := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &alice[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &bob[i])
	}

	result := solutions.CompareTheTriplets(alice, bob)
	fmt.Printf("%d %d", result[0], result[1])
}
