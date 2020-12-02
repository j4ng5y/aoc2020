package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func twoNumberProduct(number int, slice []int) int {
	var val int
	for i := range slice {
		for n := 0; n <= len(slice)-1; n++ {
			if slice[i]+slice[n] == number {
				val = (slice[i] * slice[n])
			}
		}
	}
	return val
}

func threeNumberProduct(number int, slice []int) int {
	var val int
	for i := range slice {
		for n := 0; n <= len(slice)-1; n++ {
			for nn := 0; nn <= len(slice)-1; nn++ {
				if slice[i]+slice[n]+slice[nn] == number {
					val = (slice[i] * slice[n] * slice[nn])
				}
			}
		}
	}
	return val
}

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	strEntries := strings.Split(string(f), "\n")
	var intEntries []int
	for _, v := range strEntries {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		intEntries = append(intEntries, i)
	}

	fmt.Println(twoNumberProduct(2020, intEntries))
	fmt.Println(threeNumberProduct(2020, intEntries))
}
