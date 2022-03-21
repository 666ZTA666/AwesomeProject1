package YALIB

import (
	"fmt"
	"os"
	"strings"
)

func IndaxRune(s, sep string, file *os.File) int {
	var i, idx, n int
	idx = -1
	for {
		idx = strings.Index(s[n:], sep)
		if idx < 0 && i == 0 {
			return idx
		} else if idx >= 0 {
			i++
			n += idx + 1
			fmt.Fprintln(file, n)
			fmt.Println(n)
		} else if idx < 0 && i > 0 {
			return i
		} else {
			fmt.Println("аварийный выход хуй пойми почему", idx, i)
			break
		}
	}
	return i
}
