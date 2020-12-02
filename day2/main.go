package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Rule struct {
	Letter      string
	PositionOne int
	PositionTwo int
	Pattern     string
}

func (R *Rule) IsValidSledRental() bool {
	switch {
	case strings.Count(strings.ToLower(R.Pattern), strings.ToLower(R.Letter)) < R.PositionOne:
		return false
	case strings.Count(strings.ToLower(R.Pattern), strings.ToLower(R.Letter)) > R.PositionTwo:
		return false
	default:
		return true
	}
}

func (R *Rule) IsValidTobogganCorp() bool {
	pos1 := strings.ToLower(string(R.Pattern[R.PositionOne-1]))
	pos2 := strings.ToLower(string(R.Pattern[R.PositionTwo-1]))
	l := strings.ToLower(R.Letter)
	switch {
	case pos1 == l && pos2 == l:
		return false
	case pos1 == l, pos2 == l:
		return true
	default:
		return false
	}
}

func NewRules(b []byte) (R []*Rule, err error) {
	lines := strings.Split(string(b), "\n")
	for _, v := range lines {
		line := strings.Split(v, " ")
		minmax := strings.Split(line[0], "-")
		pos1, err := strconv.Atoi(minmax[0])
		if err != nil {
			return R, err
		}
		pos2, err := strconv.Atoi(minmax[1])
		if err != nil {
			return R, err
		}
		letter := strings.TrimSuffix(line[1], ":")
		pattern := line[2]
		R = append(R, &Rule{
			Letter:      letter,
			PositionOne: pos1,
			PositionTwo: pos2,
			Pattern:     pattern,
		})
	}
	return R, nil
}

func main() {

	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	r, err := NewRules(f)
	if err != nil {
		log.Fatal(err)
	}

	SledRentalPlace := 0
	TobogganCorp := 0
	for _, v := range r {
		if v.IsValidSledRental() {
			SledRentalPlace = SledRentalPlace + 1
		}
		if v.IsValidTobogganCorp() {
			TobogganCorp = TobogganCorp + 1
		}
		continue
	}
	fmt.Printf("Sled Rental Count: %d\n", SledRentalPlace)
	fmt.Printf("Toboggan Corp Count: %d\n", TobogganCorp)
}
