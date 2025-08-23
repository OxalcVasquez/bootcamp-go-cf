package main

import "fmt"

type User interface {
	PrintName()
}

type Professor struct {
	Id   int
	Name string
}

func (p Professor) PrintName() {
	fmt.Println(p.Name)
}

func PrintStructName(u interface{}) {
	fmt.Println(u)
}

func main() {
	p := Professor{Id: 1, Name: "Jordan"}
	p.PrintName()
	PrintStructName(p)
}
