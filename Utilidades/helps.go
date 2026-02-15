package helps

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// CleanScreen limpia la pantalla de la terminal de forma multiplataforma
func CleanScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CleanScreenANSI limpia la pantalla usando códigos ANSI (Linux/Mac)
func CleanScreenANSI() {
	fmt.Print("\033[H\033[2J")
}
