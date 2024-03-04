package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
)

func execCommand(quad, x, y string) string {
	cmd := exec.Command(quad, x, y)
	output, _ := cmd.Output()
	return string(output)
}

func isValidFunc(file, x, y string) bool {
	files := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	intX, _ := strconv.Atoi(x)
	intY, _ := strconv.Atoi(y)
	existFile := false
	if intX <= 0 || intY <= 0 {
		return false
	}
	for _, elem := range files {
		if elem == file {
			existFile = true
			break
		}
	}
	if !existFile {
		return existFile
	}
	return true
}

func findSameResults(quad, x, y, resultOfGiven string) []string {
	files := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	sameQuads := []string{quad[2:]}
	for i, elem := range files {
		if elem == quad {
			continue
		}
		output := execCommand(elem, x, y)
		if resultOfGiven == output {
			sameQuads = append(sameQuads, quads[i])
		}
	}
	sort.Strings(sameQuads)
	return sameQuads
}

func quadChecker(file, x, y string) {
	outputOfGiven := execCommand(file, x, y)
	sameQuads := findSameResults(file, x, y, outputOfGiven)
	for i, elem := range sameQuads {
		fmt.Printf("[%s] [%s] [%s]", elem, x, y)
		if i != len(sameQuads)-1 {
			fmt.Printf(" || ")
		}
	}
	fmt.Printf("\n")
}

func main() {
	args := os.Args
	if len(args[1:]) == 3 {
		if isValidFunc(args[1], args[2], args[3]) {
			quadChecker(args[1], args[2], args[3])
		} else {
			fmt.Printf("Not a quad function\n")
		}
	} else {
		fmt.Printf("Not a quad function\n")
	}
}
