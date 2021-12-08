package main

import (
	"fmt"
	"strconv"

	common "github.com/dylanhitt/Advent-of-Code-2021/common"
)

func main() {
	lines, _ := common.ReadLines("input.txt")
	prev := 0
	count := 0
	for _, line := range lines.Slice {
		num, _ := strconv.Atoi(line)
		if num > prev {
			count++
		}
		prev = num
	}
	fmt.Println(count)
}
