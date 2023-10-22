package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune) string {
	var sb strings.Builder // create a string builder
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		sb.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])) // append the formatted character to the string builder
	}
	sb.WriteString("\n") // append a newline character
	return sb.String()   // return the string value of the string builder
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gorainbow")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	say, err := cowsay.Say(
		string(output),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}

	r := []rune(say)
	fmt.Print(print(r))

}
