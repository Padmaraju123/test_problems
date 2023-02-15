package main

import (
	"fmt"
	"sort"
)

var dup,ms int

func delete_dup(sli []int){

	sort.Ints(sli)

	le := len(sli)-1
	for i,_:= range sli[:le]{
		if sli[i] == sli[i+1]{
			dup = sli[i]
			break
		}
	}

	for _,v := range sli{
		for i:=v ; i<v+len(sli);i++{
			if v ==i {
				continue
			}else{
				ms = i+1
				break
			}
		}
	}

	fmt.Println(dup,ms)
	

	
}

func main() {
	var size int
	fmt.Scanln(&size)

	sli := []int{}

	for i := 0; i < size; i++ {
		var vv int
		fmt.Scanln(&vv)
		sli = append(sli, vv)
	}

	delete_dup(sli)
}
