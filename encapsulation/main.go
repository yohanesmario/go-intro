package main

import (
	"fmt"

	"github.com/yohanesmario/go-intro/encapsulation/capsule"
)

func echo() {
	fmt.Println("abcdef")
}

func main() {
	fmt.Println(capsule.GetString())

	c := capsule.Capsule{PublicContent: "abcd"}
	fmt.Println(c.SetContent("abc").GetContent())
	c.PublicContent = "test"
	fmt.Println(c.GetPublicContent())
}
