package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, readFileErr := ioutil.ReadFile("input.txt")
	if readFileErr != nil {
		panic(readFileErr)
	}
	lines := strings.Split(string(data), "\n\n")

	totalCount := 0
	onlyYes := 0
	for _, l := range lines {
		totalCount += part1(l)
		onlyYes += part2(l)
	}

	log.Printf("Total Yes answers is: %d", totalCount)
	log.Printf("Total Questions with only yes answeres is: %d", onlyYes)
}

func part1(answers string) int {
	m := map[string]int{}

	s := strings.Split(answers, "")

	total := 0
	for _, a := range s {
		if _, ok := m[a]; ok || a == "\n" {
			continue
		}
		m[a] = 1
		total++
	}

	return total
}

func part2(answers string) int {
	m := map[string]int{}
	s := strings.Split(answers, "")

	total := 0
	personsInGroup := 1
	for _, a := range s {
		if a == "\n" {
			personsInGroup++
			continue
		}
		if val, ok := m[a]; ok {
			m[a] = val + 1
		} else {
			m[a] = 1
		}
	}

	for _, v := range m {
		if v == personsInGroup {
			total++
		}
	}

	return total
}
