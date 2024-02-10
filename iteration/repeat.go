package iteration

import "strings"

func Repeat(s string) string {
	return strings.Repeat(s, 5)
}

func Repeat2(c string) string {
	s := ""
	for i := 0; i < 5; i++ {
		s = s + c
	}
	return s
}
