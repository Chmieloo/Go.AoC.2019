package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day3.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var line string
	var slice []string
	var instructions = make(map[int][]string)
	//var array = make(map[int8]int8)
	i := 0

	for {
		line, err = r.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			break
		}
		slice = strings.Split(line, ",")
		instructions[i] = slice
		i++
	}

	var direction string
	var steps, px, py, j int
	var visitedCoordinates = make(map[string]int)
	px = 0
	py = 0
	length := 0

	fmt.Println(instructions[0])

	for _, ins := range instructions[0] {
		direction = ins[0:1]
		steps, _ = strconv.Atoi(ins[1:])
		length++

		switch direction {
		case "R":
			for j = px; j < steps; j++ {
				visitedCoordinates[strconv.Itoa(j)+"."+strconv.Itoa(py)] = length
			}
			px += steps
		case "L":
			for j = px; j < steps; j-- {
				visitedCoordinates[strconv.Itoa(j)+"."+strconv.Itoa(py)] = length
			}
			px -= steps
		case "D":
			for j = py; j < steps; j-- {
				visitedCoordinates[strconv.Itoa(px)+"."+strconv.Itoa(j)] = length
			}
			py -= steps
		case "U":
			for j = py; j < steps; j++ {
				visitedCoordinates[strconv.Itoa(px)+"."+strconv.Itoa(j)] = length
			}
			py += steps
		}

		fmt.Println(ins, px, py)
	}

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
