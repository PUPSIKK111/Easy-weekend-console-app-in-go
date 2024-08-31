package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Получаем текущую дату и время
	//Get wether and time now
	today := time.Now()

	// Определяем день недели
	// Determine the day of the week
	//Дни недели можно писать на русском и вобще на любом языке!
	switch today.Weekday() {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	case time.Saturday:
		fmt.Println("Saturday")
	case time.Sunday:
		fmt.Println("Sunday")
	default:
		log.Panic("Incorrect day")
	}
}
