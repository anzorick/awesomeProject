package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "10"
	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

}
