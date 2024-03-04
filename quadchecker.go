package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getPreviousOutput() string {
	input, _ := ioutil.ReadAll(os.Stdin)
	str := string(input)
	return str
}

func execCommand(quad, x, y string) string {
	cmd := exec.Command(quad, x, y)
	output, _ := cmd.Output()
	return string(output)
}

func getDimensions(str string) (int, int, bool) {
	err := false
	lines := strings.Split(str, "\n")
	y := len(lines) - 1
	if y == 0 {
		return 0, 0, true
	}
	for _, char := range lines[0] {
		if char == ' ' {
			return 0, 0, true
		}
	}
	for i := 0; i < len(lines)-1; i++ {
		if len(lines[0]) != len(lines[i]) {
			return 0, 0, true
		}
	}
	x := len(lines[0])
	return x, y, err
}

func findSameResults(x int, y int, resultOfGiven string) []string {
	files := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	sameQuads := []string{}
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)
	for _, elem := range files {
		output := execCommand(elem, strX, strY)
		if resultOfGiven == output {
			sameQuads = append(sameQuads, elem[2:])
		}
	}
	return sameQuads
}

func quadChecker(previousOutput string) {
	x, y, err := getDimensions(previousOutput)
	if !err {
		sameQuads := findSameResults(x, y, previousOutput)
		for i, elem := range sameQuads {
			fmt.Printf("[%s] [%d] [%d]", elem, x, y)
			if i != len(sameQuads)-1 {
				fmt.Printf(" || ")
			}
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("Not a quad function\n")
	}
}

func main() {
	previousOutput := getPreviousOutput()
	if len(previousOutput) > 0 {
		quadChecker(previousOutput)
	} else {
		fmt.Printf("Not a quad function\n")
	}
}
