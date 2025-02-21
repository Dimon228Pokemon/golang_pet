package main

//go mod init pet_projects
import (
	"fmt"
	"pet_projects/perform"
)

func main() {

	fmt.Println("Дарова! Я тут калькулятор намутил, проверь потыкай")
	flag := true
	var history []string
	for flag {
		fmt.Println("Нужна ли инструкция по применению?(да/нет)")
		var choice string
		fmt.Scan(&choice)
		if choice == "да" {
			fmt.Println("Давай я тебе дам инструкцию что можно делать:")
			fmt.Println("Инструкция по применению:")
			fmt.Println("*1 - сложить два числа")
			fmt.Println("*2 - вычесть из первого числа второе")
			fmt.Println("*3 - умножить")
			fmt.Println("*4 - поделить")
		}
		fmt.Println("введи два числа:")
		var a float64
		var b float64
		fmt.Scan(&a, &b)
		fmt.Println("введи код операции:")
		var kod int
		fmt.Scan(&kod)
		firstNumber := perform.Number{Value: a, Is_zero: a == 0}
		secondNumber := perform.Number{Value: b, Is_zero: b == 0}
		switch kod {
		case 1:
			result := perform.Addition(firstNumber, secondNumber)
			fmt.Printf("Результат:%f\n", result)
			history = append(history, "+")
		case 2:
			result := perform.Subtraction(firstNumber, secondNumber)
			fmt.Printf("Результат:%f\n", result)
			history = append(history, "-")
		case 3:
			result := perform.Multiplication(firstNumber, secondNumber)
			fmt.Printf("Результат%f\n", result)
			history = append(history, "*")
		case 4:
			result, err := perform.Division(firstNumber, secondNumber)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Результат:%f\n", result)
				history = append(history, "/")
			}
		default:
			fmt.Printf("Я без понятие что это за код\n")
		}
		fmt.Println("Хотите завершить работу?(да/нет)")
		var cont string
		fmt.Scan(&cont)
		if cont == "да" {
			flag = false
			fmt.Print(history)
		}
	}
}
