package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func getfuel(mass int) (fuel int) {
	fuel = (mass / 3) - 2
	return fuel
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuel := getfuel(mass)
		total += uint64(fuel)
	}

	fmt.Println(total)
}
