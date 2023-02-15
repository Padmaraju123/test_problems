package main

import (
	"fmt"
	"math"
)

func first_square(m, n int){
	for i := m; i <= n; i++ {
		vv := math.Sqrt(float64(i))
		
	}
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	first_square(m, n)
}	
