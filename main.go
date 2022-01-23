package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kotaoue/go-tput"
)

var (
	f = flag.String("func", "", "func name to run")
	a = flag.String("arg", "", "argument of sending to function")
)

func init() {
	flag.Parse()
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	switch *f {
	case "Cols":
		c, err := tput.Cols()
		if err != nil {
			return err
		}
		fmt.Printf("Cols: %d", c)
		return nil
	case "HR":
		return tput.HR()
	case "Setaf":
		i, err := strconv.Atoi(*a)
		if err != nil {
			return err
		}
		b, err := tput.Setaf(i)
		if err != nil {
			return err
		}
		fmt.Printf("byte: %s color: %d", b, i)
		return nil
	}
	return nil
}
