package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kshiva1126/goybox/cmd/mkpswd/controller"
)

var chars charFlags

func main() {
	var (
		nc = flag.Int("nchar", 8, "The number of characters")
		np = flag.Int("npass", 1, "The number of passwords")
	)
	flag.Var(&chars, "char", charFlagsUsage)
	flag.Parse()
	if len(chars) == 0 {
		chars = append(chars, charDefaultValue)
	}

	outputs, err := controller.CreatePassword(chars, nc, np)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(2)
	}

	fmt.Println(outputs)
}
