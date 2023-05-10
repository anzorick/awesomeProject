package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Напишите программу, создающую текстовый файл только для чтения, и проверьте,
	//что в него нельзя записать данные.
	//Рекомендация
	//Для проверки создайте файл, установите режим только для чтения, закройте его,
	//а затем, открыв, попытайтесь прочесть из него данные.
	file, err := os.Create("test.txt")
	if err := os.Chmod("test.txt", 0444); err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer.WriteString("say hi")
	writer.Flush()

	file, err = os.Open("test.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
}
