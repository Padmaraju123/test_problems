package main

import (
	"fmt"
	"unicode"
)

func shift_numbers(word string)string{
	out1 := ""
	out2 := ""
	for _,v := range word{
		if unicode.IsDigit(v)==true{
			out2+=string(v)
		}else{
			out1+=string(v)
		}
	}
	return out1+out2

}

func main(){
	var word string 
	fmt.Scanln(&word)
	fmt.Println(shift_numbers(word))
}