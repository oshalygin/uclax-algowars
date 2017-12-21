package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/oshalygin/uclax-algowars/solutions"

	"github.com/fatih/color"
)

func main() {

	startMessage := "The start of something great"

	color.Green(startMessage)
	color.White(strings.Repeat("=", len(startMessage)))

	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	input := make([]int, n)
	println(n)
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &input[i])
	}

	result := solutions.PlusMinus(input)
	fmt.Println(result)

}
