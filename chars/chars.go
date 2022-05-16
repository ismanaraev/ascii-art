package chars

import (
	"ascii-art/args"
	"ascii-art/colors"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"
)

var Standardhash = []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}

func ReadCharFile(file string) string {
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
	hash := sha256.New()
	hash.Write(data[:n])
	if string(hash.Sum(nil)) != string(Standardhash) {
		fmt.Println("standard.txt file is corrupted")
		os.Exit(0)
	}
	return string(data[:n])
}

func CreateCharMap(allchars string) map[string][]string {
	var ct int
	var tmp []string
	var newcharset [][]string
	charset := strings.Split(allchars, "\n")
	charmap := make(map[string][]string)
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
		charmap[string(rune(i+32))] = char
	}
	charmap["\n"] = make([]string, 8)
	return charmap
}

func CheckString(input string, charmap map[string][]string) bool {
	for _, item := range input {
		if _, ok := charmap[string(item)]; !ok {
			return false
		}
	}
	return true
}

func PrintLine(item string, charmap map[string][]string, regmap map[string][]int, Indexlist []args.Index, n int) {
	if item == "" {
		fmt.Print("\n")
		return
	}
	for i := 0; i < 8; i++ {
		for _, letter := range item {
			if in, ok := args.CheckIndex(Indexlist, n); ok {
				dexcol := in.Color
				fmt.Print(colors.GetANSIColor(dexcol[0], dexcol[1], dexcol[2]) + charmap[string(letter)][i])
			} else if col, ok := regmap[""]; ok {
				if len(regmap) != 1 {
					args.Help()
				}
				fmt.Print(colors.GetANSIColor(col[0], col[1], col[2]) + charmap[string(letter)][i])
			} else if col, ok := regmap[string(letter)]; ok {
				fmt.Print(colors.GetANSIColor(col[0], col[1], col[2]) + charmap[string(letter)][i])
			} else {
				fmt.Print("\033[0m" + charmap[string(letter)][i])
			}
			n++
		}
		n = n - len(item)
		fmt.Print("\033[0m\n")
	}
}
