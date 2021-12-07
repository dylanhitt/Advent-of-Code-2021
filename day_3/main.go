package main

import (
	"fmt"
	"strconv"

	common "github.com/dylanhitt/Advent-of-Code-2021/common"
)

func main() {
	lines, _ := common.ReadLines("input.txt")
	half := len(lines) / 2
	counts := map[int]int{}
	for i := range lines[0] {
		counts[i] = 0
	}

	for _, line := range lines {
		for i, bit := range line {
			if string(bit) == "1" {
				counts[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""
	for i := range lines[0] {
		if counts[i] > half {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
			continue
		}
		gamma = gamma + "0"
		epsilon = epsilon + "1"
	}

	gammaNu, _ := strconv.ParseInt(gamma, 2, 64)
	epNum, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammaNu * epNum)

	oxy := findOne(lines, 0, max)
	co := findOne(lines, 0, min)

	oxyNum, _ := strconv.ParseInt(oxy, 2, 64)
	fmt.Println(oxyNum)
	coNum, _ := strconv.ParseInt(co, 2, 64)
	fmt.Println(coNum)
	fmt.Println(oxyNum * coNum)
}

func findOne(strings []string, position int, f func(a int, b int) rune) string {
	if len(strings) <= 1 {
		return strings[0]
	}

	oneCount := 0
	zeroCount := 0

	for _, s := range strings {
		bit := s[position]
		if string(bit) == "1" {
			oneCount++
		}
		if string(bit) == "0" {
			zeroCount++
		}
	}

	filter := f(oneCount, zeroCount)
	matched := []string{}
	for _, s := range strings {
		if s[position] == byte(filter) {
			matched = append(matched, s)
		}
	}

	// if len(matched) == 0 {
	// 	matched = strings
	// }
	fmt.Println(matched)

	position++
	return findOne(matched, position, f)
}

func max(a, b int) rune {
	if a >= b {
		return '1'
	}
	return '0'
}

func min(a, b int) rune {
	if a < b {
		return '1'
	}
	return '0'
}
