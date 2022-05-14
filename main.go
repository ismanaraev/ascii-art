package main

import (
	args "ascii-art/args"
	"ascii-art/chars"
	"fmt"
	"os"
	"strings"
)

func main() {
	var regmap = make(map[string][]int)
	var n int
	var Indexlist []args.Index
	if len(os.Args) < 2 {
		return
	}
	input := os.Args[1]
	if len(os.Args) > 2 {
		if !args.CheckArgs(os.Args[2:], regmap, &Indexlist) {
			fmt.Println("Invalid arguments")
			return
		}
	}
	input = strings.ReplaceAll(input, "\\n", "\n")
	if input == "" {
		return
	}
	charstring := chars.ReadCharFile("standard.txt")
	charmap := chars.CreateCharMap(charstring)
	if !chars.CheckString(input, charmap) {
		fmt.Println("invalid Chars")
		return
	}
	inputlines := strings.Split(input, "\n")
	if len(input) < len(inputlines) {
		inputlines = inputlines[1:]
	}
	for _, item := range inputlines {
		chars.PrintLine(item, charmap, regmap, Indexlist, n)
		n += len(item)
	}
}
