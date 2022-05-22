package colors

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

//This function recieves hex string in #ff00ff format and parses it to []int
func HextoRGB(s string) []int {
	var res []int
	for i := 0; i < 5; i = i + 2 {
		r, err := strconv.ParseInt(s[i:i+2], 16, 64) //We don't have to worry about getting out of range, input is always valid because of regex
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, int(r))
	}
	return res
}

//This function converts hsl to RGB colors using mathmatical formula
func HslToRGB(H, S, L int) (int, int, int) {
	h := float64(H)
	s := float64(S) / 100
	l := float64(L) / 100
	C := (1.0 - math.Abs((2.0*l)-1.0)) * s
	H1 := h / 60.0
	X := C * (1.0 - math.Abs((math.Mod(H1, 2.0))-1.0))
	m := l - (C / 2.0)
	var R, G, B float64
	switch {
	case H1 >= 0.0 && H1 < 1.0:
		R, G, B = C, X, 0.0
	case H1 >= 1.0 && H1 < 2.0:
		R, G, B = X, C, 0.0
	case H1 >= 2.0 && H1 < 3.0:
		R, G, B = 0.0, C, X
	case H1 >= 3.0 && H1 < 4.0:
		R, G, B = 0.0, X, C
	case H1 >= 4.0 && H1 < 5.0:
		R, G, B = X, 0.0, C
	case H1 >= 5.0 && H1 < 6.0:
		R, G, B = C, 0.0, X
	}
	R = math.Round((R + m) * 255)
	G = math.Round((G + m) * 255)
	B = math.Round((B + m) * 255)
	return int(R), int(G), int(B)
}

//This function recieves rgb colors and processes them to string formatted as \033[38;2;r;g;bm
func GetANSIColor(color []int) string {
	r, g, b := color[0], color[1], color[2]
	if r == 500 {
		return "\033[8m"
	}
	if r == 501 {
		return "\033[5m"
	}
	fmt.Print("\033[0m")
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}
