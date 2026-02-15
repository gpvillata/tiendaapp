package helps

import "fmt"

func CleanScreen() {
	fmt.Print("\033[H\033[2J")
}
