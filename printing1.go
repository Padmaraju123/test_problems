package main 

import "fmt"

func main(){
	var n int 
	fmt.Scanln(&n)
	for n>0{
		fmt.Println(n)
		n--
	}
}