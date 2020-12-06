package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"regexp"
)

func main() {
	data, readFileErr := ioutil.ReadFile("input.txt")
	if readFileErr != nil {
		panic(readFileErr)
	}
	lines := strings.Split(string(data), "\n\n")
	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	log.Printf("The number of valid passports are: %d", part1(lines, validFields))
	log.Printf("The number of valid passports are: %d", part2(lines, validFields))
}

func part1(lines, validFields []string) int {
	correctPassports := 0
	for _, p := range lines {
		numberOfValidFields := 0
		for _, vf := range validFields {
			for _, f := range strings.Split(string(p), "\n") {
				if strings.Contains(f, vf) || vf == "cid" {
					numberOfValidFields++
				}
			}
		}
		if numberOfValidFields >= len(validFields) {
			correctPassports++
		}
	}
	return correctPassports
}

func part2(lines, validFields []string) int {
	correctPassports := 0
	for _, p := range lines {
		numberOfValidFields := 0
		fields := strings.Fields(p)
		for _, vf := range validFields {
			for _, f := range fields {
				if strings.Contains(f, vf) || vf == "cid" {
					if validate(f) {
						numberOfValidFields++
					}
					break
				}
			}
		}
		if numberOfValidFields >= len(validFields) {
			correctPassports++
		}
	}
	return correctPassports
}

func validate(f string) bool {
	fieldValid := false
	if f[0:3] == "byr" {
		i, _ := strconv.Atoi(f[4:])
		if i >= 1920 && i <= 2002 {
			fieldValid = true
		}
	} else if f[0:3] == "iyr" {
		i, _ := strconv.Atoi(f[4:])
		if i >= 2010 && i <= 2020 {
			fieldValid = true
		}
	} else if f[0:3] == "eyr" {
		i, _ := strconv.Atoi(f[4:])
		if i >= 2020 && i <= 2030 {
			fieldValid = true
		}
	} else if f[0:3] == "hgt" {
		value := []string{f[4 : len(f)-2], f[len(f)-2:]}
		length, _ := strconv.Atoi(value[0])
		unit := value[1]
		if unit == "in" && length >= 59 && length <= 76 {
			fieldValid = true
		} else if unit == "cm" && length >= 150 && length <= 193 {
			fieldValid = true
		}
	} else if f[0:3] == "hcl" {
		matched, _ := regexp.MatchString(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`, f[4:])
		if matched {
			fieldValid = true
		}
	} else if f[0:3] == "ecl" {
		validEyeColor := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		valid := false
		for _, c := range validEyeColor {
			if strings.Contains(f[4:], c) {
				valid = true
				break
			}
		}
		if valid {
			fieldValid = true
		}
	} else if f[0:3] == "pid" {
		matched, _ := regexp.MatchString(`^[0-9]*$`, f[4:])
		if matched && len(f[4:]) == 9 {
			fieldValid = true
		}
	} else if f[0:3] == "pid" {
		fieldValid = true
	}
	return fieldValid
}