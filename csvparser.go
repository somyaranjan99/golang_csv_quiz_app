package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csvfile := flag.String("csv", "problem.csv", "csv file path")
	timer := flag.Int("time", 30, "expired time to complete")
	flag.Parse()

	file, err := os.Open(*csvfile)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	csvData := quesionAndAnswer(data)
	expiredTimer := time.NewTimer(time.Duration(*timer) * time.Second)

	//fmt.Println(csvData)
	var right_answer int
	ansChan := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ansChan <- scanner.Text()

		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	for i, qAndA := range csvData {
		fmt.Printf("Qustion # %d:#  :%s=\n", i+1, qAndA.Question)

		select {
		case <-expiredTimer.C:
			fmt.Println("Time expired")
			return
		case answer := <-ansChan:
			if qAndA.Answer == answer {
				right_answer++
			}
		}
	}

	fmt.Printf("Total score :=%d", right_answer)
}

func quesionAndAnswer(data [][]string) []Problem {
	result := make([]Problem, len(data))
	for i, databuild := range data {
		result[i] = Problem{Question: databuild[0], Answer: databuild[1]}
	}
	return result
}
