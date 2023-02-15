package main

import (
	"fmt"
	"sort"
)

func missingAndDuplicate(Data []int) {
	var final []int
	sort.Ints(Data)
	var Duplicate int
	for i := 0; i < len(Data)-1; i++ {
		if Data[i] == Data[i+1] {
			Duplicate = Data[i]
		}
	}
	final = append(final, Duplicate)

	var missing int
	min := Data[0]
	for i := min; i < min+len(Data); i++ {
		for _, val := range Data {
			if i == val {
				continue
			} else {
				missing = i
			}
		}
	}
	final = append(final, missing)
	fmt.Println(final)
}
func main() {
	input1 := []int{17, 19, 18, 18, 21}
	input2 := []int{28, 31, 28, 31, 29, 30}
	missingAndDuplicate(input1)
	missingAndDuplicate(input2)
}
