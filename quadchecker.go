package main

import (
	"fmt"
	"os"
	"os/exec"
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
	sameQuads := []string{}
	for i, elem := range files {
		if elem == quad {
			continue
		}
		output := execCommand(elem, x, y)
		if resultOfGiven == output {
			sameQuads = append(sameQuads, quads[i])
		}
	}
	return sameQuads
}

func quadChecker(file, x, y string) {
	outputOfGiven := execCommand(file, x, y)
	sameQuads := findSameResults(file, x, y, outputOfGiven)
	fmt.Printf("[%s] [%s] [%s]", file[2:], x, y)
	for _, elem := range sameQuads {
		fmt.Printf(" || [%s] [%s] [%s]", elem, x, y)
	}
	fmt.Printf("\n")
}

func main() {
	args := os.Args
	if isValidFunc(args[1], args[2], args[3]) {
		quadChecker(args[1], args[2], args[3])
	} else {
		fmt.Printf("Not a quad function\n")
	}
}
