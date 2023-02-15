package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	fmt.Println(sli_b)
	sent := string(sli_b)
	fmt.Printf("%T %v\n", sent, sent)

}
