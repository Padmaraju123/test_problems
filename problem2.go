package main

import "fmt"

func prime_numbers(n1,n2 int){
	
	for v:=n1;v<=n2;v++{
		if v == 1 || v == 2{
			continue
		}else{
			status := true
			for i:=2;i<v;i++{
				if v%i==0{
					status = false
					break
				}
			}
			if status{
				fmt.Println(v)
			}
		}
		
	}
}

func main(){
	var num1,num2 int
	fmt.Scan(&num1,&num2)
	prime_numbers(num1,num2)
}