package main

import (
	"bufio"
	"fmt"
	"os"
)

func GameMenu() {
	fmt.Println("Добро пожаловать в игру Виселица!\n 1.Начать игру\n 2.Выйти из игры")
	fmt.Print("Введите значение: ")
	input := getUserInput()
	if input == "1" {
		fmt.Println("Вы начали игру!")
	} else if input == "2" {
		os.Exit(0)
	} else {
		fmt.Println("Вы ввели неверное значение")
	}

}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}
