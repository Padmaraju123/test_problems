package main

import (
	"bufio"
	"fmt"
	"os"
)

func string_byteArry(word string) []byte {
	out := []byte{}
	for _, v := range word {
		vv := byte(v)
		out = append(out, vv)
	}
	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	
	final := string_byteArry(sent)
	fmt.Printf("%T %v\n", final, final)

	final1 := []byte(sent)
	fmt.Printf("%T %v\n",final1,final1)
}
