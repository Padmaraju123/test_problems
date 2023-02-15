package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	word1, _ := reader.ReadString('\n')
	//here the word has end with \n because we put the delimit for that
	//to remove that we use trim
	
}
