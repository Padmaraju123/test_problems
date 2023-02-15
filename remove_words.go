package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func remove_words(word_S []string,k int)string{
	out := ""
	for _,word := range word_S{
		le := len(word)
		if le == k{
			continue
		}
		out+=word+" "
	}
	return out
	
}

func main() {
	var k int
	fmt.Scanln(&k)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	word_S := strings.Split(sent, " ")
	fmt.Println(remove_words(word_S,k))
}
