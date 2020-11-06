package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	plusPerator()
	fmt.Println("========================")
	bytesBufferWriteString()
	fmt.Println("========================")
	usingSprintf()
	fmt.Println("========================")
	usingJoin()
	fmt.Println("========================")
}
func plusPerator() {
	str1 := "hello,"
	str2 := "where do you want to go?\n"
	str3 := "I'm coming"
	result := str1 + str2 + str3 + "Where are you?"
	fmt.Println("+ operator concatenate of str1,str2,str3...: ", result)
}
func bytesBufferWriteString() {
	var b bytes.Buffer
	b.WriteString("hello Sir1,")
	b.WriteString("hello Sir2,")
	fmt.Println("bytesBufferWriteString2 b:  ", b.String())
	b.WriteString("hello Sir3")
	fmt.Println("bytesBufferWriteString3 b:  ", b.String())

	//Important point
	//bytes.Buffer.WriteString("hi Sir1,")//C.E. bcoz WriteString is method not function. function can call through package name only(except inbuild func like len...)
	//method can call only through variable ref of struct/interface NOT through struct/interface name directly
	var b2 bytes.Buffer
	b2.WriteString("hi Sir1,")
	fmt.Println("bytesBufferWriteString b2:  ", b2.String())
}

func usingSprintf() {
	str := fmt.Sprintf("s1:%s s2:%s s3:%s", "hello Sir1", "hello Sir2", "hello Sir3")
	fmt.Println("UsingSprintf :  ", str)
}

func usingJoin() {
	fruitsSlice := []string{"Apple", "Orange", "Mango"}
	resultStr1 := strings.Join(fruitsSlice, "****")
	fmt.Println("UsingSprintf resultStr1 :  ", resultStr1)

	str := strings.Join([]string{"hello Sir1", "hello Sir2", "hello Sir3"}, "--")
	fmt.Println("UsingSprintf :  ", str)
}
