package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	csvfileName := flag.String("csv", "problems.csv", "a csv file in QnA format")
	
	timeLimit := flag.Int("limit", 5, "time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvfileName)
	if err != nil{
		exit(fmt.Sprintf("failed to open the csv file %s\n", *csvfileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err!=nil{
		exit("Failed to parse the provided the csv file")
	}
	problems := parseLines(lines)	
	//fmt.Println(problems)	//result in form of struct variable	

	count := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//<-timer.C

	for i, p := range problems{
		select{
		case <-timer.C:
			fmt.Printf("You've Scored %d out of %d\n", count, len(problems))
			return	//not using break, as it will come out of entire loop
		default: 
			fmt.Printf("Problem #%d: %s =? \n", i+1, p.q)
			var answer string
			fmt.Scanf("%s \n", &answer)
			if answer == p.a{
				//fmt.Println("Correct Answer !!!")
				count++
			}
		}			
	}
}
//parseLines will take input as 2D lines array and return slice of problems
func parseLines(lines [][]string) []problem{
	ret := make([]problem, len(lines))
	for i, line := range lines{
		ret[i] = problem{
			q: strings.TrimSpace(line[0]), //column 1
			a: strings.TrimSpace(line[1]), //column 2
		}
	}
	return ret

}

type problem struct {
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

