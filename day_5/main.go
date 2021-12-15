package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/dylanhitt/Advent-of-Code-2021/common"
)

type coordinate struct {
	x int
	y int
}

type vent struct {
	start  coordinate
	end    coordinate
	slope  int
	maxX   int
	maxY   int
	minX   int
	minY   int
	isVert bool
	isHort bool
	isDiag bool
}

type vents struct {
	vents  []vent
	boundX int
	boundY int
}

type point struct {
	count int
}

func main() {
	lineList, _ := common.ReadLines("input.txt")
	vf := genField(lineList.Slice)

	field := make([][]point, vf.boundY+1)
	for y := 0; y < len(field); y++ {
		field[y] = make([]point, vf.boundX+1)
	}

	for _, v := range vf.vents {
		if v.isHort {
			for x := v.minX; x <= v.maxX; x++ {
				field[v.maxY][x].count++
			}
			continue
		}
		if v.isVert {
			for y := v.minY; y <= v.maxY; y++ {
				field[y][v.maxX].count++
			}
			continue
		}
		if v.isDiag {
			distance := v.end.x - v.start.x
			for x := 0; x <= distance; x++ {
				field[v.start.y+(x*v.slope)][x+v.start.x].count++
			}
		}
	}

	moreThanTwo := 0
	for _, row := range field {
		for _, v := range row {
			if v.count > 1 {
				moreThanTwo++
			}
		}
	}

	fmt.Println(moreThanTwo)
}

func genField(lines []string) vents {
	ventArray := []vent{}

	maxX := 0
	maxY := 0
	for _, rc := range lines {
		v := createVent(strings.Split(rc, " -> "))
		maxX = int(math.Max(float64(maxX), float64(v.maxX)))
		maxY = int(math.Max(float64(maxY), float64(v.maxY)))
		ventArray = append(ventArray, v)
	}

	return vents{
		vents:  ventArray,
		boundX: maxX,
		boundY: maxY,
	}
}

func createVent(rawCoords []string) vent {
	if len(rawCoords) > 2 {
		fmt.Println(rawCoords)
		fmt.Println("Invalid coordinate set")
		os.Exit(1)
	}
	return newVent(createCoordinate(rawCoords[0]), createCoordinate(rawCoords[1]))
}

func newVent(one coordinate, two coordinate) vent {
	maxX := math.Max(float64(one.x), float64(two.x))
	maxY := math.Max(float64(one.y), float64(two.y))
	minX := math.Min(float64(one.x), float64(two.x))
	minY := math.Min(float64(one.y), float64(two.y))

	start := one
	end := two
	if one.x >= two.x {
		start = two
		end = one
	}

	slope := 0

	if (start.x - end.x) != 0 {
		slope = (start.y - end.y) / (start.x - end.x)
	}

	return vent{
		start:  start,
		end:    end,
		slope:  slope,
		maxX:   int(maxX),
		maxY:   int(maxY),
		minX:   int(minX),
		minY:   int(minY),
		isVert: start.x == end.x,
		isHort: start.y == end.y,
		// rise over run has to be one
		isDiag: math.Abs(float64(slope)) == 1,
	}
}

func createCoordinate(rawCoord string) coordinate {
	raw := strings.Split(rawCoord, ",")
	if len(raw) > 2 {

		fmt.Println("Invalid coordinate")
		os.Exit(1)
	}
	x, err := strconv.Atoi(raw[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	y, err := strconv.Atoi(raw[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return coordinate{x: x, y: y}
}
