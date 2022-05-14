package args

import (
	colors "ascii-art/colors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var colorword = regexp.MustCompile(`^--color=([A-Za-z]+)$`)
var colornum = regexp.MustCompile(`^--color=rgb\((\d{1,3})[,;] {0,1}(\d{1,3})[,;] {0,1}(\d{1,3})\)$`)
var colorhsl = regexp.MustCompile(`^--color=hsl\((\d{1,3})°{0,1} {0,1}(\d{1,3})%{0,1} {0,1}(\d{1,3})%{0,1}\)$`)
var colorhex = regexp.MustCompile(`^--color=#([0-9a-fA-F]{6})$`)
var regdex = regexp.MustCompile(`^\[(\d*?)(:){0,1}(\d*?)\]$`)
var n int

type Index struct {
	Start int
	End   int
	Color []int
}

func (I Index) Validate() bool {
	if I.Start < 0 || I.Start > len(os.Args[1]) {
		return false
	}
	if I.End < 0 || I.End > len(os.Args[1]) {
		return false
	}
	return true
}

func (I Index) Match(s int) bool {
	if I.Start <= s && I.End > s {
		return true
	}
	return false
}

func ValidateIndexList(dexes []Index) bool {
	for _, item := range dexes {
		if !item.Validate() {
			return false
		}
	}
	return true
}

func CheckIndex(dexes []Index, pos int) (Index, bool) {
	for _, dex := range dexes {
		if dex.Match(pos) {
			return dex, true
		}
	}
	return Index{Start: -1}, false
}

func SetIndex(s string, rgb []int) []Index {
	var Indexlist []Index
	l := regdex.FindStringSubmatch(s)
	index := Index{}
	var delim bool
	if l[2] != "" {
		delim = true
	}
	var err error
	index.Start, err = strconv.Atoi(l[1])
	if err != nil {
		if l[1] == "" {
			index.Start = 0
		} else {
			log.Fatal(err)
		}
	}
	index.End, err = strconv.Atoi(l[3])
	if err != nil {
		if l[3] == "" {
			index.End = len(os.Args[1])
		} else {
			log.Fatal(err)
		}
	}
	if !delim {
		index.Start = index.End
		index.End++
	}
	var dexcol = make([]int, 3)
	for i := 0; i < 3; i++ {
		dexcol[i] = rgb[i]
	}
	index.Color = dexcol
	Indexlist = append(Indexlist, index)
	return Indexlist
}

func SetTarget(args []string, color []int, regmap map[string][]int, Indexlist *[]Index) {
	var target string
	if len(args) > 1 {
		if regdex.MatchString(args[1]) {
			*Indexlist = SetIndex(args[1], color)
			return
		}
		target = args[1]
		if strings.HasPrefix(args[1], "--color") {
			target = ""
		}
	} else {
		target = ""
	}
	if len(target) > 1 {
		for _, item := range target {
			regmap[string(item)] = color
		}
		return
	}
	regmap[target] = color
}

func GetNumbers(nums []string) []int {
	var numbers []int
	for _, item := range nums[1:] {
		n, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func CheckArgs(args []string, regmap map[string][]int, Indexlist *[]Index) bool {
	var color []string
	var nums []string
	var numbers []int
	var good bool
	for i, arg := range args {
		switch {
		case colorword.MatchString(arg):
			good = true
			color = colorword.FindStringSubmatch(arg)
			numbers, ok := colors.Colormap[color[1]]
			if !ok {
				return false
			}
			SetTarget(args[i:], numbers, regmap, Indexlist)
		case colornum.MatchString(arg):
			good = true
			nums = colornum.FindStringSubmatch(arg)
			numbers = GetNumbers(nums)
			SetTarget(args[i:], numbers, regmap, Indexlist)
		case colorhsl.MatchString(arg):
			good = true
			nums = colorhsl.FindStringSubmatch(arg)
			numbers = GetNumbers(nums)
			numbers[0], numbers[1], numbers[2] = colors.HslToRGB(numbers[0], numbers[1], numbers[2])
			SetTarget(args[i:], numbers, regmap, Indexlist)
		case colorhex.MatchString(arg):
			good = true
			nums = colorhex.FindStringSubmatch(arg)
			numbers = colors.HextoRGB(nums[1])
			SetTarget(args[i:], numbers, regmap, Indexlist)
		}
	}
	if !ValidateIndexList(*Indexlist) {
		return false
	}
	fmt.Println(len(regmap), len(*Indexlist))
	return good
}
