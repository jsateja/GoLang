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

func main() {
    file, err := os.Create("problem.csv")
    if err != nil {
    	log.Fatalln("Couldn't create the csv file", err)
	}
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    var data [][]string

	for i := 0; i < 1000; i++ {
		firstNumber, secondNumber, result := generateRandomAddition()
		question := fmt.Sprintf("%d+%d", firstNumber, secondNumber)
		answer := strconv.Itoa(result)
		data := append(data, []string{question, answer})

	for _, value := range data {
		writer.Write(value)
	}

	readFile(file.Name())
	}
}

func readFile(csvFile string) {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

     reader := csv.NewReader(file)

     for {
     	question, err := reader.Read()
     		if err == io.EOF {
     			break
			}
			if err != nil {
				log.Fatal(err)
			}
			answer := bufio.NewReader(os.Stdin)
			fmt.Println("Question:", question[0])
			text, _ := answer.ReadString('\n')
			result := verifyAndCount(text, question[1])
			println(result)
		 //TODO Count player points after good question and move quizz logic to new function, also remove spaghetticode
		 }
	 file.Close()
	 }

func verifyAndCount(userAnswer string, trueAnswer string) bool {
     if strings.TrimRight(userAnswer, "\n") == trueAnswer {
     	return true
	 } else {
	 	return false
	 }
}

func generateRandomAddition() (int, int, int) {
	    min := 0
	    max := 10
		firstNumber := rand.Intn((max - min +1) + min)
		secondNumber := rand.Intn((max - min +1) + min)
		result := firstNumber + secondNumber
		return firstNumber, secondNumber, result
}