package main

import (
	"fmt"
	"strconv"
	"strings"

	common "github.com/dylanhitt/Advent-of-Code-2021/common"
)

func main() {
	lines, _ := common.ReadLines("input.txt")

	depth := 0
	horizontal := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		fmt.Println(words)
		num, _ := strconv.Atoi(words[1])
		switch words[0] {
		case "forward":
			horizontal = horizontal + num
		case "down":
			depth = depth + num
		case "up":
			depth = depth - num
		}
	}

	fmt.Println(depth * horizontal)
}
