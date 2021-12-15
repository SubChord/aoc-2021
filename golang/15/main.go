package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("inp")
	g := parseGrid(lines)

	risk(g)
	//fmt.Printf("Part 1: %d\n", risk(g))
	fmt.Printf("Part 1: %d\n", dijkstra(g, node{}, node{x: len(g[0]) - 1, y: len(g) - 1}))

	g2 := make(grid, len(g)*5)
	for i := 0; i < len(g2); i++ {
		g2[i] = make([]int, len(g[0])*5)
	}

	for y := 0; y < len(g2); y++ {
		yIncr := y / len(g)
		for x := 0; x < len(g2[y]); x++ {
			xIncr := x / len(g[0])
			g2[y][x] = (g[y%len(g)][x%len(g[0])] + yIncr + xIncr) % 9
			if g2[y][x] == 0 {
				g2[y][x] = 9
			}

		}
	}

	fmt.Printf("Part 2: %d\n", dijkstra(g2, node{}, node{x: len(g2[0]) - 1, y: len(g2) - 1}))
}

func dijkstra(grid grid, source node, target node) int {
	dist := make(map[node]int)
	q := make(PriorityQueue, 0)

	heap.Push(&q, &Item{value: nodeWithDist{
		node: source,
		dist: 0,
	}, priority: 0})

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			n := node{x: x, y: y}
			dist[n] = math.MaxInt32
		}
	}

	dist[source] = 0

	for len(q) > 0 {
		nodeWithDistance := heap.Pop(&q).(*Item).value.(nodeWithDist)
		d := nodeWithDistance.dist
		if nodeWithDistance.node == target {
			return d
		}

		neighbours := getNeigbours(len(grid)-1, len(grid[0])-1, nodeWithDistance)
		for i := 0; i < len(neighbours); i++ {
			neighbour := neighbours[i]
			newDistance := d + grid[neighbour.y][neighbour.x]
			if dist[neighbour] > newDistance {
				dist[neighbour] = newDistance
				heap.Push(&q, &Item{
					value:    nodeWithDist{node: neighbour, dist: newDistance},
					priority: newDistance,
				})
			}
		}
	}

	return dist[target]
}

func getNeigbours(ymax int, xmax int, nodeWithDistance nodeWithDist) []node {
	neigbors := []node{}
	if nodeWithDistance.node.x > 0 {
		neigbors = append(neigbors, node{x: nodeWithDistance.node.x - 1, y: nodeWithDistance.node.y})
	}
	if nodeWithDistance.node.x < xmax {
		neigbors = append(neigbors, node{x: nodeWithDistance.node.x + 1, y: nodeWithDistance.node.y})
	}
	if nodeWithDistance.node.y > 0 {
		neigbors = append(neigbors, node{x: nodeWithDistance.node.x, y: nodeWithDistance.node.y - 1})
	}
	if nodeWithDistance.node.y < ymax {
		neigbors = append(neigbors, node{x: nodeWithDistance.node.x, y: nodeWithDistance.node.y + 1})
	}
	return neigbors
}

func risk(g grid) int {
	// zero based grid with same size as g
	newGrid := make(grid, len(g))
	for i := range newGrid {
		newGrid[i] = make([]int, len(g[i]))
	}

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if x != 0 || y != 0 {
				newGrid[y][x] = g[y][x]
			}

			if y == 0 && x > 0 {
				newGrid[y][x] += newGrid[y][x-1]
			} else if x == 0 && y > 0 {
				newGrid[y][x] += newGrid[y-1][x]
			} else if x > 0 && y > 0 {
				top := newGrid[y-1][x]
				left := newGrid[y][x-1]

				newGrid[y][x] += int(math.Min(float64(top), float64(left)))
			}
		}
	}

	return newGrid[len(g)-1][len(g[0])-1]
}

func parseGrid(lines []string) grid {
	g := grid{}
	for _, line := range lines {
		ints := []int{}
		ints2 := []int{}
		split := strings.Split(line, "")
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, i)
			ints2 = append(ints2, i)
		}
		g = append(g, ints)
	}
	return g
}

type grid [][]int

type node struct {
	x, y int
}

type nodeWithDist struct {
	node
	dist int
}

func (g grid) print() {
	for _, row := range g {
		for _, i := range row {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
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
