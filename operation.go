package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GameMenu() {
	fmt.Println("Добро пожаловать в игру Виселица!\n 1.Начать игру\n 2.Выйти из игры")
	fmt.Print("Введите значение: ")
	input := getUserInput()
	switch input {
	case "1":
		fmt.Println("Вы начали игру!")
		fmt.Println(getWord())
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

func getWord() []string {
	data, err := os.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(data), "\r\n")
	return words
}
