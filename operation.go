package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
)

func gameMenu() {
	fmt.Println("Добро пожаловать в игру Виселица!\n 1.Начать игру\n 2.Выйти из игры")
	fmt.Print("Введите значение: ")
	input := getUserInput()
	switch input {
	case "1":
		fmt.Println("Вы начали игру!")
		fmt.Println(getRandomElemen())
	case "2":
		os.Exit(0)
	default:
		fmt.Println("Вы ввели неверное значение")
	}
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func getRandomElemen() string {
	data, err := os.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(data), "\r\n")
	randomIndex := rand.IntN(len(words))
	return words[randomIndex]
}
