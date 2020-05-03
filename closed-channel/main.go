package main

import "fmt"

func main() {
	chann := make(chan string)

	go runChannel("Salam", chann)

	for msg := range chann {
		fmt.Println(msg)
	}

}

func runChannel(text string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- text
	}

	// menutup chanel
	close(c)
}
