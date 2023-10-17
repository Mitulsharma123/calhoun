package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main(){
	csvfileName := flag.String("csv", "problems.csv", "a csv file in QnA format")
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
	fmt.Println(lines)
			
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

