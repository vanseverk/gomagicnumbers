package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/magicnumbers", magicnumbers)

	http.ListenAndServe(":8090", nil)
}

func magicnumbers(w http.ResponseWriter, req *http.Request) {
	magicNumber := rand.Intn(100)

	magicOne, magicTwo, magicThree := getMagicNumbers(magicNumber)

	fmt.Fprintf(w, "These are three magic numbers: %d, %d, and %d\n", magicOne, magicTwo, magicThree)
}

func getMagicNumbers(baseNumber int) (int, int, int) {
	magicNumbersChannel := make(chan int)

	go generateMagicNumber("one", baseNumber, generateMagicNumberOne, magicNumbersChannel)
	go generateMagicNumber("two", baseNumber, generateMagicNumberTwo, magicNumbersChannel)
	go generateMagicNumber("three", baseNumber, generateMagicNumberThree, magicNumbersChannel)

	return <-magicNumbersChannel, <-magicNumbersChannel, <-magicNumbersChannel
}

func generateMagicNumber(numberName string, randomNumber int, generatorFunction func(randomNumber int) int, resultChannel chan int) {
	result := generatorFunction(randomNumber)

	fmt.Println("Generated number ", numberName, " result is ", result)

	resultChannel <- result
}

func generateMagicNumberOne(randomNumber int) int {
	return randomNumber
}

func generateMagicNumberTwo(randomNumber int) int {
	numbers := []int{}

	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	return numbers[randomNumber%10]
}

type MagicNumbersHolder struct {
	magicNumber int
}

type MagicNumbersCalculator interface {
	doCalculation() int
}

func (mnh MagicNumbersHolder) doCalculation() int {
	return mnh.magicNumber * 5
}

func generateMagicNumberThree(randomNumber int) int {
	var magicNumberCalculator MagicNumbersCalculator = MagicNumbersHolder{randomNumber}

	return magicNumberCalculator.doCalculation()
}
