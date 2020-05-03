package main

import (
	"fmt"
	"time"
)

// mengirim data antara satu goroutine ke goroutine lainnya kita bisa menggunakan channel di go lang

var itemsChannel = make(chan string)
var cleanedChannel = make(chan string)

func main() {
	items := [7]string{"batu", "harta", "kerang", "harta", "batu", "kerang"}
	/*
		Kita mau menyelam mencari harta karun,
		 terselip diantara banyak barang
		 tugasnya adalah membagi tugas yaitu
		 1. Penyelam , 2. Bersihkan, 3. Disimpan
		 membuat channel tsb
	*/

	go menyelam(items)
	go membersihkan()
	go menyimpan()

	time.Sleep(500 * time.Millisecond)
}

func menyelam(items [7]string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("Berhasil mendapatkan " + item)
			itemsChannel <- item
		}
	}
}

func membersihkan() {
	for rawItem := range itemsChannel {
		fmt.Println("Berhasil membersihkan " + rawItem)
		cleanedChannel <- "HartaBersih"
	}
}

func menyimpan() {
	for cleanedItem := range cleanedChannel {
		fmt.Println("Berhasil menyimpan " + cleanedItem)
	}
}
