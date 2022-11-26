package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	lyrics := []string{}
	for i := startBottles; i > (startBottles - takeDown); i-- {
		l := fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(intToStr(i)), genBottle(i))
		lyrics = append(lyrics, l)
		lyrics = append(lyrics, l)
		lyrics = append(lyrics, "And if one green bottle should accidentally fall,")

		next := i - 1
		lyrics = append(lyrics, fmt.Sprintf("There'll be %s green %s hanging on the wall.", intToStr(next), genBottle(next)))

		if takeDown > 1 && next+takeDown != startBottles {
			lyrics = append(lyrics, "")
		}
	}
	return lyrics
}

func genBottle(x int) string {
	if x == 1 {
		return "bottle"
	}
	return "bottles"
}

func intToStr(x int) string {
	switch x {
	case 0:
		return "no"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	default:
		return ""
	}
}
