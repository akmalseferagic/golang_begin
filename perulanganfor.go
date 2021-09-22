package main

import (
	"fmt"
)

func main() {

	// for biasa
	//	for i := 1; i <= 100; i++ {
	//		fmt.Println("Aku mencintaimu dengan ikhlas :", i)
	//	}

	// for while
	//i := 1
	//for i <= 100 {
	//	fmt.Println("aku mengihlaskanmu :", i)
	//	i++

	// for untuk mengambil collection

	judul := "belajar untuk ikhlas"

	for index, letter := range judul {
		//for _,letter := range judul {

		fmt.Println("index :", index, "letter :", string(letter))
		//fmt.Println("letter :", string(letter))
	}

}
