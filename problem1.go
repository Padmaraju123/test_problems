package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func conversion(sent string)string{
	out := ""
	for _,v := range sent{
		if unicode.IsLower(v)==true{
			k:=unicode.ToUpper(v)
			out+=string(k)
		}else if string(v)==" "{
			out += " "
		}else{
			g := unicode.ToLower(v)
			out+=string(g)
		}
	}
	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	
	fmt.Println(conversion(sent))
}
