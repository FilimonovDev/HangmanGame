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
	input := getUserInput("Введите значение:\n 1.Начать игру\n 2.Выйти из игры", validationMenu)
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

func getUserInput(promt string, validationFunc func(string) bool) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(promt)
		scanner.Scan()
		text := scanner.Text()
		if validationFunc(text) {
			return text
		}
		fmt.Println("Вы ввели неверное значение, попробуйте еще раз")
	}
}

func validationMenu(choise string) bool {
	return choise == "1" || choise == "2"
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
