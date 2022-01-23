package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kotaoue/go-tput"
)

var (
	f = flag.String("fg", "Reading", "test target function group name")
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
	case "Reading":
		c, err := tput.Cols()
		if err != nil {
			return err
		}
		fmt.Printf("cols: %d\n", c)
		return tput.HR()
	case "Text":
		_, err := tput.Setaf(tput.Red)
		if err != nil {
			return err
		}
		fmt.Println("color is red")

		_, err = tput.Smul()
		if err != nil {
			return err
		}
		fmt.Println("set underline")
		tput.Sgr0()

		var opt []*tput.Option
		opt = append(opt, &tput.Option{Attribute: tput.TextColor, Color: tput.Cyan})
		opt = append(opt, &tput.Option{Attribute: tput.UnderLine})
		opt = append(opt, &tput.Option{Attribute: tput.BoldText})
		tput.Printf(opt, "%s\n", "tput.Printf")

		return nil
	}
	return nil
}
