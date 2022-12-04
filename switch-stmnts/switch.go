package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("The number ", i, " is ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	default:
		fmt.Println("Ponni Nadhi da vibe karo!!")
	}
	//It takes the day from the local machine's time&date !! Cool
	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday:
		fmt.Println("It's the weekday")
	case time.Friday, time.Sunday:
		fmt.Println("It's weekend ley")
		// default:
		// 	fmt.Println("It's a weekday")
	}

	//Similar to above switch, system time is taken here
	current_time := time.Now()
	switch {
	case current_time.Hour() < 12:
		fmt.Println("It's before noon")
	case current_time.Hour() > 12:
		fmt.Println("It's past noon")

	}
}
