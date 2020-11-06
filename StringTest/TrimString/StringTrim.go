package main

import (
	"fmt"
	"strings"
)

var str1 = "!!Welcome to GeeksforGeeks !!"
var str2 = "@@This is the tutorial of Golang$$"

func main() {
	strTrim()
	fmt.Println("=======================")
	strTrimleft()
	fmt.Println("=======================")
	strTrimRight()
	fmt.Println("=======================")
	strTrimSpace()
	fmt.Println("=======================")
	strTrimPrefix()
	fmt.Println("=======================")
	strTrimSuffix()
	fmt.Println("=======================")
}
func strTrim() {
	fmt.Println("strTrim", strings.Trim(str1, "!"))
	fmt.Println("strTrim", strings.Trim(str2, "$@"))
}

func strTrimleft() {
	fmt.Println("strTrimleft: ", strings.TrimLeft(str1, "!"))
	fmt.Println("strTrimleft", strings.TrimLeft(str2, "$@"))
}

func strTrimRight() {
	fmt.Println("strTrimRight: ", strings.TrimRight(str1, "!"))
	fmt.Println("strTrimRight", strings.TrimRight(str2, "$@"))
}

func strTrimSpace() {
	fmt.Println("strTrimSpace: ", strings.TrimSpace(str1+"   "))
	fmt.Println("strTrimSpace", strings.TrimSpace(str2+"   "))
}

func strTrimPrefix() {
	fmt.Println("strTrimPrefix: ", strings.TrimPrefix(str1, "!!"))
	fmt.Println("strTrimPrefix", strings.TrimPrefix(str2, "@@"))
}

func strTrimSuffix() {
	fmt.Println("strTrimSuffix: ", strings.TrimSuffix(str1, "forGeeks !!"))
	fmt.Println("strTrimSuffix", strings.TrimSuffix(str2, "lang$$"))
}
func strTrimSuffix2(str1, str2 string, i int) {
	fmt.Println("strTrimSuffix: ", strings.TrimSuffix(str1, "forGeeks !!"))
	fmt.Println("strTrimSuffix", strings.TrimSuffix(str2, "lang$$"))
}
