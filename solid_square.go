package main

import (
	"fmt"
	"strings"
)

func solid_square(m int) {
	i := 1
	for i <= m {
		fmt.Println(strings.Repeat("* ", i))
		i++
	}
}

func main() {
	var m int
	fmt.Scan(&m)
	solid_square(m)

}
