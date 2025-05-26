package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	quiz := flag.String("Quiz file","problems.csv","source file with the questions")
	rightAnswers := 0
	wrongAnswers := 0
	var solution int

	// read the CSV file 
	file, err := os.Open(*quiz)
	if err != nil{
		fmt.Println("The file failed to read.")
	}
	defer file.Close()

	// store file in "records" (list of lists of strings)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil{
		fmt.Println("Error reading file.")
	}

	fmt.Println("Do the Math, come on, it's gonna be fun:")
	for _, eachrecord := range records {
		// priklad
		fmt.Printf("%v: \n", eachrecord[0])
		goodSolution, err := strconv.Atoi(eachrecord[1])
		if err != nil {
			panic(err)
		}
		// Tadyk se zeptat na user input
		fmt.Scan(&solution)
		if solution == goodSolution {
			rightAnswers += 1
		} else {
			wrongAnswers += 1
		}
	}

	fmt.Printf("Number of right answers: %d\n", rightAnswers)
	fmt.Printf("Number of wrong answers: %d\n", wrongAnswers)
	fmt.Printf("Your score: %d out of %d \n", rightAnswers, rightAnswers+wrongAnswers)
	fmt.Println("Goog job, clever one!")

}