package main

import "fmt"

var ss int = 0

func sum_of_digits(num int)int{

	for num>0{
		re := num%10
		ss+=re
		num = num/10
	}
	return ss
}

func main(){
	var num int
	fmt.Scanln(&num)
	fmt.Println(sum_of_digits(num))
}