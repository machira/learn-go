package iteration

import "strings"

func Repeat(character string, times int) string {
	if times <= 0 {
		return character
	}
	return strings.Repeat(character, times)
}
