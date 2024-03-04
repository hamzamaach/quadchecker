package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("./quadchecker")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
