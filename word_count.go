package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func words_count(sli []string) {

	mm := map[string]bool{}
	ss := []string{}
	for _, vv := range sli {
		if mm[vv] == true {
			continue
		}
		ss = append(ss, vv)
		mm[vv] = true
	}

	sort.Strings(ss)

	for _, ww := range ss {
		count := 0
		for _, kk := range sli {
			if ww == kk {
				count++
			}
		}
		fmt.Printf("%s %d\n", ww, count)
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	sli_s := strings.Split(sent, " ")
	words_count(sli_s)
}
