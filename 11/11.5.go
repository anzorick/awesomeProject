package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Напишите программу, которая выведет все части строки
	//a10 10 20b 20 30c30 30 dd,
	//которые можно привести к числу в десятичном формате.
	s := "a10 10 20b 20 30c30 30 dd"
	space := " "
	for strings.Contains(s, space) {
		spaceIndex := strings.Index(s, space)
		word := s[:spaceIndex]
		i, err := strconv.Atoi(word)
		if err == nil {
			fmt.Println(i)
		}
		s = s[spaceIndex+1:]
		s = strings.Trim(s, space)
	}
	i, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println(i)
	}
}
