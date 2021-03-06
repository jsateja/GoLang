package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// It creates csv file with simple questions
func createFile(csvFileName string, numberOfquestions int) string {
	file, err := os.Create(csvFileName)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < numberOfquestions; i++ {
		firstNumber, secondNumber, result := generateRandomAddition()
		question := fmt.Sprintf("%d + %d", firstNumber, secondNumber)
		answer := strconv.Itoa(result)
		quizz_data := []string{question, answer}

		returnError := writer.Write(quizz_data)
		if returnError != nil {
			fmt.Println(returnError)
		}
	}
	writer.Flush()

	return file.Name()
}

// Generates questions: simple addition
func generateRandomAddition() (int, int, int) {
	min := 0
	max := 10
	firstNumber := rand.Intn((max - min + 1) + min)
	secondNumber := rand.Intn((max - min + 1) + min)
	result := firstNumber + secondNumber
	return firstNumber, secondNumber, result
}

//Reads the csv file with questions
func readFile(csvFile string) [][]string {
	file, err := os.OpenFile(csvFile, os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	questions := [][]string{}

	for {
		question, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questions = append(questions, question)
	}
	return questions

}

// Start quiz game, read the player answer from the keyboard
func quiz(question []string) int {
	answer := bufio.NewReader(os.Stdin)
	fmt.Println("Question:", question[0])
	userAnswer, _ := answer.ReadString('\n')
	result := verifyAndCount(userAnswer, question[1])

	if result {
		return 1
	} else {
		return 0
	}
}

// Verify the answer provided by the player
func verifyAndCount(userAnswer string, trueAnswer string) bool {
	if strings.TrimRight(userAnswer, "\n") == trueAnswer {
		return true
	} else {
		return false
	}
}

func main() {
	var csvFileName string
	var numberOfQuestions int
	var quizTime int

	flag.StringVar(&csvFileName, "file-name", "problem.csv", "name a csv file, default: problem.csv")
	flag.IntVar(&numberOfQuestions, "numb", 10, "define number of questions, default: 10")
	flag.IntVar(&quizTime, "time", 30, "define time (in seconds )of the quiz, default: 30s")

	flag.Parse()

	fileName := createFile(csvFileName, numberOfQuestions)

	questions := readFile(fileName)

	var points int

	timer := time.NewTimer(time.Duration(quizTime) * time.Second)
	for _, q := range questions {
		select {
		case <-timer.C:
			fmt.Println("Sorry, timed out!")
			fmt.Println("You scored", points, "points out of", numberOfQuestions)
			return

		default:
			points += quiz(q)
		}
	}

	fmt.Println("You scored", points, "points out of", numberOfQuestions)
}
