package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
)

func validateint(i int) (ok bool) {
	return true
}

func intsplit(u uint) (out []uint) {
	for u > 10 {
		out = append([]uint{u % 10}, out...)
		u /= 10
	}
	out = append([]uint{u % 10}, out...)

	return
}

func intjoin(u []uint) (out uint) {
	p := uint(1)

	i := len(u) - 1
	for i > 0 {
		out += (u[i] * p)
		p *= 10
		i--
	}
	out += (u[i] * p)

	return
}

func checksplint(u []uint) (ok bool) {
	// Valid integer's digits must meet criteria of:
	// each consecutive digit must be equal-or-greater than previous digit
	// there must be at least one pair of consecutive digits that are
	//  the same
	// there must be exactly 6 digits

	rep := false
	ok = true

	if len(u) != 6 {
		ok = false
		return
	}

	prev := u[0]
	for _, d := range u[1:] {
		if d < prev {
			ok = false
		}
		if d == prev {
			rep = true
		}
		prev = d
	}

	if rep == false {
		ok = false
	}

	return
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var in []string
	for scanner.Scan() {
		in = strings.Split(scanner.Text(), "-")
	}
	
	Range := make([]int, len(in))
	for i, a := range in {
		Range[i], _ = strconv.Atoi(a)
	}


	validints := 0
	for i := Range[0]; i < Range[1]; i++ {
		tmp := intsplit(uint(i))
		if ok := checksplint(tmp); ok {
			validints++
		}
	}

	fmt.Println(validints)
}
