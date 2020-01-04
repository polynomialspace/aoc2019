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

	oplist[1] = 12 // "replace position 1 with the value 12"
	oplist[2] = 2 // "and replace position 2 with the value 2"

	for offset := 0; len(oplist) > (offset + 3); offset += 4 {
		valid := op(oplist, offset)
		if !valid {
			break
		}
	}

	fmt.Println(oplist[0])
}
