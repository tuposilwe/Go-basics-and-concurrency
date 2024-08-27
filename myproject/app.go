package main

import (
	"errors"
	"fmt"
	"math"
)

func main(){
 result, err := sqr(6)

 if err != nil{
	fmt.Println(err)
 }else{
	fmt.Println(result)
 }

}

func sqr(x float64)(float64,error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbaers")
	}

	return math.Sqrt(x),nil
}