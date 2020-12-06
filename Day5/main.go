package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seatIDS := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		seatIDS = append(seatIDS, getSeatID(strings.Split(scanner.Text(), "")))
	}

	sort.Ints(seatIDS)
	highestID := seatIDS[len(seatIDS)-1]

	for i, seat := range seatIDS {
		log.Println(seat)
		if seatIDS[i+1]-seat > 1 {
			log.Printf("My seat is: %d", seat+1)
			break
		}
	}

	log.Printf("Highets seatID is: %d", highestID)
}

func getSeatID(codes []string) int {
	rowCode := codes[0:7]
	seats := codes[7:]
	totalRows := []string{}

	for r := 0; r < 128; r++ {
		totalRows = append(totalRows, fmt.Sprint(r))
	}

	for _, rc := range rowCode {
		numRows := len(totalRows)
		if rc == "F" {
			totalRows = totalRows[0:(numRows / 2)]
		} else if rc == "B" {
			totalRows = totalRows[(numRows / 2):]
		}
	}

	totalSeats := []string{"0", "1", "2", "3", "4", "5", "6", "7"}

	for i := 0; i < 3; i++ {
		numSeats := len(totalSeats)
		if seats[i] == "L" {
			totalSeats = totalSeats[0:(numSeats / 2)]
		} else {
			totalSeats = totalSeats[(numSeats / 2):]
		}
	}

	parsedRow, _ := strconv.Atoi(totalRows[0])
	parsedSeat, _ := strconv.Atoi(totalSeats[0])

	return (parsedRow * 8) + parsedSeat
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
