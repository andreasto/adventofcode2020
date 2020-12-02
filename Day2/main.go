package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	validPasswordsPart1 := 0
	validPasswordsPart2 := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		validPasswordsPart1 = validPasswordsPart1 + CheckPassword(scanner.Text())
		validPasswordsPart2 = validPasswordsPart2 + CheckPassWordPartTwo(scanner.Text())
	}

	log.Printf("1st Part: %d", validPasswordsPart1)
	log.Printf("2nd Part: %d", validPasswordsPart2)
}

func CheckPassword(text string) int {
	i := strings.Index(text, ":")

	count := 0
	for _, l := range strings.Split(text[i+2:], "") {
		search := text[i-1 : i]
		if l == search {
			count = count + 1
		}
	}

	r := strings.Split(text[0:i-2], "-")
	low, _ := strconv.Atoi(r[0])
	high, _ := strconv.Atoi(r[1])

	if count > 0 && count >= low && count <= high {
		return 1
	}
	return 0
}

func CheckPassWordPartTwo(text string) int {
	i := strings.Index(text, ":")
	password := strings.Split(text[i+2:], "")
	key := text[i-1 : i]

	r := strings.Split(text[0:i-2], "-")
	firstValue, _ := strconv.Atoi(r[0])
	secondvalue, _ := strconv.Atoi(r[1])

	firstIndex := firstValue - 1
	secondIndex := secondvalue - 1

	if password[firstIndex] == key && password[secondIndex] != key || password[firstIndex] != key && password[secondIndex] == key {
		return 1
	}

	return 0
}
