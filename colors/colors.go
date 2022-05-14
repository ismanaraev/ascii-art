package colors

import (
	"log"
	"math"
	"strconv"
)

func HextoRGB(s string) []int {
	var res []int
	for i := 0; i < 5; i = i + 2 {
		r, err := strconv.ParseInt(s[i:i+2], 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, int(r))
	}
	return res
}

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
	case H1 >= 0.0 && H1 <= 1.0:
		R, G, B = C, X, 0.0
	case H1 >= 1.0 && H1 <= 2.0:
		R, G, B = X, C, 0.0
	case H1 >= 2.0 && H1 <= 3.0:
		R, G, B = 0.0, C, X
	case H1 >= 3.0 && H1 <= 4.0:
		R, G, B = 0.0, X, C
	case H1 >= 4.0 && H1 <= 5.0:
		R, G, B = X, 0.0, C
	case H1 >= 5.0 && H1 <= 6.0:
		R, G, B = C, 0.0, X
	}
	R = math.Round((R + m) * 255)
	G = math.Round((G + m) * 255)
	B = math.Round((B + m) * 255)
	return int(R), int(G), int(B)
}

func GetANSIColor(r, g, b int) string {
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}
