package main

import "fmt"

var index int
func remove_elements(sli1,sli2 []int){
	for _,v := range sli2{
		for j,w := range sli1{
			if v==w{
				index =j
				break
			}
		}
		sli1 = append(sli1[:index],sli1[index+1:]...)
		fmt.Println(sli1)
	}
}

func main(){
	first_s := []int{10,20,30,40,50,60,70,80,90,100}

	var size int 
	fmt.Scanln(&size)

	int_S := []int{}

	for size>0{
		var nn int 
		fmt.Scanln(&nn)
		int_S = append(int_S, nn)
		size--
	}

	remove_elements(first_s,int_S)
	

}