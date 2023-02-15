package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var a, b int

func checking(sli []int) {
	mm := map[int]bool{}
	for _, d := range sli {
		if mm[d] == true {
			a = d
			break
		} else {
			mm[d] = true
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	split_S := strings.Split(sent, " ")

	int_S := []int{}
	for _, v := range split_S {
		vv, _ := strconv.Atoi(v)
		int_S = append(int_S, vv)
	}
	checking(int_S)
}
