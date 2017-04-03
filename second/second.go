package main

import (
	"fmt"
	"strings"
)

func main() {
	word := strings.Split("SAIPPUAKAUPPIAS", "")
	const length = 47

	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			letter := word[(x*47+y)%15]
			fmt.Print(letter)
		}
		fmt.Println("")
	}
}
