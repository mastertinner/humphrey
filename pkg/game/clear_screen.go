package game

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

// clearScreen clears the terminal screen for different
// operating systems.
func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
