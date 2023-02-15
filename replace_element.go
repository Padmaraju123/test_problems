package main

import "fmt"

func replace_n(sli [][]int,n int){
	for _,v := range sli{
		le := len(v)-1
		v[le]=n
	}
	fmt.Println(sli)
}

func main(){
	sli := [][]int{
		{2,3},
		{5,6,5},
		{8,4},
	}
	
	var n int 
	fmt.Scanln(&n)

	replace_n(sli,n)

}