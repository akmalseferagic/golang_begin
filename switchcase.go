package main

import "fmt"

func main() {
	number := 1

	switch number {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	default:
		fmt.Println("ulang")
	}

}
