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
	out, err := exec.Command("tput", "setaf", "2").Output()

	fmt.Printf("%s", out)
	fmt.Println("step 2")

	tput.Setaf(tput.Red)
	fmt.Println(err)

	fmt.Println("step 3")
}
