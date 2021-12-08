package main

import (
	"fmt"
	"strconv"

	common "github.com/dylanhitt/Advent-of-Code-2021/common"
)

func main() {
	l, _ := common.ReadLines("input.txt")
	lines := l.Slice
	gamma := ""
	epsilon := ""
	for i := 0; i < len(lines[0]); i++ {
		bit := commonBit(lines, i)
		gamma = gamma + string(bit)
		epsilon = epsilon + flipBit(bit)
	}

	gammaNu, _ := strconv.ParseInt(gamma, 2, 64)
	epNum, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammaNu * epNum)

	oxy, _ := strconv.ParseInt(findOne(lines, 0, max), 2, 64)
	co, _ := strconv.ParseInt(findOne(lines, 0, min), 2, 64)
	fmt.Println(oxy * co)
}

func commonBit(strings []string, position int) rune {
	oneCount, zeroCount := getCount(strings, position)
	return max(oneCount, zeroCount)
}

func findOne(strings []string, position int, f func(a int, b int) rune) string {
	if len(strings) == 1 {
		return strings[0]
	}

	oneCount, zeroCount := getCount(strings, position)
	filter := f(oneCount, zeroCount)
	matched := []string{}
	for _, s := range strings {
		if s[position] == byte(filter) {
			matched = append(matched, s)
		}
	}

	position++
	return findOne(matched, position, f)
}

func getCount(strings []string, position int) (one int, zero int) {
	for _, s := range strings {
		bit := s[position]
		if string(bit) == "1" {
			one++
			continue
		}
		zero++
	}
	return
}

func flipBit(bit rune) string {
	if string(bit) == "1" {
		return "0"
	}
	return "1"
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
