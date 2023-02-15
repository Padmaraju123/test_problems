package main

import (
	"bufio"
	"fmt"
	"os"
)

func checking(c rune) bool {
	mm := map[rune]bool{
		'a': true,
		'A': true,
		'e': true,
		'E': true,
		'i': true,
		'I': true,
		'o': true,
		'O': true,
		'u': true,
		'U': true,
	}
	if mm[c] == true {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ch, _, _ := reader.ReadRune()
	fmt.Println(checking(ch))
}
