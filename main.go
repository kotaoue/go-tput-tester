package main

import (
	"fmt"
	"os/exec"

	"github.com/kotaoue/go-tput"
)

func main() {
	// Cols
	cols, _ := tput.Cols()
	fmt.Printf("cols: %d\n", cols)

	tput.HR()

	fmt.Println("step 1")
	// out, err := exec.Command("set", "-c", "tput setaf 2").Output()
	err := exec.Command("sh", "-c", "tput setaf 2").Run()

	fmt.Println(err)

	fmt.Println("step 2")
	fmt.Println("step 3")
}
