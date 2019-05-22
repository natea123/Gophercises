/*
Part 1
Create a program that will read in a quiz provided via a CSV file (more details below)
and will then give the quiz to a user keeping track of how many questions they get
right and how many they get incorrect. Regardless of whether the answer is correct
or wrong the next question should be asked immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user
should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question
and the second column in the same row is the answer to that question.

At the end of the quiz the program should output the total number of questions
correct and how many questions there were in total. Questions given invalid answers
are considered incorrect.

Libraries to use:
csv
os
flags
channels/go routines
time

*/

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	go timer()

	var problems = flag.String("file", "problems.csv", "help message for flagname")
	flag.Parse()

	correct := 0
	incorrect := 0

	// Open file
	file, err := os.Open(*problems)
	if err != nil {
		log.Fatal(err)
	}

	// Create reader object
	r := csv.NewReader(file)

	// Read entire object into records
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range records {

		fmt.Println(line[0])

		// Prompt for user input answer
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter answer: ")
		answer, _ := reader.ReadString('\n')

		// Check if user input equals answer and increment accordingly
		if strings.TrimRight(answer, "\n") == line[1] {
			correct += 1
		} else {
			incorrect += 1
		}
	}

	fmt.Println("RESULTS:")
	fmt.Println("Correct:" + strconv.Itoa(correct))
	fmt.Println("Incorrect:" + strconv.Itoa(incorrect))

}

func timer() {
	time.Sleep(30 * time.Second)
	os.Exit(1)
}
