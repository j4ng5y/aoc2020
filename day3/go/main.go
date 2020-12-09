package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func buildFullMap(input []string) []string {
	var output = make([]string, len(input))
	for i, line := range input {
		output[i] = line
		for {
			if len(output[i]) < len(output)*7 {
				output[i] = output[i] + output[i]
				continue
			}
			break
		}
	}
	return output
}

func findTrees(r int, d int, filename string) (int, error) {
	acc := 0
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return acc, err
	}

	lines := strings.Split(string(f), "\n")
	fullMap := buildFullMap(lines)

	x := 0
	y := 0
	for i, line := range fullMap {
		if i == y {
			switch string(line[x]) {
			case "#":
				acc = acc + 1
			}
			y = y + d
			x = x + r
		}
	}

	return acc, nil
}

func main() {
	// This is a simple input parser
	var file string
	var part int
	flag.StringVar(&file, "input", "", "The input file to run against")
	flag.IntVar(&part, "part", 0, "The challenge part to print (1/2), everything else prints both")
	flag.Parse()

	// If the user didn't provide a file, error
	if file == "" {
		log.Fatalln("You didn't provide a file, please do so :D")
	}

	// Check that the file exists and isn't a directory
	stat, err := os.Stat(file)
	if err != nil {
		log.Fatalln(err)
	}
	if stat.IsDir() {
		log.Fatalln("Well, you provided something, but it's a directory. Please provide an input file. kthxbai")
	}

	// Right 1, down 1
	// Right 3, down 1
	// Right 5, down 1
	// Right 7, down 1
	// Right 1, down 2
	r1d1, err := findTrees(1, 1, file)
	r3d1, err := findTrees(3, 1, file)
	r5d1, err := findTrees(5, 1, file)
	r7d1, err := findTrees(7, 1, file)
	r1d2, err := findTrees(1, 2, file)
	if err != nil {
		log.Fatal(err)
	}

	switch part {
	case 1:
		fmt.Println("Day 3 / Part 1: ", r3d1)
	case 2:
		fmt.Println("Day 3 / Part 2: ", r1d1*r3d1*r5d1*r7d1*r1d2)
	default:
		fmt.Println("Day 3 / Part 1: ", r3d1)
		fmt.Println("Day 3 / Part 2: ", r1d1*r3d1*r5d1*r7d1*r1d2)
	}
}
