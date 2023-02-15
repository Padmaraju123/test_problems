package main

import "fmt"


func main(){
	var word string
	fmt.Scanln(&word)

	fmt.Println(string(word[0]))
	fmt.Printf("%c\n",word[0])
	fmt.Println(word[:1])

	//Last character of the string
	le := len(word)
	fmt.Println(string(word[le-1]))
	fmt.Printf("%c\n",word[le-1])
	fmt.Println(word[le-1:])
}