package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
)

type coordinate struct {
	x int
	y int
}

type line struct {
	axis int
	distance int
	lineno int
}

type coordmap map[coordinate]int

func traceline(start coordinate, l line, grid coordmap, overlap []coordinate) (coordinate, []coordinate) {
	cur := start

	//If we have a negative distance we have to walk backwards
	walk := 1
	if l.distance < 0 {
		walk = -1
	}

	//Increment against current axis during walk
	//There's maybe a more elegant way to do this?
	icur := &(cur.x)
	if l.axis == 1 {
		icur = &(cur.y)
	}

	//We can't just use a range here (I think?) due to negative walks
	for i := 0; i != l.distance; i += walk {
		(*icur) += walk
		if val, exists := grid[cur]; exists {
			if val != l.lineno {
				overlap = append(overlap, cur)
			}
		}
		grid[cur] = l.lineno
	}

	return cur, overlap
}

func mappath(path []string, grid coordmap, overlap []coordinate, lineno int) ([]coordinate) {
	const x, y = 0, 1

	axis := x
	last := coordinate{0, 0}

	for _, l := range path {
		distance, _ := strconv.Atoi(l[1:])

		switch l[0] {
			case 'U':
				axis = x
			case 'D':
				axis = x
				distance *= -1
			case 'L':
				axis = y
				distance *= -1
			case 'R':
				axis = y
		}
		l := line{axis, distance, lineno}
		last, overlap = traceline(last, l, grid, overlap)
	}

	return overlap
}

func nearestmanhattan (overlap []coordinate) (int, coordinate) {
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	//This is basically just a bad min()
	var coord coordinate
	nearer := abs(overlap[0].x) + abs(overlap[0].y)
	for _, o := range overlap {
		manhattan := abs(o.x) + abs(o.y)
		if nearer > manhattan {
			nearer = manhattan
			coord = o
		}
	}

	return nearer, coord
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var in [][]string
	for i := 0; scanner.Scan(); i++ {
		tmp := strings.Split(scanner.Text(), ",")
		in = append(in, nil)
		in[i] = append(in[i], tmp...)
	}


	grid := make(coordmap)
	overlap := make([]coordinate, 0)

	for i, path := range in {
		overlap = mappath(path, grid, overlap, i)
	}

	n, _ := nearestmanhattan(overlap)

	fmt.Println(n)
}
