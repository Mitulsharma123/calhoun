package main

import (
	"fmt"
)

func main(){
		var input string = "oneTwoThree"
		answer := 1
		for _, ch := range input{
			min := 'A'
			max := 'Z'
			if ch >= min && ch <= max{
				answer++
			}
			/*str := string(ch)
			if strings.ToUpper(str) == str{
				answer++
			}*/
		}
		fmt.Println(answer)
}

