package main

import "fmt"

func main(){
	var n int 
	fmt.Scanln(&n)
	for n>0{
		var i int
		fmt.Scanln(&i)
		fmt.Println(i)
		n--
	}
}