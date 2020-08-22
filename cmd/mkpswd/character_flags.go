package main

import "fmt"

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
