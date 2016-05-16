package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
  <letter>   = A-Z
  <digit>    = 0-9
  <special>  = "_"
  <varname>  = <letter> {<letter> | <digit> | <special>}
  <variable> = <dollar> <lparen> <varname> <rparen>
  <dollar>   = "$"
  <lparen>   = "("
  <rparen>   = ")"
*/

var eof = rune(0)

func isLetter(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z')
}
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
func isSpecial(ch rune) bool {
	return ch == '_'
}
func isDollar(ch rune) bool {
	return ch == '$'
}
func isLparen(ch rune) bool {
	return ch == '{'
}
func isRparen(ch rune) bool {
	return ch == '}'
}

func read(r *bufio.Reader) rune {
	ch, _, err := r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func unread(r *bufio.Reader) { _ = r.UnreadRune() }

func varname(r *bufio.Reader) string {
	var literal, varname string
	var ch rune

	ch = read(r)
	literal += string(ch)
	if !isLetter(ch) {
		return literal
	} else {
		varname += string(ch)
	}

	for {
		ch = read(r)
		literal += string(ch)
		if isLetter(ch) || isDigit(ch) || isSpecial(ch) {
			varname += string(ch)
		} else {
			if isRparen(ch) {
				unread(r)
				return varname
			} else {
				return literal
			}
		}
	}
}

func variable(r *bufio.Reader) string {
	var literal, varvalue string
	var ch rune

	// read "$("
	ch = read(r)
	literal += string(ch)
	if !isDollar(ch) {
		return literal
	}
	ch = read(r)
	literal += string(ch)
	if !isLparen(ch) {
		return literal
	}

	// read <varname>
	name := varname(r)
	literal += name

	// read env var with this name
	if varvalue = os.Getenv(name); varvalue == "" {
		return literal
	}
	// read ")"
	ch = read(r)
	literal += string(ch)
	if !isRparen(ch) {
		return literal
	}
	return varvalue
}

func Scan(r *bufio.Reader) string {
	var result string
	for {
		if ch := read(r); ch != eof {
			if isDollar(ch) {
				unread(r)
				result += variable(r)
			} else {
				result += string(ch)
			}
		} else {
			break
		}
	}
	return result
}

func main() {
	r := bufio.NewReader(os.Stdin)
	result := Scan(r)
	fmt.Print(result)
}
