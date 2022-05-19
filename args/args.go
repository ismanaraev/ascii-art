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

var (
	isarg        = regexp.MustCompile(`^--color=|^--output=`)
	colorword    = regexp.MustCompile(`^--color=([A-Za-z]+)$`)
	colornum     = regexp.MustCompile(`^--color=rgb\((\d{1,3})[,;] {0,1}(\d{1,3})[,;] {0,1}(\d{1,3})\)$`)
	colorhsl     = regexp.MustCompile(`^--color=hsl\((\d{1,3})°{0,1},{0,1} {0,1}(\d{1,3})%{0,1},{0,1} {0,1}(\d{1,3})%{0,1}\)$`)
	colorhex     = regexp.MustCompile(`^--color=#([0-9a-fA-F]{6})$`)
	regdex       = regexp.MustCompile(`^\[(\d*?)(:){0,1}(\d*?)\]$`)
	outfileregex = regexp.MustCompile(`^--output=(\w*?)$`)
)

//display help, someday I will write better help string
func Help() {
	fmt.Print(`Usage: go run . [STRING] [BANNER] [OPTION]

EX: go run . something standard --color=<color>
`)
	os.Exit(0)
}

//this function sets the string or index to color in case when string is given, returns the number of args read
func SetTarget(args []string, color []int, regmap map[string][]int, Indexlist *[]Index) int {
	var ct int
	if len(args) == 1 {
		regmap[""] = color
		return 0
	}
	if isarg.MatchString(args[1]) {
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

//this function takes the submatch of colornum and colorhsl and gets the numbers in rgb(255, 0, 0) or hsl(100, 20, 50) etc...
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

//rgb and hsl numbers validation
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

//This function iterates the argument list and forms Indexlist and regmap based on arguments entered
func CheckArgs(args []string, regmap map[string][]int, Indexlist *[]Index, outfile **os.File) {
	var color []string
	var nums []string
	var numbers []int
	for i := 0; i < len(args); i++ {
		//this switch checks for args by regex, to add new type of arg, just enter new regex
		switch {
		case colorword.MatchString(args[i]):
			color = colorword.FindStringSubmatch(args[i])
			colorname := strings.ToLower(color[1])
			numbers, ok := colors.Colormap[colorname]
			if !ok {
				fmt.Println("Invalid color. Available colors are: red, green, blue, black, white, cyan, gray, purple, orange, pink, yellow, lime, teal, transparent, blink ")
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
		case outfileregex.MatchString(args[i]):
			output := outfileregex.FindStringSubmatch(args[i])
			_, err := os.Create(output[1])
			if err != nil {
				log.Fatal(err)
			}
			*outfile, err = os.OpenFile(output[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

		default:
			Help()
		}
	}
	//check if index is valid, i.e. doesn't go out of range
	if !ValidateIndexList(*Indexlist) {
		fmt.Println("invalid index")
		os.Exit(0)
	}
}
