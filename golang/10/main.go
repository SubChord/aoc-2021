package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	lines := readLines("inp")

	pt1ScoreMap := make(map[string]int)
	pt2Scores := []int{}

lineLoop:
	for _, line := range lines {
		stack := make([]string, 0)
		for _, c := range line {
			switch c {
			case '(':
				stack = append(stack, string(')'))
			case '{':
				stack = append(stack, string('}'))
			case '[':
				stack = append(stack, string(']'))
			case '<':
				stack = append(stack, string('>'))
			default:
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if string(c) != pop {
					pt1ScoreMap[string(c)]++
					continue lineLoop
				}
			}
		}

		// part 2
		score := 0
		// loop over stack in revers
		for i := len(stack) - 1; i >= 0; i-- {
			score *= 5
			if stack[i] == ")" {
				score += 1
			} else if stack[i] == "]" {
				score += 2
			} else if stack[i] == "}" {
				score += 3
			} else if stack[i] == ">" {
				score += 4
			}
		}
		pt2Scores = append(pt2Scores, score)
	}

	//): 3 points.
	//]: 57 points.
	//}: 1197 points.
	//>: 25137 points.

	s := 0
	for k, v := range pt1ScoreMap {
		switch k {
		case ")":
			s += 3 * v
		case "]":
			s += 57 * v
		case "}":
			s += 1197 * v
		case ">":
			s += 25137 * v
		}
	}

	fmt.Printf("Part 1: %d\n", s)

	// find middle number of scores
	sort.Ints(pt2Scores)
	middleIndex := len(pt2Scores) / 2
	fmt.Printf("Part 2: %d", pt2Scores[middleIndex])
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
