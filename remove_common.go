package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func common_elements(sli1,sli2 []int)[]int{
	mm := map[int]bool{}

	for _,v := range sli1{
		if mm[v]==true{
			continue
		}
		mm[v]=true
	}
	
	out := []int{}

	for _ , b := range sli2{
		if mm[b]==true{
			out = append(out, b)
		}
	}
	return out
}

func main() {
	reader1 := bufio.NewReader(os.Stdin)
	sli1, _, _ := reader1.ReadLine()

	reader2 := bufio.NewReader(os.Stdin)
	sli2, _, _ := reader2.ReadLine()

	sent1 := string(sli1)
	sent2 := string(sli2)

	split_1 := strings.Split(sent1, " ")

	split_2 := strings.Split(sent2, " ")

	int_1 := []int{}
	for _, v := range split_1 {
		j, _ := strconv.Atoi(v)
		int_1 = append(int_1, j)
	}

	int_2 := []int{}
	for _, v := range split_2 {
		j, _ := strconv.Atoi(v)
		int_2 = append(int_2, j)
	}
	
	fmt.Println(common_elements(int_1,int_2))


}
