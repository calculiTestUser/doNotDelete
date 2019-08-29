package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func main() {
	fmt.Println("main")
	f1()
}

