package main

import (
 "fmt"
)

type person struct {
	name string
	age int
}

func main(){
	p := person{name: "Rudiger", age: 22}
	fmt.Println(p)
}

