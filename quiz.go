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
)

func userPreference() {

}
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
		question := fmt.Sprintf("%d+%d", firstNumber, secondNumber)
		answer := strconv.Itoa(result)
		quizz_data := []string{question,answer}

		returnError := writer.Write(quizz_data)
		if returnError != nil {
			fmt.Println(returnError)
		}
	}
	writer.Flush()

	return file.Name()
}

func generateRandomAddition() (int, int, int) {
	min := 0
	max := 10
	firstNumber := rand.Intn((max - min +1) + min)
	secondNumber := rand.Intn((max - min +1) + min)
	result := firstNumber + secondNumber
	return firstNumber, secondNumber, result
}

//think about using bufio as a reader
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

func quiz(question []string) int {
	answer := bufio.NewReader(os.Stdin)
	fmt.Println("Question:", question[0])
	text, _ := answer.ReadString('\n')
	result := verifyAndCount(text, question[1])

	if result {
		return 1
	} else {
		return 0
	}
}

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

	flag.StringVar(&csvFileName,"file-name", "problem.csv", "name a csv file")
	flag.IntVar(&numberOfQuestions,"numb", 10, "define number of questions")

	flag.Parse()

	fileName := createFile(csvFileName, numberOfQuestions)

	questions := readFile(fileName)

	var points int

	for _ , q := range questions {
		points += quiz(q)
	}
    fmt.Println("You have", points, "points out of", numberOfQuestions)
}
