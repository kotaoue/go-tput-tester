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
	// sh.Command("tput", "setaf", strconv.Itoa(1)).Run()
	// out, err := exec.Command("set", "-c", "tput setaf 2").Output()
	// out, err := exec.Command("set", "-c", "tput setaf 2").Output()

	cmd := exec.Command("tput", "setaf", "2")
	/*
		stdin, _ := cmd.StdinPipe()
		io.WriteString(stdin, "hoge")
		stdin.Close()
	*/
	out, _ := cmd.Output()
	fmt.Printf("%s", out)

	fmt.Println("step 2")
}
