package main

import (
	"fmt"
	"strings"
)

func main(){
		var input string = "oneTwoThree"
		answer := 1
		for _, ch := range input{
			str := string(ch)
			if strings.ToUpper(str) == str{
				answer++
			}
		}
		fmt.Printf("Number of words prsent in the string is %d\n", answer)
}

