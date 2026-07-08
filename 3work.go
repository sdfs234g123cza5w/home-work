package main

import "fmt"

func main() {
  var name string
  var weight float64
  var distance float64
  var things int



BaseRate := 5.50
TaxRate := 0.12
DistanceRate := 2.0
FragileFee := 0.2
fmt.Print("Введите имя: ")
fmt.Scan(&name)
fmt.Println("Введите вес товара: ")
fmt.Scan(&weight)
fmt.Println("Введите дистанцию: ")
fmt.Scan(&distance)
fmt.Println("Введите количество хрупких товаров: ")
fmt.Scan(&things)

baseCost := (weight * BaseRate) * (1 + FragileFee * float64(things)) + (distance * DistanceRate)
totalCost := baseCost * (1 + TaxRate)

fmt.Println("\n--- Отчет о доставке ---")
fmt.Println("Отправитель:", name)
fmt.Printf("Итоговая стоимость: %.2f\n", totalCost)
}