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

type passwordPattern struct {
	Letter      string
	PositionOne int
	PositionTwo int
	Pattern     string
}

// isValidSledRental checks a parent structure for adherence to the "sled company up the road"'s password policy.
//
// The number in "PositionOne" of the struct is the minimum number of times an occurrence of "Letter" should exist (case 1).
// The number in "PositionTwo" of the struct is the maximum number of times an occurrence of "Letter" should exist (case 2).
// If both cases do not match, the password adheres to the policy.
//
// Arguments:
//     None
//
// Returns:
//     (bool): Whether or not the password provided adheres to the "sled company up the road"'s password policy.
func (P *passwordPattern) isValidSledRental() bool {
	switch {
	case strings.Count(strings.ToLower(P.Pattern), strings.ToLower(P.Letter)) < P.PositionOne:
		return false
	case strings.Count(strings.ToLower(P.Pattern), strings.ToLower(P.Letter)) > P.PositionTwo:
		return false
	default:
		return true
	}
}

// isValidTobogganCorp checks a parent structure for adherence to Toboggan Corporation's password policy.
//
// The number in "PositionOne" of the struct is the (1 indexed) position that "Letter" should exist.
// The number in "PositionTwo" of the struct is the (1 indexed) position that "Letter" should exist.
// If either "PositionOne" or "PositionTwo" match "Letter", but not BOTH are set to "Letter", the password adheres to the policy.
// If both are set to "Letter", the password does not adhere to the policy.
//
// Arguments:
//     None
//
// Returns:
//     (bool): Whether or not the password provided adheres to Toboggan Corporation's password policy.
func (P *passwordPattern) isValidTobogganCorp() bool {
	pos1 := strings.ToLower(string(P.Pattern[P.PositionOne-1]))
	pos2 := strings.ToLower(string(P.Pattern[P.PositionTwo-1]))
	l := strings.ToLower(P.Letter)
	switch {
	case pos1 == l && pos2 == l:
		return false
	case pos1 == l, pos2 == l:
		return true
	default:
		return false
	}
}

// newPasswordPattern is a function that creates a new slice of pointers to passwordPattern structures from a provided input file.
//
// 1. First, we read the file and error if that is unsuccessful
// 2. Then, we do a number of string conversions (and error out if necessary) to get the necessary fields to build a passwordPattern struct
// 3. Finally, we append the pointer to the newly created passwordPattern structure and return it.
//
// Arguments:
//     f (string): The name of the file to parse
//
// Returns:
//     P ([]*passwordPattern): The final slice of pointers to passwordPattern structures, nil if an error occurred
//     err (error): An error if one exists, nil otherwise
func newPasswordPattern(f string) (P []*passwordPattern, err error) {
	// Read the file to bytes and error out if necessary
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("something went awry reading your input file, err: %+v", err)
	}

	// Split the byte string into individual lines
	lines := strings.Split(string(b), "\n")

	// For each line in lines:
	for i, v := range lines {
		// Split the line based on spaces
		line := strings.Split(v, " ")

		// Split the PositionOne and PositionTwo out of the first element of the slice
		pos := strings.Split(line[0], "-")

		// Convert the string of PostionOne to an integer and error out if the string can't be converted
		pos1, err := strconv.Atoi(pos[0])
		if err != nil {
			return P, fmt.Errorf("(%s: line %d) %s can't be converted to an integer", f, i, pos[0])
		}

		// Convert the string on PostionTwo to an integer and error out if the string can't be converted
		pos2, err := strconv.Atoi(pos[1])
		if err != nil {
			return P, fmt.Errorf("(%s: line %d) %s can't be converted to an integer", f, i, pos[1])
		}

		// Trim the ending colon from the character match
		letter := strings.TrimSuffix(line[1], ":")

		// Set the password pattern to validate
		pattern := line[2]

		// Build the new passwordPattern struct and append the pointer to the slice
		P = append(P, &passwordPattern{
			Letter:      letter,
			PositionOne: pos1,
			PositionTwo: pos2,
			Pattern:     pattern,
		})
	}

	// Return the slice of pointers to passwordPattern structures
	return P, nil
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

	// build the new slice of pointers to the passwordPattern struct and error if necessary
	p, err := newPasswordPattern(file)
	if err != nil {
		log.Fatal(fmt.Errorf("parsing the input file failed for some reason, err: %+v", err))
	}

	// initialize counters
	SledRentalPlace := 0
	TobogganCorp := 0
	// run validations
	for _, v := range p {
		if v.isValidSledRental() {
			SledRentalPlace = SledRentalPlace + 1
		}
		if v.isValidTobogganCorp() {
			TobogganCorp = TobogganCorp + 1
		}
		continue
	}
	// print the output
	src := fmt.Sprintf("Sled Rental Count: %d\n", SledRentalPlace)
	tcc := fmt.Sprintf("Toboggan Corp Count: %d\n", TobogganCorp)
	switch part {
	case 1:
		fmt.Print(src)
	case 2:
		fmt.Print(tcc)
	default:
		fmt.Print(src)
		fmt.Print(tcc)
	}
}
