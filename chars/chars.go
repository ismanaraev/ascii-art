package chars

import (
	"ascii-art/args"
	"ascii-art/colors"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//This type defines an ascii char and consists of fields Lines and Width, Lines are ascii-symbol lines from top to bottom, one line each and Width are its Width
type Char struct {
	Lines []string
	Width int
}

//This function checks font name specified by option and points it to the corresponding file
func CheckFontName(file string) (string, error) {
	switch file {
	case "standard":
		file = "ascii/standard.txt"
	case "shadow":
		file = "ascii/shadow.txt"
	case "thinkertoy":
		file = "ascii/thinkertoy.txt"
	case "doom":
		file = "ascii/doom.txt"
	default:
		return "", errors.New("Invalid font name")
	}
	return file, nil
}

//This function Reads the file containing ascii-chars and validates it by number of lines
func ReadCharFile(file string) (string, error) {
	chars, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer chars.Close()
	data := make([]byte, 100000)
	n, err := chars.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	res := string(data[:n])
	filelen := strings.Split(res, "\n")
	if len(filelen) != 856 {
		return "", errors.New("invalid file, file must have 855 lines")
	}
	return res, nil
}

//This function Creates a string to Char map using the contents of a ascii-art char file as a single continious string
func CreateCharMap(allchars string) map[string]Char {
	var ct int
	var tmp []string
	var newcharset [][]string
	charset := strings.Split(allchars, "\n")
	charmap := make(map[string]Char)
	charset = charset[1:]
	for _, item := range charset {
		ct++
		tmp = append(tmp, item)
		if ct == 8 {
			newcharset = append(newcharset, tmp)
			tmp = []string{}
		}
		if ct == 9 {
			tmp = []string{}
			ct = 0
		}
		continue
	}
	for i, char := range newcharset {
		charmap[string(rune(i+32))] = Char{Lines: char, Width: len(char[0])}
	}
	charmap["\n"] = Char{Lines: make([]string, 8), Width: 1}
	return charmap
}

//This function Check if string has only ascii characters
func CheckString(input string, charmap map[string]Char) bool {
	for _, item := range input {
		if _, ok := charmap[string(item)]; !ok {
			return false
		}
	}
	return true
}

//This function finds the width of the all the ascii-chars of a given string
func GetStrWidth(s string, charmap map[string]Char) int {
	var ct int
	for _, ch := range s {
		ct += charmap[string(ch)].Width
	}
	return ct
}

//This function Counts the number of words in string
func CountWords(s string) int {
	f := strings.Split(s, " ")
	return len(f)
}

//this function gets the width of a current terminal, required by SetAlignment
func GetTermWidth() int {
	tw, err := exec.Command("tput", "cols").Output()
	if err != nil {
		log.Fatal(err)
	}
	tw = tw[:len(tw)-1]
	terminalWidth, err := strconv.Atoi(string(tw))
	if err != nil {
		log.Fatal(err)
	}
	return terminalWidth
}

//This function handles --alignment flag. It does it by altering the input string,
//In the given string we replace space chars, which have width 6, with non-ascii chars added to the map with
//the calculated width with respect to the terminal width and free space available
func SetAlignment(align, item string, charmap map[string]Char) string {
	n := GetStrWidth(item, charmap)
	terminalwidth := GetTermWidth()
	switch align {
	case "left":
		//Here we get the free space available to us in the current terminal
		shift := terminalwidth - n
		//Then, we create the string slice with required width
		newchar := []string{}
		for i := 0; i != 8; i++ {
			newchar = append(newchar, strings.Repeat(" ", shift))
		}
		//append the slice to the non-ascii char in the charmap
		charmap["Л"] = Char{Lines: newchar, Width: shift}
		//and add it to the string
		item = item + "Л"
	case "right":
		//Same as in left
		shift := terminalwidth - n
		var newchar []string
		for i := 0; i != 8; i++ {
			newchar = append(newchar, strings.Repeat(" ", shift))
		}
		charmap["Р"] = Char{Lines: newchar, Width: shift}
		item = "Р" + item
	case "center":
		//As usual, we get the available free space, then divide it by 2 to get the size for shift
		shift := terminalwidth - n
		lshift := shift / 2
		rshift := shift - lshift
		//Then, create two string slices with required width
		var lchar []string
		var rchar []string
		for i := 0; i != 8; i++ {
			lchar = append(lchar, strings.Repeat(" ", lshift))
			rchar = append(rchar, strings.Repeat(" ", rshift))
		}
		//add the non-ascii chars to the charmap
		charmap["Л"] = Char{Lines: lchar, Width: lshift}
		charmap["Р"] = Char{Lines: rchar, Width: rshift}
		//add the non-ascii chars with required width to the string both on left and right
		item = "Л" + item + "Р"
	case "justify":
		//Here we Trim the spaces on left and right, so justify will work as intended
		item = strings.Trim(item, " ")
		//Count Freespace
		freespace := terminalwidth - n
		//Count words, we need number of words to find the width of each space between them
		words := CountWords(item)
		//Here we add 1 to the word, so we could avoid ZeroDivisionError
		if words == 1 {
			words = 2
		}
		//Here we divide the freespace to number of spaces between words, to count the new width for each space
		sh := freespace / (words - 1)
		//Because the standard ascii space width is 6, we add six to the newly counted width value
		sh += 6
		//Again, create string slice with needed width
		var suchar []string
		for i := 0; i != 8; i++ {
			suchar = append(suchar, strings.Repeat(" ", sh))
		}
		//Add the slice to the map
		charmap["Щ"] = Char{Lines: suchar, Width: sh}
		//And here we replace each space in string with the char that has required number of spaces
		item = strings.ReplaceAll(item, " ", "Щ")
	}
	return item
}

//This function checks if the char should be colored, if yes, then returns the slice of integers, containing rgb color
func CheckPrintColor(s rune, Indexlist []args.Index, regmap map[string][]int, n int) []int {
	if in, ok := args.CheckIndex(Indexlist, n); ok {
		dexcol := in.Color
		return dexcol
	} else if col, ok := regmap[""]; ok {
		return col
	} else if col, ok := regmap[string(s)]; ok {
		return col
	}
	return []int{}
}

//This function iterates over the slice of string in Char and prints them
func PrintLine(item string, charmap map[string]Char, regmap map[string][]int, Indexlist []args.Index, n int) {
	if item == "" {
		fmt.Print("\n")
		return
	}
	//Because each character has height equal to 8
	for i := 0; i < 8; i++ {
		for _, letter := range item {
			//Check if we should color the string
			if color := CheckPrintColor(letter, Indexlist, regmap, n); len(color) != 0 {
				fmt.Print(colors.GetANSIColor(color) + charmap[string(letter)].Lines[i])
			} else {
				fmt.Print(charmap[string(letter)].Lines[i])
			}
			n++
		}
		//This n represents the true index of the char, given by main, here we reset it back to the value we had at the beginning of this line, because we print it line by line
		n = n - len(item)
		fmt.Print("\n")
	}
}

//Here we check if input overflows the terminal and won't fir and split the string, so all the characters that won't fit will be places on next line
func ValidateInput(inputlines []string, charmap map[string]Char) []string {
	tw := GetTermWidth()
	var wordlen int
	var word string
	var newinputlines []string
	for _, item := range inputlines {
		for _, ch := range item {
			if l := wordlen + charmap[string(ch)].Width; l <= tw {
				wordlen += charmap[string(ch)].Width
				word += string(ch)
			} else {
				newinputlines = append(newinputlines, word)
				wordlen = charmap[string(ch)].Width
				word = string(ch)
			}
		}
		wordlen = 0
		if word != "" {
			newinputlines = append(newinputlines, word)
			word = ""
		}
	}
	return newinputlines
}

//This function uses some dirty hack, i.e. it executes the command echo -e to interpret the backslash characters in the input
func FormatString(s string) string {
	formatted_str, err := exec.Command("echo", "-e", s).Output()
	if err != nil {
		log.Fatal(err)
	}
	s = string(formatted_str)
	s = s[:len(s)-1]
	return s
}
