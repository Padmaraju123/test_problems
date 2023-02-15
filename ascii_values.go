package main

import (
	"bufio"
	"fmt"
	"os"
)

func ascii_values(word string) {
	le := len(word)
	for _, v := range word[:le-2] {
		fmt.Println(v)
		fmt.Printf("%c\n",v)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	ascii_values(word)

}
