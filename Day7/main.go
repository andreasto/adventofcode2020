package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name  string
	Count int
}

type Bag struct {
	Name     string
	Count    int
	Children []Bag
}

func main() {
	data, readFileErr := ioutil.ReadFile("input.txt")
	if readFileErr != nil {
		panic(readFileErr)
	}
	lines := strings.Split(string(data), "\n")

	b := []Parent{}

	for _, l := range lines {
		cases := strings.Split(l, " contain ")

		parentColor := strings.ReplaceAll(cases[0], " bags", "")
		nb := Parent{
			Name: parentColor,
		}
		children := []Child{}
		for i, c := range cases {
			if i == 0 {
				continue
			}
			newBagString := strings.Split(c, ", ")
			for _, nbs := range newBagString {
				newSplit := strings.Split(nbs, "")
				count := newSplit[0]
				parsedCount, err := strconv.Atoi(count)

				if err != nil {
					continue
				}

				clean := strings.Join(newSplit[1:], "")
				r := strings.NewReplacer("bags", "",
					"bags,", "",
					"bag", "",
					"bags.", "",
					"bag.", "",
					".", "",
				)
				clean = r.Replace(clean)
				clean = strings.Trim(clean, " ")
				cb := Child{
					Name:  clean,
					Count: parsedCount,
				}
				children = append(children, cb)
			}
			nb.Children = children
		}
		b = append(b, nb)
	}

	total := 0
	totalPartTwo := 0

	for _, t := range b {
		hasGolden := canHaveGoldenBag(t.Name, b)

		if hasGolden {
			total++
		}
	}

	for _, t := range b {
		if t.Name == "shiny gold" {
			for _, c := range t.Children {
				totalPartTwo += c.Count + (c.Count * checkChild(b, c))
			}
			break
		}
	}

	log.Printf("Total of cases that can contain Shiny golden cases: %d", total)
	log.Printf("Part 2 result: %d", totalPartTwo)
}

func canHaveGoldenBag(name string, bags []Parent) bool {
	singleBag := Parent{}
	for _, b := range bags {
		if b.Name == name {
			singleBag = b
			break
		}
	}

	temp := false
	for _, cb := range singleBag.Children {
		if cb.Name == "shiny gold" {
			temp = true
		}
		if canHaveGoldenBag(cb.Name, bags) {
			temp = true
		}
	}

	return temp
}

func checkChild(parents []Parent, bag Child) int {
	totalCount := 0

	for _, p := range parents {
		if p.Name == bag.Name {
			if len(p.Children) == 0 {
				return totalCount
			}
			for _, child := range p.Children {
				totalCount += child.Count + (child.Count * checkChild(parents, child))
			}
		}
	}

	return totalCount
}
