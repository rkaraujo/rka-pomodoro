package main

import (
	"fmt"
	"strings"
	"time"
)

const Pomodoro = 25 * time.Minute

func main() {

main_loop:
	for {
		end := time.Now().Add(Pomodoro)
		fmt.Printf("Pomodoro will end at: %v\n", end.Format(time.RFC1123))

	timer_loop:
		for {
			now := time.Now()

			fmt.Printf("\rRemaining %s     ", end.Sub(now).Round(time.Second))

			if now.Equal(end) || now.After(end) {
				fmt.Println("\nPomodoro!!!")
				break timer_loop
			}

			time.Sleep(1 * time.Second)
		}

		if !confirm("Start another one? (y/n): ") {
			break main_loop
		}
	}
}

func confirm(msg string) bool {
	var input string

	fmt.Println(msg)
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	input = strings.ToLower(input)

	if input == "y" || input == "yes" {
		return true
	}
	return false
}
