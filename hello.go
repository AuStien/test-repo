package main

import "fmt"

type Hello struct {
	Name string
}

func (h Hello) Say() {
	fmt.Printf("Hello, %s!\n", h.Name)
}
