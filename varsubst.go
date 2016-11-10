package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

var envVarRegexp = regexp.MustCompile(`\$\{[A-Z_0-9]+\}`)

func Scan(reader io.Reader) string {
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		return ""
	}
	return envVarRegexp.ReplaceAllStringFunc(string(input), func(placeholder string) string {
		name := placeholder[2 : len(placeholder)-1] // drop ${ and }
		if varvalue := os.Getenv(name); varvalue != "" {
			return varvalue
		}
		return placeholder
	})
}

func main() {
	fmt.Print(Scan(bufio.NewReader(os.Stdin)))
}
