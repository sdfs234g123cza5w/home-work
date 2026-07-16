package main

import "fmt"

func main() {
	fmt.Println(5 == 5)

	fmt.Println(10 != 3)

	fmt.Println(7 > 12)

	fmt.Println(15 < 20)

	fmt.Println(8 >= 8)

	fmt.Println(6 <= 4)

	fmt.Println((10 > 5) && (3 < 1))

	fmt.Println((10 > 5) || (3 < 1))

	fmt.Println(!(5 == 5))

	fmt.Println(!(7 < 3))

	fmt.Println(true && false)

	fmt.Println(false || false)

	fmt.Println(true || false)

	fmt.Println((4 + 6 == 10) && (9 > 2))

	fmt.Println((12 / 3 == 4) || (8 < 5))

	age := 20
	hasTicket := true

	canEnter := age >= 18 && hasTicket

	fmt.Println("Can enter:", canEnter)

	isLoggedIn := true
  
	isAdmin := false

	hasAccess := (isLoggedIn && isAdmin) || (isLoggedIn && !isAdmin)

	fmt.Println("Has access:", hasAccess)
}
