package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

func (p passport) isSet() bool {
	return !(p.byr == "" || p.iyr == "" || p.eyr == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "")
}

// Return true if fields are valid by the following rules:
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
//   If cm, the number must be at least 150 and at most 193.
//   If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func (p passport) isValid() bool {

	if !p.isSet() {
		return false
	}

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	//   If cm, the number must be at least 150 and at most 193.
	//   If in, the number must be at least 59 and at most 76
	if strings.HasSuffix(p.hgt, "cm") {
		hgt, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if err != nil {
			return false
		}
		if hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		hgt, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if err != nil {
			return false
		}
		if hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	r := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if !r.MatchString(p.hcl) {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	switch p.ecl {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	r = regexp.MustCompile(`^[0-9]{9}$`)
	if !r.MatchString(p.pid) {
		return false
	}

	// cid (Country ID) - ignored, missing or not.

	return true
}

func readFileToPassports(path string) ([]passport, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var passports []passport
	scanner := bufio.NewScanner(file)
	currPassport := passport{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// if there is a blankline, that should be considered the end of
		// a passport entry
		if len(line) <= 0 {
			passports = append(passports, currPassport)
			currPassport = passport{}
		} else {
			// parse the line. spaces are delimiters, keys are split on ':'
			for _, v := range strings.Split(line, " ") {
				kvp := strings.Split(v, ":")
				fieldVal := kvp[1]
				switch field := kvp[0]; field {
				case "byr":
					currPassport.byr = fieldVal
				case "iyr":
					currPassport.iyr = fieldVal
				case "eyr":
					currPassport.eyr = fieldVal
				case "hgt":
					currPassport.hgt = fieldVal
				case "hcl":
					currPassport.hcl = fieldVal
				case "ecl":
					currPassport.ecl = fieldVal
				case "pid":
					currPassport.pid = fieldVal
				case "cid":
					currPassport.cid = fieldVal
				default:
					fmt.Printf("'%s'\n", field)
					return nil, errors.New("Unrecognized Field: " + field)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// save the passport we were building at the end of file
	passports = append(passports, currPassport)

	return passports, nil
}

// Determine the number of trees you would encounter if you traversed with a
// slope of  Right 3, down 1.
func part1(input []passport) int {
	fieldSet := 0
	for _, passport := range input {
		if passport.isSet() {
			fieldSet++
		}

	}
	return fieldSet
}

func part2(input []passport) int {
	valid := 0
	for _, passport := range input {
		if passport.isValid() {
			valid++
		}

	}
	return valid
}

func main() {
	passports, err := readFileToPassports("/2020/aoc04/input")
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(passports)) // 226
	fmt.Print("Part 2: ")
	fmt.Println(part2(passports)) // 160
}
