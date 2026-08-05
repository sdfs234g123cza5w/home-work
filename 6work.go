package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {
        fmt.Println(i)
    }
	
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    number := 1
    for i := 1; i <= 10; i++ {
        fmt.Println(number, "*", i, "=", number*i)
    }

	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		if i % 3 == 0{
			fmt.Println(i)
		}
	}
	
	var number2 int
	fmt.Scan(&number2)

	counter := 0
	for number2 > 0{
		counter += 1
		number2 /= 10
	}
	fmt.Println(counter)

	text := "Developer"
	
	for _, value := range text {
         fmt.Printf("Value: %c\n", value)
    }

	balance := 3000
	var m int
	for{
		fmt.Println("1 - вывести текущий баланс")
		fmt.Println("2 - увеличить баланс на 500")
		fmt.Println("3 - уменьшить баланс на 200")
		fmt.Println("0 - Выход")
		fmt.Println("Введите число: ")
		fmt.Scan(&m)
		switch m{
    	  case 1:
         	fmt.Println(balance)
    	  case 2:
         	balance += 500
    	  case 3:
         	balance -= 200
		  case 0:
			break
}

	}



}
