package main

import (
	"fmt"
)

func main(){
     a := []int{4,3,7,1,2}
  
	a = append(a, 13)

	fmt.Println(a)
}