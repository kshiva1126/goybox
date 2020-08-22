package controller

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"

	"github.com/kshiva1126/goybox/cmd/mkpswd/service"
)

type charFlags []string

func (cf *charFlags) sort() error {
	// Check for expected value.
	err := service.AllowValue(*cf)
	if err != nil {
		return err
	}

	cf2 := new(charFlags)
	for _, c := range []string{"l", "u", "n", "s", "c"} {
		bool, err := service.Contains(c, []string(*cf))
		if err != nil {
			return err
		}

		if bool {
			*cf2 = append(*cf2, c)
		}
	}

	*cf = *cf2

	return nil
}

var passwordLetters = map[string]string{
	"l": "abcdefghijklmnopqrstuvwxyz",
	"u": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"n": "0123456789",
	"s": "!\"#$%'()*,./:;<=>?@[]^_`{|}~",
	"c": "loIO01\"',.:;^_`|~",
}

// CreatePassword returns the generated password.
func CreatePassword(chars []string, nc *int, np *int) ([]string, error) {
	cf := charFlags(chars)
	err := cf.sort()
	if err != nil {
		return nil, err
	}

	letters := concatPasswordLetters(cf)

	outputs, err := outputsPasswords(letters, nc, np)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return outputs, nil
}

func concatPasswordLetters(cf charFlags) string {
	var letters string
	for _, c := range cf {
		if c == "c" {
			pl, _ := passwordLetters[c]
			for _, excludeLetterRune := range pl {
				loc := strings.IndexRune(letters, excludeLetterRune)
				if loc != -1 {
					letters = letters[:loc] + letters[loc+1:]
				}
			}
			continue
		}

		pl, _ := passwordLetters[c]
		letters += pl
	}
	return letters
}

func outputsPasswords(letters string, nc *int, np *int) ([]string, error) {
	var output string
	var outputs []string
	buf := make([]byte, *nc)
	for i := 0; i < *np; i++ {
		if _, err := rand.Read(buf); err != nil {
			return outputs, err
		}
		for _, v := range buf {
			output += string(letters[int(v)%len(letters)])
		}
		outputs = append(outputs, output)
		output = ""
	}

	return outputs, nil
}
