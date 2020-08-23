package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kshiva1126/goybox/cmd/mkpswd"
)

type charFlags []string

var charDefaultValue string = "n"

func (cf *charFlags) String() string {
	return fmt.Sprintf("%v", *cf)
}

func (cf *charFlags) Set(value string) error {
	*cf = append(*cf, value)
	return nil
}

var charFlagsUsage = `Multiple selections are available (default: n)

l: Lowercase characters
u: Uppercase characters
n: Including number
s: Including symbolic characters
c: Not including confusing characters, like ` + "l o I O 0 1 \" ' , . : ; ^ _ ` | ~"

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

	outputs, err := mkpswd.CreatePassword(chars, nc, np)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(2)
	}

	fmt.Println(outputs)
}
