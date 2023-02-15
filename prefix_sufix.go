package main

import "fmt"

func prefix_sufix(w1,w2 string)bool{
	le := len(w1)
	ch := string(w1[le-1])
	var ls int
	for i,v := range w2{
		if string(v)==ch{
			ls = i
		}
	}
	ww1 := string(w2[:ls+1])
	ww2 := string(w1[le-(ls+1):])
	if ww1==ww2{
		return true
	}
	return false

	
}

func main(){
	var w1 ,w2 string 
	fmt.Scan(&w1,&w2)
	fmt.Println(prefix_sufix(w1,w2))
}