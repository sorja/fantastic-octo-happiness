package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arr [10]int
	break_ := false
	nums := 99998
	i := 0
	for true {
		x := strings.Split(strconv.Itoa(i), "")

		for _, v := range x {
			foo, _ := strconv.Atoi(v)
			arr[foo]++
			if arr[foo] >= nums {
				break_ = true
				break
			}
		}

		if break_ {
			break
		}

		strconv.Itoa(i)
		i++
	}

	fmt.Println(arr, i)
}
