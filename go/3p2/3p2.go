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
	start coordinate
	axis int
	distance int
	lineno int
	len int
}

type coordmap map[coordinate]int

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func traceline(l line, grid coordmap, overlap coordmap) (coordinate, coordmap) {
	cur := l.start

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
				overlap[cur] += (1 + l.len + abs(i))
			}
		}
		grid[cur] = l.lineno
	}

	return cur, overlap
}

func mappath(path []string, grid coordmap, overlap coordmap, lineno int) (coordmap) {
	const x, y = 0, 1

	axis := x
	last := coordinate{0, 0}
	len := 0

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
		l := line{last, axis, distance, lineno, len}
		last, overlap = traceline(l, grid, overlap)
		len += abs(distance)
		
	}

	return overlap
}

func nearest (distances []int) (int) {
	//This is basically just a bad min()
	nearer := distances[0]
	for _, d := range distances {
		if nearer > d {
			nearer = d
		}
	}

	return nearer
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
	overlap := make(coordmap)

	// We have two paths, but we need to calculate overlaps for both,
	// Since we can't know the overlaps for the first line without running
	// the second, we rerun the first 
	for i, path := range in {
		overlap = mappath(path, grid, overlap, i)
	}
	overlap = mappath(in[0], grid, overlap, 0)

	distances := make([]int, 0)
	for _, val := range overlap {
		distances = append(distances, val)
	}

	n := nearest(distances)

	fmt.Println(n)

}
