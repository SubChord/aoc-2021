package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readLines("inp")
	paths := make(map[string][]string)
	largeCaves := map[string]bool{}
	smallCaves := map[string]bool{}
	for _, line := range lines {
		split := strings.Split(line, "-")
		if split[0][0] >= 'a' {
			smallCaves[split[0]] = true
		} else {
			largeCaves[split[0]] = true
		}

		if split[1][0] >= 'a' {
			smallCaves[split[1]] = true
		} else {
			largeCaves[split[1]] = true
		}

		paths[split[0]] = append(paths[split[0]], split[1])
		paths[split[1]] = append(paths[split[1]], split[0])
	}

	pt1VisitsLeft := map[string]int{}
	for k, _ := range smallCaves {
		pt1VisitsLeft[k] = 1
	}

	// part 1
	p1 := findPath("start", smallCaves, largeCaves, paths, pt1VisitsLeft)
	fmt.Printf("Part 1: %d\n", len(p1))

	opts := map[string]bool{}
	for k, _ := range smallCaves {
		pt2VisitsLeft := map[string]int{}
		for k2, _ := range smallCaves {
			pt2VisitsLeft[k2] = 1
		}

		if k != "start" && k != "end" {
			pt2VisitsLeft[k] = 2
		}

		p2 := findPath("start", smallCaves, largeCaves, paths, pt2VisitsLeft)
		for _, p := range p2 {
			opts[strings.Join(p.nodes, "")] = true
		}
	}

	fmt.Printf("Part 2: %d\n", len(opts))
}

type path struct {
	nodes []string
	done  bool
}

func findPath(cave string, smallCaves map[string]bool, largeCaves map[string]bool, options map[string][]string, visitsLeft map[string]int) []path {
	// if end return
	if cave == "end" {
		return []path{path{nodes: []string{cave}, done: true}}
	}

	// if visited return
	if visitsLeft[cave] == 0 && smallCaves[cave] {
		return []path{}
	}

	visitsLeft[cave]--

	ret := []path{}
	for _, caveName := range options[cave] {
		copyVisited := make(map[string]int)
		for k, v := range visitsLeft {
			copyVisited[k] = v
		}
		paths := findPath(caveName, smallCaves, largeCaves, options, copyVisited)
		for _, p := range paths {
			if p.done {
				ret = append(ret, path{
					nodes: append([]string{cave}, p.nodes...),
					done:  true,
				})
			}
		}
	}

	return ret
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
