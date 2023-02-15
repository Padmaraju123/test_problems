package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var sum, count int

func avg_sum(c, ss int)int{
	return ss / c

}

func numbers_instring(sent string)int{
	for _, v := range sent {
		if unicode.IsNumber(v) == true {
			count++
			k, _ := strconv.Atoi(string(v))
			sum += k
		}
	}
	fmt.Println(sum)
	return avg_sum(count, sum)

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	fmt.Println(numbers_instring(sent))
}
