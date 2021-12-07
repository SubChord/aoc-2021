package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("inp")
	numbers := []int{}
	for _, v := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}

	boards := getBoards(lines[2:])

outer:
	for _, number := range numbers {
		for i := range boards {
			boards[i].markNumber(number)
			if boards[i].checkWin() {
				fmt.Printf("Part 1: %d\n", number*boards[i].sumUnmarked())
				break outer
			}
		}
	}

	// reset boards
	boards = getBoards(lines[2:])

	done := []int{}
	for _, number := range numbers {
		for i := range boards {
			boards[i].markNumber(number)
			if boards[i].checkWin() {
				found := false
				for _, v := range done {
					if i == v {
						found = true
					}
				}
				if !found {
					done = append(done, i)
				}
			}
		}

		if len(done) == len(boards) {
			fmt.Printf("Part 2: %d\n", number*boards[done[len(done)-1]].sumUnmarked())
			break
		}
	}
}

type pair struct {
	seen bool
	val  int
}

type board [][]pair

func (b board) checkWin() bool {
f1:
	for _, row := range b {
		for _, p := range row {
			if p.seen == false {
				continue f1
			}
		}

		return true
	}

f2:
	for i := 0; i < len(b); i++ {
		for _, row := range b {
			if row[i].seen == false {
				continue f2
			}
		}

		return true
	}

	return false
}

func (b *board) markNumber(n int) {
	for i, row := range *b {
		for j, p := range row {
			if p.val == n {
				(*b)[i][j].seen = true
			}
		}
	}
}

func (b board) sumUnmarked() int {
	sum := 0
	for _, row := range b {
		for _, p := range row {
			if p.seen == false {
				sum += p.val
			}
		}
	}

	return sum
}

func getBoards(lines []string) []board {
	boards := []board{}
	b := board{}
	for _, line := range lines {
		if line == "" {
			boards = append(boards, b)
			b = board{}
			continue
		}
		numbers := []pair{}
		for _, s := range strings.Split(strings.Trim(line, " "), " ") {
			if s == "" {
				continue
			}
			n, err := strconv.Atoi(strings.Trim(s, " "))
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, pair{val: n})
		}

		b = append(b, numbers)
	}

	boards = append(boards, b)
	return boards
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
