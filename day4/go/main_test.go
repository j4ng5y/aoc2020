package main

import "testing"

func Test_byrIsValid(t *testing.T) {
	good := document{
		byr: "1987",
	}
	if !good.byrIsValid() {
		t.Fail()
	}
	bad := document{
		byr: "1919",
	}
	if bad.byrIsValid() {
		t.Fail()
	}
}

func Test_iyrIsValid(t *testing.T) {
	good := document{
		iyr: "2015",
	}
	if !good.iyrIsValid() {
		t.Fail()
	}
	bad := document{
		iyr: "2030",
	}
	if bad.iyrIsValid() {
		t.Fail()
	}
}

func Test_eyrIsValid(t *testing.T) {
	good := document{
		eyr: "2022",
	}
	if !good.eyrIsValid() {
		t.Fail()
	}
	bad := document{
		eyr: "1919",
	}
	if bad.eyrIsValid() {
		t.Fail()
	}
}

func Test_hgtIsValid(t *testing.T) {
	// hgt valid:   60in
	// hgt valid:   190cm
	// hgt invalid: 190in
	// hgt invalid: 190

	good := document{
		hgt: "60in",
	}
	if !good.hgtIsValid() {
		t.Fail()
	}
	good = document{
		hgt: "190cm",
	}
	if !good.hgtIsValid() {
		t.Fail()
	}
	bad := document{
		hgt: "190in",
	}
	if bad.hgtIsValid() {
		t.Fail()
	}
	bad = document{
		hgt: "190",
	}
	if bad.hgtIsValid() {
		t.Fail()
	}
}

func Test_hclIsValid(t *testing.T) {
	// hcl valid:   #123abc
	// hcl invalid: #123abz
	// hcl invalid: 123abc

	good := document{
		hcl: "#123abc",
	}
	if !good.hclIsValid() {
		t.Fail()
	}
	bad := document{
		hcl: "#123abz",
	}
	if bad.hclIsValid() {
		t.Fail()
	}
	bad = document{
		hcl: "123abc",
	}
	if bad.hclIsValid() {
		t.Fail()
	}
}

func Test_eclIsValid(t *testing.T) {
	good := document{
		ecl: "amb",
	}
	if !good.eclIsValid() {
		t.Fail()
	}
	bad := document{
		ecl: "amber",
	}
	if bad.eclIsValid() {
		t.Fail()
	}
}

func Test_pidIsValid(t *testing.T) {
	// pid valid:   000000001
	// pid invalid: 0123456789

	good := document{
		pid: "000000001",
	}
	bad := document{
		pid: "0123456789",
	}

	if !good.pidIsValid() {
		t.Logf("good pid expected to pass validation, but didn't")
		t.Fail()
	}
	if bad.pidIsValid() {
		t.Logf("bad pid expected to fail validation, but didn't")
		t.Fail()
	}
}
