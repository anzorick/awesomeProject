package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("не удалось создать файл", err)
		return
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("введите число от 1 до 100")
	file.WriteString("введите число от 1 до 100 \n")
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			file.WriteString(fmt.Sprintf("Введенно число %d \n", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("число должно быть в диапазоне от 1 до 100")
				file.WriteString("число должно быть в диапазоне от 1 до 100 \n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("вы угадали число")
			file.WriteString("вы угадали число \n")
			return
		} else if answer < n {
			fmt.Println("загаданное число больше")
			file.WriteString("загаданное число больше \n")
		} else {
			fmt.Println("загаданное число меньше")
			file.WriteString("загаданное число меньше \n")
		}
	}
}
