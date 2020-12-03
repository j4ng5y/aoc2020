package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// twoNumberProduct returns the product of two integers from within a slice that add up to a particular value.
// If no two values from within the slice add up to the provided number, the return value is 0.
//
// Arguments:
//     number (int): The number that two integers should add up to.
//     slice ([]int): The slice of integers to parse
//
// Returns:
//     (int): Either the product of the two values from the "slice" that add up to "number", or zero
func twoNumberProduct(number int, slice []int) int {
	var val int
	// for each element in slice
	for i := range slice {
		// for each index of the previous slice
		for n := 0; n <= len(slice)-1; n++ {
			// if the sum of the two digits is number, set val to the product of those numbers
			if slice[i]+slice[n] == number {
				val = (slice[i] * slice[n])
			}
		}
	}
	return val
}

// threeNumberProduct returns the product of three integers from within a slice that add up to a particular value.
// If no three values from within the slice add up to the provided number, the return value is 0.
//
// Arguments:
//     number (int): The number that three integers should add up to.
//     slice ([]int): The slice of integers to parse
//
// Returns:
//     (int): Either the product of the three values from the "slice" that add up to "number", or zero
func threeNumberProduct(number int, slice []int) int {
	var val int
	// for each element in slice
	for i := range slice {
		// for each index of the previous slice
		for n := 0; n <= len(slice)-1; n++ {
			// and doing it again for good measure
			for nn := 0; nn <= len(slice)-1; nn++ {
				// the the sum of the three digits is number, set val to the product of those numbers
				if slice[i]+slice[n]+slice[nn] == number {
					val = (slice[i] * slice[n] * slice[nn])
				}
			}
		}
	}
	return val
}

func main() {
	// Basic argument parsing
	var file string
	var part int
	flag.StringVar(&file, "input", "", "The input file to parse")
	flag.IntVar(&part, "part", 0, "The part of the challenge to display (1/2), all other integers (or no value), displays both")
	flag.Parse()

	// If the provide input file is blank, error out
	if file == "" {
		log.Fatalln("You didn't provide an input file, please do so next time :D")
	}

	// Check that the file exists and is not a directory
	s, err := os.Stat(file)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening your input file, err: %+v", err))
	}
	if s.IsDir() {
		log.Fatal(fmt.Errorf("%s is an directory, please provide a file :D", file))
	}

	// Read in the file
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(fmt.Errorf("an error occurred reading your file, err %+v", err))
	}

	// Split the strings from the read in file into a slice on newlines
	strEntries := strings.Split(string(f), "\n")

	// For each line in strEntries:
	var intEntries []int
	for _, v := range strEntries {
		// convert the string value into an integer and error if it fails
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(fmt.Errorf("an error occurred converting %s to an integer, err: %+v", v, err))
		}
		// and append the value to the slice of ints
		intEntries = append(intEntries, i)
	}

	// run the specific functions to get the values requested and pin it to a string to print in just a second
	one := fmt.Sprintf("Day 1 / Part 1: %d", twoNumberProduct(2020, intEntries))
	two := fmt.Sprintf("Day 1 / Part 2: %d", threeNumberProduct(2020, intEntries))

	// finally, print all or some of the things
	switch part {
	case 1:
		fmt.Println(one)
	case 2:
		fmt.Println(two)
	default:
		fmt.Println(one)
		fmt.Println(two)

	}
}
