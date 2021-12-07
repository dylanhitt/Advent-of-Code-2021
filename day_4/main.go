package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dylanhitt/Advent-of-Code-2021/common"
)

func main() {
	lines, _ := common.ReadLines("input.txt")
	lineList := common.LineList{Slice: lines}
	draws := strings.Split(lineList.Remove(0), ",")
	boards := createBoards(lineList.Slice)

	firstWinner := 0
	lastWinner := 0
	for _, draw := range draws {
		for i := 0; i < len(boards); i++ {
			n, _ := strconv.Atoi(draw)
			if v, ok := boards[i].lookup[draw]; ok {
				boards[i].mark(v)
			}
			if boards[i].checkWin() {
				if firstWinner == 0 {
					firstWinner = boards[i].unmarkedSum() * n
				}
				lastWinner = boards[i].unmarkedSum() * n
				boards = append(boards[:i], boards[i+1:]...)
				i--
			}
		}
	}
	fmt.Println(firstWinner)
	fmt.Println(lastWinner)
}

func createBoards(lines []string) []board {
	boards := []board{}
	b := newBoard()
	y := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		nums := strings.Fields(line)
		for x, num := range nums {
			c := coords{x: x, y: y}
			b.add(c, num)
		}

		y++
		if b.isFull() {
			boards = append(boards, b)
			y = 0
			b = newBoard()
		}
	}
	return boards
}

type coords struct {
	x int
	y int
}

type field struct {
	marked bool
	val    int
	hasVal bool
}

type board struct {
	fields [5][5]field
	lookup map[string]coords
}

func newBoard() board {
	a := [5][5]field{}
	return board{
		fields: a,
		lookup: map[string]coords{},
	}
}

func (b *board) checkWin() bool {
	for i := range b.fields {
		if b.checkRow(i) {
			return true
		}
		if b.checkCol(i) {
			return true
		}
	}
	return false
}

func (b *board) checkRow(y int) bool {
	for x := range b.fields {
		if !b.fields[y][x].marked {
			return false
		}
	}
	return true
}

func (b *board) checkCol(x int) bool {
	for y := range b.fields {
		if !b.fields[y][x].marked {
			return false
		}
	}
	return true
}

func (b *board) add(c coords, val string) {
	n, _ := strconv.Atoi(val)
	b.fields[c.y][c.x] = field{
		marked: false,
		val:    n,
		hasVal: true,
	}
	b.lookup[val] = c
}

func (b *board) mark(c coords) {
	b.fields[c.y][c.x].marked = true
}

func (b *board) isFull() bool {
	return b.fields[4][4].hasVal
}

func (b *board) unmarkedSum() int {
	sum := 0
	for _, row := range b.fields {
		for _, val := range row {
			if !val.marked {
				sum = sum + val.val
			}
		}
	}
	return sum
}
