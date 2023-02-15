package main

import "fmt"

var count int
func string_rotation(s1,s2 string)int{

	le := len(s1)
	for i:=0;i<le;i++{
		out := s1[le-(i+1):]+s1[:le-(i+1)]
		count++
		if out == s2{
			break
		}
	}
	return count
}

func main(){
	var str1,str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	fmt.Println(string_rotation(str1,str2))
}