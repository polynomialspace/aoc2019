package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
)

type orbmap map[string][]string

func recursemapcnt(orbitmap orbmap, orbit string) (int) {
	cnt := 0
	for _, v := range orbitmap[orbit] {
		cnt += recursemapcnt(orbitmap, v)
		cnt++
	}
	return cnt
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

	orbitmap := make(orbmap)

	for _, v := range s {
		orbitmap[v[0]] = append(orbitmap[v[0]], v[1])
	}

	cnt := 0
	for k, _ := range orbitmap {
		cnt += recursemapcnt(orbitmap, k)
	}
	fmt.Println(cnt)
}
