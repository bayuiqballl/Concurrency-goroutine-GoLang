package main

import (
	"fmt"
	"time"
)

func printAngka() {
	for i := 0; i < 5; i++ {
		// waktu menunggu
		time.Sleep(200 * time.Millisecond)
		fmt.Println(i)
	}
}

func printText() {
	for i := 0; i < 5; i++ {
		// waktu menunggu
		time.Sleep(500 * time.Millisecond)
		fmt.Println("text", i)
	}
}

func main() {
	start := time.Now()
	go printAngka()
	go printText()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println(time.Since(start))

}

// // menggunakan go routine
//  > printAngka .. .. .. ..
// > printText     ..  ..  ..  ..  ..
// > Main ----------- 3s
