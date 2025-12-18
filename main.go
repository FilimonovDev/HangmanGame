package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Добро пожаловать в игру Виселица!\n 1.Начать игру\n 2.Выйти из игры")
	fmt.Print("Введите значение: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Читает ВСЮ строку до \n
	text := scanner.Text()
	if text == "1" {
		fmt.Println("Вы начали игру!")
	} else if text == "2" {
		os.Exit(0)
	} else {
		fmt.Println("Вы ввели неверное значение")
	}

}
