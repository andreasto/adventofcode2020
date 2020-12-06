package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	p1 := calc(lines, 3, 1)

	log.Printf("Number of trees encountered is: %d", p1)

	p2 := calc(lines, 1, 1)
	p3 := calc(lines, 5, 1)
	p4 := calc(lines, 7, 1)
	p5 := calc(lines, 1, 2)

	log.Printf("Number of trees encountered is: %d", p2*p1*p3*p4*p5)
}

func calc(input []string, right, down int) int {
	trees := 0
	index := right
	rowLength := len(input[0])
	for i := down; i < len(input); i += down {
		row := input[i]
		if string(row[index]) == "#" {
			trees++
		}

		index = (index + right) % rowLength
	}

	return trees
}
