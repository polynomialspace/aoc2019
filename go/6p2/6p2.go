package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
)

type orbmap map[string][]string

type hop struct {
	name string
	dist int
}

func traceorbit(orbitmap orbmap, cur hop) (hops []hop) {
	hops = append(hops, cur)
	for _, v := range orbitmap[cur.name] {
		tmp := traceorbit(orbitmap, hop{v, 1+cur.dist})
		hops = append(hops, tmp...)
	}

	return hops
}

func nearestintersect(routes [2][]hop) (intersect []hop) {
	for _, v1 := range routes[0] {
		for _, v2 := range routes[1] {
			if v1.name == v2.name {
				intersect = []hop{v1, v2}
				return intersect
			}
		}
	}
	return nil
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	s := make([][]string, 0)
	for scanner.Scan() {
		in := strings.Split(scanner.Text(), ")")
		s = append(s, in)
	}

	rorbitmap := make(orbmap)
	for _, v := range s {
		/* Reversed orbit map */
		rorbitmap[v[1]] = append(rorbitmap[v[1]], v[0])
	}


	ends := []string{"YOU", "SAN"}
	routes := [2][]hop{{}}

	/* Reverse paths for "YOU" and "SAN" */
	routes[0] = traceorbit(rorbitmap, hop{ends[0], 0})
	routes[1] = traceorbit(rorbitmap, hop{ends[1], 0})

	/* first returned intersect ends up being the shortest */
	intersect := nearestintersect(routes)
	hopcnt := intersect[0].dist + intersect[1].dist - 2

	fmt.Println(hopcnt)
}
