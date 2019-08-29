package main

import (
	"fmt"
	"strings"
)

func foobar() {
	fmt.Println(strings.ContainsAny("shwetatest", "i & u & a"))
	fmt.Println("abc" < "def")
}
