package main

import "fmt"

func duplicate_val(sli []int)(int,int){
	mm := map[int]bool{}

	for i,v := range sli{
		if mm[v]==true{
			return i,v 
		}
		mm[v]=true
	}
	return -1,-1
}

func main(){
	var size int
	fmt.Scanln(&size)

	sli := make([]int,size)
	for i:=0;i<size;i++{
		fmt.Scan(&sli[i])
	}
	
	pos,val := duplicate_val(sli)
	fmt.Println(pos,val)
}