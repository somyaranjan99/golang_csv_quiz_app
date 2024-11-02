package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main_r() {
	csvFile := flag.String("csv", "problem.csv", "question and answer")
	timer := flag.Int("limit", 30, "Limit of timer")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Error occured while parsing %s\n", *csvFile))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error occured while reading %s\n", *csvFile))
	}
	allAnswer := questinAnswer(lines)
	timerLimit := time.NewTimer(time.Duration(*timer) * time.Second)
	correct := 0
	for i, qandn := range allAnswer {
		fmt.Printf("Problem# %d# %s= \n", i+1, qandn.question)
		ansCha := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansCha <- answer
		}()
		select {
		case <-timerLimit.C:
			fmt.Println("Time has been expired")
			return
		case answer := <-ansCha:
			if qandn.answer == answer {
				fmt.Println(answer)
				correct++
			}

		}

	}
	fmt.Println("you have given correct answer", correct)

}
func questinAnswer(linnes [][]string) []problem {
	result := make([]problem, len(linnes))
	for i, qandn := range linnes {
		result[i] = problem{question: qandn[0], answer: qandn[1]}
	}

	return result
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
