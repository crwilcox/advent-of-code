package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadFileToLines takes a path and outputs the contents to a line array.
func ReadFileToLines(path string) ([]string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

//ReadFileToIntArray takes a path and outputs the contents to an int array.
func ReadFileToIntArray(path string) ([]int, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}

		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// StringSliceToUintSlice takes a string slice and converts each entry to a uint64.
func StringSliceToUintSlice(strings []string) ([]uint64, error) {
	uints := make([]uint64, len(strings))
	for i, v := range strings {
		v, err := strconv.ParseUint(v, 2, 32)
		if err != nil {
			return nil, err
		}
		uints[i] = v
	}
	return uints, nil
}

//ReadFileToUIntArray takes a path and outputs the contents to an int array.
func ReadFileToUIntArray(path string) ([]uint64, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.ParseUint(strings.TrimSpace(line), 2, 32)
		if err != nil {
			return nil, err
		}

		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// ReadFileTo2DByteArray takes a path and outputs the contents to a [][]byte array.
func ReadFileTo2DByteArray(path string) ([][]byte, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)

		var runeSlice = []byte(line)
		if len(line) > 0 {
			lines = append(lines, runeSlice)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
