package args

import (
	"log"
	"os"
	"strconv"
)

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
	if I.End < I.Start {
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
