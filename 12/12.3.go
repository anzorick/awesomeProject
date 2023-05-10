package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("не удалось открыть файл", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 256)
	//if _, err := io.ReadFull(f, buf); err != nil {
	//	fmt.Println("не смогли прочитать последовательность байтов в файл", err)
	//	return
	//}
	//fmt.Println("\n", buf)

	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("не удалось открыть файл", err)
		return
	}
	fmt.Println(string(buf))

}
