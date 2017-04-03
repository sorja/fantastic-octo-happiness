package main

import (
	"fmt"
	"strings"
)

func main() {
	word := strings.Split("SAIPPUAKAUPPIAS", "")
	result := ""
	length := 47
	for i := 0; i < length*length; i++ {
		b := i % len(word)
		result += word[b]
		if i%46 == 0 && i != 0 {
			result += "\n"
		}
	}
	fmt.Println(result)
}
