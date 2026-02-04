package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"unicode/utf8"
)

func gameMenu() {
	input := getUserInput("Нажмите 1, чтобы начать игру, либо 2 чтобы выйти:\n 1.Начать игру\n 2.Выйти из игры", validationMenu)
	switch input {
	case "1":
		fmt.Println("Вы начали игру!")
		gameSession()
	case "2":
		os.Exit(0)
	default:
		fmt.Println("Вы ввели неверное значение")
	}
}

func gameSession() {
	hiddenWord := getRandomElemen()
	enteredLetters := make(map[rune]bool)
	attempts := 6

	for attempts > 0 {
		fmt.Println(hiddenWord)
		fmt.Println(displayWord(hiddenWord, enteredLetters))
		fmt.Printf("У вас осталось %d попыток\n", attempts)
		input := strings.ToLower(getUserInput("Введите букву:", validationLetter))
		letter := []rune(input)[0]
		if enteredLetters[letter] {
			fmt.Println("Вы уже вводили эту букву")
		} else {
			enteredLetters[letter] = true
		}

		if strings.ContainsRune(hiddenWord, letter) {
			fmt.Println("Такая буква есть в слове!")
			if getGuessedWord(hiddenWord, enteredLetters) {
				fmt.Println("Поздравляю, вы отгадали слово!")
				fmt.Println("Загаданное слово: ", hiddenWord)
				return
			}
		} else {
			fmt.Println("Такой буквы нет в слове")
			attempts--
		}

	}
	if attempts == 0 {
		fmt.Println("Вы проиграли!")
		fmt.Println("Загаданное слово :", hiddenWord)
	}
}

func getUserInput(promt string, validationFunc func(string) bool) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(promt)
		scanner.Scan()
		text := scanner.Text()
		if validationFunc(text) && utf8.RuneCountInString(text) == 1 {
			return text
		}
		fmt.Println("Вы ввели неверное значение, попробуйте еще раз")
	}
}

func validationMenu(choise string) bool {
	return choise == "1" || choise == "2"
}

func validationLetter(letter string) bool {
	return (letter >= "а" && letter <= "я") || (letter >= "А" && letter <= "Я")
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

func displayWord(word string, letters map[rune]bool) string {
	var result []rune
	for _, value := range word {
		if letters[value] {
			result = append(result, value)
		} else {
			result = append(result, '_')
		}
		result = append(result, ' ')
	}
	return string(result)

}

func getGuessedWord(word string, letters map[rune]bool) bool {
	for _, value := range word {
		if !letters[value] {
			return false
		}
	}
	return true
}
