package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

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

func quizz(question []string) int {
	answer := bufio.NewReader(os.Stdin)
	fmt.Println("Question:", question[0])
	text, _ := answer.ReadString('\n')
	result := verifyAndCount(text, question[1])
	println(result)

	//TODO Count player points after good answer

	var points int

	if result {
		points++
	}
	return points
}

func verifyAndCount(userAnswer string, trueAnswer string) bool {
     if strings.TrimRight(userAnswer, "\n") == trueAnswer {
     	return true
	 } else {
	 	return false
	 }
}

func main() {
	file, err := os.Create("problem.csv", )
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}
	defer file.Close()


	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 5; i++ {
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
	questions := readFile(file.Name())

	var points int

	for _ , q := range questions {
		points += quizz(q)
	}
    fmt.Println("You have", points, "points out of 10")
}
