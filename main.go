package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"strings"
)

func getCommand(line string) string {
	// take the first word of the line
	if strings.Contains(line, " ") {
		return strings.Split(line, " ")[0]
	} else {
		return line
	}
}

func parseLine(line string) string {
	switch getCommand(line) {
	case "bc":
		line = line[3:]
		number, err := basicMath(line)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return fmt.Sprintf("%f", number)
	case "p2r": // polar to rectangular
		return p2r(line)
	case "r2p": // rectangular to polar
		return r2p(line)
	case "exit":
		return "exit"
	default:
		return "unknown command: " + line
	}
}

func main() {
	println("welcome to mathtool")
	println("(c) husky/floppydiskette 2022")
	println("")
	println("type help for a list of commands")

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	for {
		// read input
		line, err := rl.Readline()
		if err != nil {
			panic(err)
		}

		// parse input
		println(parseLine(line))
	}
}
