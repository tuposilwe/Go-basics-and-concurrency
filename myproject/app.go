package main

import (
	"fmt"
)

func main(){
  vertices := make(map[string]int)

  vertices["triangle"] =2
  vertices["square"] = 3
  vertices["dodecagon"] = 12

  delete(vertices,"triangle")

  fmt.Println(vertices)
}