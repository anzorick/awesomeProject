package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("введите имя")
	var username string
	fmt.Scan(&username)
	fmt.Println("введите пароль")
	var password string
	fmt.Scan(&password)
	fmt.Println("введите возраст")
	var age int
	fmt.Scan(&age)

	file, err := os.Create("creadentials.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("ваш логин %s \n", username))
	b.WriteString(fmt.Sprintf("ваш пароль %s \n", password))
	b.WriteString(fmt.Sprintf("ваш возраст %d \n", age))
	_, err = file.Write(b.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}
