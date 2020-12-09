package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type document struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func (d document) byrIsValid() bool {
	if d.byr == "" {
		return false
	}
	by, err := strconv.Atoi(d.byr)
	if err != nil {
		return false
	}

	switch {
	case by >= 1920 && by <= 2002:
		return true
	default:
		return false
	}
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func (d document) iyrIsValid() bool {
	if d.iyr == "" {
		return false
	}

	iy, err := strconv.Atoi(d.iyr)
	if err != nil {
		return false
	}

	switch {
	case iy >= 2010 && iy <= 2020:
		return true
	default:
		return false
	}
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func (d document) eyrIsValid() bool {
	if d.eyr == "" {
		return false
	}

	ey, err := strconv.Atoi(d.eyr)
	if err != nil {
		return false
	}

	switch {
	case ey >= 2020 && ey <= 2030:
		return true
	default:
		return false
	}
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func (d document) hgtIsValid() bool {
	if strings.HasSuffix(d.hgt, "cm") {
		h := strings.TrimSuffix(d.hgt, "cm")
		hint, err := strconv.Atoi(h)
		if err != nil {
			return false
		}
		if hint >= 150 && hint <= 193 {
			return true
		}
		return false
	}
	if strings.HasSuffix(d.hgt, "in") {
		h := strings.TrimSuffix(d.hgt, "in")
		hint, err := strconv.Atoi(h)
		if err != nil {
			return false
		}
		if hint >= 59 && hint <= 76 {
			return true
		}
		return false
	}
	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func (d document) hclIsValid() bool {
	r := regexp.MustCompile("#[a-f0-9]{6}")
	return r.MatchString(d.hcl)
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func (d document) eclIsValid() bool {
	switch d.ecl {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func (d document) pidIsValid() bool {
	r := regexp.MustCompile("^[0-9]{9}$")
	return r.MatchString(d.pid)
}

// cid (Country ID) - ignored, missing or not.
func (d document) isValid(part int) bool {
	switch part {
	case 1:
		if d.byr != "" && d.iyr != "" && d.eyr != "" && d.hgt != "" && d.hcl != "" && d.ecl != "" && d.pid != "" {
			return true
		}
		return false
	case 2:
		if !d.byrIsValid() {
			return false
		}
		if !d.iyrIsValid() {
			return false
		}
		if !d.eyrIsValid() {
			return false
		}
		if !d.hgtIsValid() {
			return false
		}
		if !d.hclIsValid() {
			return false
		}
		if !d.eclIsValid() {
			return false
		}
		if !d.pidIsValid() {
			return false
		}
		return true
	default:
		return false
	}
}

func parseFile(filename string) ([]*document, error) {
	var docs []*document
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return docs, err
	}

	lines := strings.Split(string(b), "\n")

	var t [][]string
	prev := 0
	for i, l := range lines {
		if l == "" {
			t = append(t, lines[prev:i])
			prev = i
		}
	}

	for _, v := range t {
		var doc = new(document)
		for _, vv := range v {
			elems := strings.Split(vv, " ")
			for _, elem := range elems {
				kv := strings.Split(elem, ":")
				switch strings.ToLower(string(kv[0])) {
				case "byr":
					doc.byr = kv[1]
				case "iyr":
					doc.iyr = kv[1]
				case "eyr":
					doc.eyr = kv[1]
				case "hgt":
					doc.hgt = kv[1]
				case "hcl":
					doc.hcl = kv[1]
				case "ecl":
					doc.ecl = kv[1]
				case "pid":
					doc.pid = kv[1]
				case "cid":
					doc.cid = kv[1]
				}
			}
		}
		docs = append(docs, doc)
	}

	return docs, nil
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

	d, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	switch part {
	case 1:
		validCount := 0
		for _, doc := range d {
			if doc.isValid(1) {
				validCount = validCount + 1
			}
		}
		fmt.Println("Day 4 / Part 1: ", validCount)
	case 2:
		validCount := 0
		for _, doc := range d {
			if doc.isValid(2) {
				validCount = validCount + 1
			}
		}
		fmt.Println("Day 4 / Part 2: ", validCount)
	default:
		validCount := 0
		for _, doc := range d {
			if doc.isValid(1) {
				validCount = validCount + 1
			}
		}
		fmt.Println("Day 4 / Part 1: ", validCount)
		validCount = 0
		for _, doc := range d {
			if doc.isValid(2) {
				validCount = validCount + 1
			}
		}
		fmt.Println("Day 4 / Part 2: ", validCount)
	}
}
