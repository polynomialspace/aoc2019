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
	Add = 1  // Add
	Mul = 2  // Multiply
	Inp = 3  // receive user Input
	Out = 4  // send Output to user
	JIT = 5  // Jump If True
	JIF = 6  // Jump If False
	TLT = 7  // Test Less Than
	TEQ = 8  // Test EQuals
	
	Hlt = 99 // Halt
)

type Instruction func (arg ...int) (out *int)

type op struct {
	inst Instruction // lambda to execute instruction
	len  int
	argp []int // argument position list
}


func intsplit(i int) (split []int) {
	for i > 1 {
		split = append(split, i % 10)
		i /= 10
	}
	split = append(split, i % 10)

	return
}

func mode(intcode int) ([]int) {
	// drop intcode itself, we only care about mode
	intcode /= 100

	mode := make([]int, 3)
	for i, m := range intsplit(intcode) {
		mode[i] = m
	}

	return mode
}

func (o op) exec (oplist []int) {
	outp := o.argp[len(o.argp)-1]

	args := make([]int, len(o.argp))
	for i, val := range o.argp {
		args[i] = oplist[val]
	}

	if ret := o.inst(args...); ret != nil {
		oplist[outp] = *ret
	}
}

func (o op) argspos(oplist []int, mode []int, offset int) (op) {
	const (
		position  = 0
		immediate = 1
	)

	args := make([]int, o.len - 1)

	for i := range args {
		pos := 1 + offset + i

		switch mode[i] {
		case immediate:
			args[i] = pos
		case position:
			args[i] = oplist[pos]
		}
	}

	o.argp = args

	return o
}

func fn(intcode int, off *int) (Instruction, int) {
	intcode %= 100 // drop mode

	var inst Instruction
	var length int

	switch intcode {
	case Hlt:
		inst = nil
	case Add:
		inst = func (arg ...int) (*int) {
			out := arg[0] + arg[1]
			return &out
		}	
		length = 4
	case Mul:
		inst = func (arg ...int) (*int) {
			out := arg[0] * arg[1]
			return &out
		}
		length = 4
	case Inp:
		inst = func (arg ...int) (*int) {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Printf("INPUT> ")
			scanner.Scan()
			out, _ := strconv.Atoi(scanner.Text())
			return &out
		}
		length = 2
	case Out:
		inst = func (arg ...int) (*int) {
			fmt.Printf("OUTPUT: %d\n", arg[0])
			return nil
		}
		length = 2
	case JIT:
		inst = func (arg ...int) (*int) {
			if arg[0] != 0 {
				*off = arg[1]
			}
			return nil
		}
		length = 3
	case JIF:
		inst = func (arg ...int) (*int) {
			if arg[0] == 0 {
				*off = arg[1]
			}
			return nil
		}
		length = 3
	case TLT:
		inst = func (arg ...int) (*int) {
			var out int
			if arg[0] < arg[1] {
				out = 1
			} else {
				out = 0
			}
			return &out
		}
		length = 4
	case TEQ:
		inst = func (arg ...int) (*int) {
			var out int
			if arg[0] == arg[1] {
				out = 1
			} else {
				out = 0
			}
			return &out
		}
		length = 4
	default:
		inst = nil
	}

	if inst == nil {
		fmt.Printf("Halting with opcode %d\n", intcode)
	}

	return inst, length
}

func parse(intcodes []int, offset *int) (opcode op, halt bool){
	var o op
	intcode := intcodes[*offset]
	o.inst, o.len = fn(intcode, offset)
	if o.inst == nil {
		return op{}, true
	}
	o = o.argspos(intcodes, mode(intcode), *offset)

	*offset += o.len
	return o, false
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	in := []string{}
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		in = strings.Split(scanner.Text(), ",")
	}

	intcodes := make([]int, len(in))
	for i,s := range in {
		intcodes[i], _ = strconv.Atoi(s)
	}

	off := 0
	o, hlt := parse(intcodes, &off)
	for !hlt {
		o.exec(intcodes)
		o, hlt = parse(intcodes, &off)
	}

}
