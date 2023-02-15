package main

import (
	"bufio"
	"fmt"
	"os"
)

func (b byte)bool{
	if (b>=97 && b<=122) 
}



func checking1(voch byte)bool{
	if voch == 'a' || voch == 'e' || voch == 'i' || voch == 'o' || voch == 'u' ||
		voch == 'A' || voch == 'E' || voch == 'I' || voch == 'O' || voch == 'U'{
			return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	byt, _ := reader.ReadByte()
	fmt.Println(byt)
	fmt.Println(checking1(byt))
	fmt.Println(checking2(byt))
}
