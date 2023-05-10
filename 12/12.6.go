package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("введите число от 1 до 100")
	b.WriteString("введите число от 1 до 100 \n")
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			b.WriteString(fmt.Sprintf("Введенно число %d \n", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("число должно быть в диапазоне от 1 до 100")
				b.WriteString("число должно быть в диапазоне от 1 до 100 \n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("вы угадали число")
			b.WriteString("вы угадали число \n")
			break
		} else if answer < n {
			fmt.Println("загаданное число больше")
			b.WriteString("загаданное число больше \n")
		} else {
			fmt.Println("загаданное число меньше")
			b.WriteString("загаданное число меньше \n")
		}
	}
	fileName := "log.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("сохраненный лог:")
	fmt.Println(string(resultBytes))
}
