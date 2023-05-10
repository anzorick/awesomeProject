package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Напишите программу, которая на вход получала бы строку, введённую пользователем,
	//а в файл писала № строки, дату и сообщение в формате:
	//2020-02-10 15:00:00 продам гараж.
	//При вводе слова exit программа завершает работу.
	file, err := os.Create("log1.txt")
	if err != nil {
		fmt.Println("не удалось создать файл", err)
		return
	}
	defer file.Close()

	count := 0
	for {
		strInp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		strInp = strings.Trim(strInp, "\n")

		if strInp == "exit" {
			break
		}
		count++

		dataTime := time.Now()
		message := strconv.Itoa(count) + " " + dataTime.Format("2006-01-02 15:04:05") + " " + strInp
		file.WriteString(message)
		// НЕ ВЫХОДИТ ИЗ ЦИКЛА
	}

}
