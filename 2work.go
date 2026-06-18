package main

import "fmt"

func main() {

	age := 14
	fmt.Println(age)
	age = 15
	fmt.Println(age)

	height := 170
	fmt.Println(height)
	height_in_meters = 1.7

	isStudent := true
	fmt.Println(isStudent)

	temperature := 25
	    if temperature >= 15 {
        fmt.Println("Погода теплая")
    } else {
        fmt.Println("Погода холодная")
    }

	favoriteQuote := "Что разум человека может постигнуть и во что он может поверить, того он способен достичь."
	fmt.Println(favoriteQuote)

	PI := 3.14
	fmt.Println(PI)

	PI = "3.1415" //Go запрещает записывать текст в числовую переменную
	fmt.Println(PI)
}
