package main

import "fmt"

func main() {
	const size = 26
	const length_ = (size + size - 1)
	mid := length_/2 + 1
	i := 0
	j := 0

	letters := ""

	for true {
		foo := (i - j)
		currentChar := toCharStr(foo)
		i++
		if i >= mid {
			j += 2
		}
		// fmt.Print(string(currentChar))
		letters += string(currentChar)
		if i >= length_ {
			break
		}
	}

	/*
		2 lazy 3 optimze
	*/

	arr := [length_][length_]string{}

	for y := 0; y < length_; y++ {
		x := 0
		position := 0
		for true {
			arr[y][x] = toCharStr(position)
			if x < y {
				position++
			}
			if x > length_-y-2 {
				position--
			}
			x++
			if x >= length_ {
				break
			}
		}
	}
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr[y]); x++ {
			fmt.Print(arr[y][x])
		}
		fmt.Println("")
	}
}

func toCharStr(i int) string {
	return string('A' + i)
}
