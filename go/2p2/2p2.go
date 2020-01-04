package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
)

const (
	Add = 1  // ArgPos, ArgPos, ResPos
	Mul = 2  // ArgPos, ArgPos, ResPos
	Hlt = 99 // -
)

func op(oplist []int, offset int) (halt bool){
/* Not much error handling here */
	op := oplist[offset]
	xpos := oplist[1+offset]
	ypos := oplist[2+offset]
	rpos := oplist[3+offset]

	switch op {
	case Hlt:
		return false
	case Add:
		x := oplist[xpos]
		y := oplist[ypos]
		oplist[rpos] = x + y
	case Mul:
		x := oplist[xpos]
		y := oplist[ypos]
		oplist[rpos] = x * y
	default:
		return false
	}

	return true
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
		in = strings.Split(scanner.Text(), ",")
	}

	oplist := make([]int, len(in))

	for i,s := range in {
		oplist[i], _ = strconv.Atoi(s)
	}

	orig := make([]int, len(oplist))
	copy(orig, oplist) // Technically should be checked?

	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			copy(oplist, orig)
			oplist[1] = noun
			oplist[2] = verb
			

			for offset := 0; len(oplist) > (offset + 3);
			 offset += 4 {
				valid := op(oplist, offset)
				if !valid {
					break
				}
			}

			if oplist[0] == 19690720 {
				out := 100 * noun + verb
				fmt.Println(out)
			}
		}
	}
}
