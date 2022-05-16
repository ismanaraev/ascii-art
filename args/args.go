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
var colorhsl = regexp.MustCompile(`^--color=hsl\((\d{1,3})°{0,1},{0,1} {0,1}(\d{1,3})%{0,1},{0,1} {0,1}(\d{1,3})%{0,1}\)$`)
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

func Help() {
	fmt.Print(`Usage: go run . [STRING] [OPTION]

EX: go run . something --color=<color>
`)
	os.Exit(0)
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

func SetTarget(args []string, color []int, regmap map[string][]int, Indexlist *[]Index) int {
	var ct int
	if len(args) == 1 {
		regmap[""] = color
		return 0
	}
	if strings.HasPrefix(args[1], "--color=") {
		return ct
	}
	if regdex.MatchString(args[1]) {
		*Indexlist = append(*Indexlist, SetIndex(args[1], color)...)
		return ct + 1
	}
	for _, ch := range args[1] {
		regmap[string(rune(ch))] = color
	}
	ct++
	return ct
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

func CheckNumbers(nums []int, mode string) bool {
	var checkval int
	if mode == "rgb" {
		checkval = 255
	}
	if mode == "hsl" {
		if nums[0] > 360 {
			return false
		}
		nums = nums[1:]
		checkval = 100
	}
	for _, item := range nums {
		if item > checkval || item < 0 {
			return false
		}
	}
	return true
}

func CheckArgs(args []string, regmap map[string][]int, Indexlist *[]Index) {
	var color []string
	var nums []string
	var numbers []int
	for i := 0; i < len(args); i++ {
		switch {
		case colorword.MatchString(args[i]):
			color = colorword.FindStringSubmatch(args[i])
			colorname := strings.ToLower(color[1])
			numbers, ok := colors.Colormap[colorname]
			if !ok {
				fmt.Println("Invalid color. Available colors are: red, green, blue, black, white, cyan, gray, purple, orange, pink, yellow, lime, teal ")
				os.Exit(0)
			}
			n := SetTarget(args[i:], numbers, regmap, Indexlist)
			i += n
		case colornum.MatchString(args[i]):
			nums = colornum.FindStringSubmatch(args[i])
			numbers = GetNumbers(nums)
			if !CheckNumbers(numbers, "rgb") {
				Help()
			}
			n := SetTarget(args[i:], numbers, regmap, Indexlist)
			i += n
		case colorhsl.MatchString(args[i]):
			nums = colorhsl.FindStringSubmatch(args[i])
			numbers = GetNumbers(nums)
			if !CheckNumbers(numbers, "hsl") {
				Help()
			}
			numbers[0], numbers[1], numbers[2] = colors.HslToRGB(numbers[0], numbers[1], numbers[2])
			n := SetTarget(args[i:], numbers, regmap, Indexlist)
			i += n
		case colorhex.MatchString(args[i]):
			nums = colorhex.FindStringSubmatch(args[i])
			numbers = colors.HextoRGB(nums[1])
			n := SetTarget(args[i:], numbers, regmap, Indexlist)
			i += n
		default:
			Help()
		}
	}
	if !ValidateIndexList(*Indexlist) {
		Help()
	}
}
