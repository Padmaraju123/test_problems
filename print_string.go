package main

import "fmt"

func main(){
	var word string 
	fmt.Scanln(&word)
	for i,v := range word{
		fmt.Println(i,string(v))
	}
}