package common

import (
	"bufio"
	"os"
)

type LineList struct {
	Slice []string
}

// ReadLines reads lines from a file into an array
func ReadLines(path string) (LineList, error) {
	file, err := os.Open(path)
	if err != nil {
		return LineList{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return LineList{lines}, scanner.Err()
}

func (c *LineList) Remove(i int) string {
	removed := c.Slice[i]
	c.Slice = append(c.Slice[:i], c.Slice[i+1:]...)
	return removed
}
