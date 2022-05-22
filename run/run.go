package run

import (
	"ascii-art/args"
	"ascii-art/chars"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Run() error {
	var regmap = make(map[string][]int)
	var n int
	var Indexlist []args.Index
	var outfile *os.File
	var align string
	if len(os.Args) < 3 {
		args.Help()
		return errors.New("Not enough arguments")
	}
	input := os.Args[1]
	if len(os.Args) > 3 {
		err := args.CheckArgs(os.Args[3:], regmap, &Indexlist, &outfile, &align)
		if err != nil {
			args.Help()
			return err
		}
	}
	input = chars.FormatString(input)
	if input == "" {
		return errors.New("empty input string")
	}
	file, err := chars.CheckFontName(os.Args[2])
	if err != nil {
		return err
	}
	charstring, err := chars.ReadCharFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	charmap := chars.CreateCharMap(charstring)
	if !chars.CheckString(input, charmap) {
		fmt.Println("invalid Chars")
		return errors.New("invalid chars in the input")
	}
	inputlines := strings.Split(input, "\n")
	if len(input) < len(inputlines) {
		inputlines = inputlines[1:]
	}
	inputlines = chars.ValidateInput(inputlines, charmap)
	if outfile != nil {
		os.Stdout = outfile
	}
	for _, item := range inputlines {
		if align != "" {
			item = chars.SetAlignment(align, item, charmap)
		}
		chars.PrintLine(item, charmap, regmap, Indexlist, n)
		n += len(item)
	}
	return nil
}
