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
	Add = 1  // ArgPos , ArgPos, ResPos
	Mul = 2  // ArgPos , ArgPos, ResPos
	Inp = 3  // Ext In , ResPos
	Out = 4  // Ext Out, ArgPos
	
	Hlt = 99 // -
)

type Instruction func (a int, b int, out *int)

type op struct {
	opcode int
	inst Instruction //closure to execute instruction
	off  int   // offset in operation list, effective instptr
	len  int
	mode []int // Immediate, Position
	argp []int // argument position list
}

func opgetmode(o op) (op) {
	// drop opcode itself, we only care about mode
	opcode := o.opcode
	opcode /= 100

	umod := intsplit(uint(opcode))
	smod := make([]int, 3)
	for i, m := range umod {
		smod[(len(umod)-1)-i] = int(m)
	}
	o.mode = smod

	return o
}

func intsplit(u uint) (out []uint) {
	for u > 1 {
		out = append([]uint{u % 10}, out...)
		u /= 10
	}
	out = append([]uint{u % 10}, out...)

	return
}

func intrev(s []int) ([]int) {
	l := len(s)-1

	for i := 0; i < l/2; i++ {
		s[i],s[l-i] = s[l-i],s[i]
	}

	return s
}

func instrexec(o op, oplist []int) {
	out := o.argp[len(o.argp)-1]
	fmt.Printf("argp: %d\n", o.argp)
	a := oplist[o.argp[0]]
	b := 0
	if len(o.argp) > 2 {
		b = oplist[o.argp[1]]
	}
	fmt.Printf("Executing op [%d], argspos %d (%d %d), respos [%d] (%d)\n", o.opcode, o.argp, a, b, out, oplist[out])
	o.inst(a, b, &(oplist[out]))
	fmt.Printf("Execute complete: out: %d\n", oplist[out])
}

func opgetargspos(o op, oplist []int) (op) {
	args := make([]int, o.len - 1)
	//m := intrev(o.mode)
	m := o.mode
	offset := o.off

	fmt.Printf("rmod: %d\n", m)
	for i := 0; i < len(args); i++ {
		if m[i] == 1 {
			// immediate mode
			fmt.Println("Immediate mode: ", 1+offset+i)
			args[i] = 1 + offset + i
		} else {
			// position mode
			fmt.Println("Position mode")
			argpos := oplist[1+offset+i]
			args[i] = argpos
		}
	}
	o.argp = args

	return o
}

func opgetfn(o op) (op) {
	length := 0
	opcode := o.opcode % 100

	var inst Instruction
	switch opcode {
	case Hlt:
		inst = nil
	case Add:
		inst = func(a int, b int, out *int) {
			*out = a + b
		}	
		length = 4
	case Mul:
		inst = func(a int, b int, out *int) {
			*out = a * b
		}
		length = 4
	case Inp:
		inst = func(a int, b int, out *int) {
			//unused: b
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Printf("INPUT> ")
			scanner.Scan()
			s := scanner.Text()
			in, _ := strconv.Atoi(s)
			*out = in
		}
		length = 2
	case Out:
		inst = func(a int, b int, out *int) {
			//unused: b, out
			fmt.Printf("OUTPUT: %d\n", a)
		}
		length = 2
	default:
		inst = nil
	}

	o.len = length
	o.inst = inst

	if inst == nil {
		fmt.Printf("Halting with opcode %d\n", opcode)
	}

	return o
}

func opparse(oplist []int, offset *int) (opcode op, halt bool){
	o := op{off: *offset}
	o.opcode = oplist[*offset]
	fmt.Printf("Starting command at offset %d [%d]\n", o.off, o.opcode)
	o = opgetfn(o)
	if o.inst == nil {
		return op{}, false
	}
	o = opgetmode(o)
	o = opgetargspos(o, oplist)

	*offset += o.len
	
	return o, true
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

	off := 0
	o, hlt := opparse(oplist, &off)
	for hlt {
		instrexec(o, oplist)
		o, hlt = opparse(oplist, &off)
	}

}
