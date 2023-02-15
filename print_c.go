package main 

import "fmt"

func printing(word string){
	for _,v := range word{
		fmt.Println(string(v))
	}
}

func main(){
	var word string 
	fmt.Scanln(&word)
	printing(word)
}