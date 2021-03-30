package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main()	{
	filename, duration := initFlags()
	workingDir, _ := os.Getwd()
	questions := parseCSV(workingDir + "\\data\\"+ filename + ".csv")
	startGame(questions, duration)
}

func startGame(questions map[string]string, duration int) {
	currentQuestion, score := 1, 0
	
	timer := time.NewTimer(time.Duration(duration) * time.Second)
	for key, val := range questions {
		// this many not work in unix terminals
		fmt.Printf("\033[H\033[2JProblem #%v: %v = ", currentQuestion, key)
		answerCh := make(chan string)
		go parseUserInput(answerCh)

		select {
		case <-timer.C:
			fmt.Printf("\nTime has passed!\nYou scored %v/%v", score, len(questions))
			return
		case answer := <-answerCh:
			if (answer == val) {
				score++
			}
	
			currentQuestion++
		}
	}
	fmt.Printf("\nYou scored %v/%v", score, len(questions))
}

func parseUserInput(ch chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	ch<- strings.ToLower(strings.TrimSpace(text))
}

func initFlags() (string, int) {
	csvHelpMessage := "The filename of the csv under the data folder, extension excluded"
	durationHelpMessage := "The total duration of the game"

	var filename string
	var duration int

	flag.StringVar(&filename, "csv", "problems", csvHelpMessage)
	flag.IntVar(&duration, "dur", 30, durationHelpMessage)
	flag.Parse()

	return filename, duration
}

func parseCSV(filepath string) map[string]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	r := csv.NewReader(file)
	questions := make(map[string]string) 

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		
		questions[record[0]] = strings.ToLower(strings.TrimSpace(record[1]))
	}

	return questions
}

