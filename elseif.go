package main

import "fmt"

func main() {
	score := 20
	var grade string

	if score <= 20 {
		grade = "E"
	} else if score <= 50 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else if score <= 80 {
		grade = "B"
	} else if score <= 90 {
		grade = "A"
	} else {
		grade = "A+"
	}

	fmt.Println(grade)

}

//if
//if else
//else if
//if bersarang
