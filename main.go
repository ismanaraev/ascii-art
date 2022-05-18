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
	var outfile *os.File
	if len(os.Args) < 3 {
		args.Help()
		return
	}
	input := os.Args[1]
	if len(os.Args) > 3 {
		args.CheckArgs(os.Args[3:], regmap, &Indexlist, &outfile)
	}
	input = strings.ReplaceAll(input, "\\n", "\n")
	if input == "" {
		return
	}
	charstring := chars.ReadCharFile(os.Args[2])
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
		if outfile != nil {
			chars.WriteToFile(outfile, item, charmap, regmap, Indexlist)
		} else {
			chars.PrintLine(item, charmap, regmap, Indexlist, n)
			n += len(item)
		}
	}
}
