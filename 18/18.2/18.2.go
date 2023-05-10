package main

var c int

func changeC(b int) {
	c = b
}

func main() {
	for i := 0; i < 10; i++ {
		changeC(i)
	}
}
