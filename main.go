package main

import (
	"fmt"
	"strings"
)

func _1_1() {
	print("kek!!!")
}

func main() {
	done := map[string]bool{
		"1.1": true}

	hr := "\n====================\n"

	println("HELLO, THIS IS THE GO EXERCISES FOLDER")
	println("MADE BY DMITRY ROSIN, STUDENT OF KOSYGIN RSU, GROUP MPM-121")
	println("To see the help page, type 'help', to exit the program, type 'exit'")
	println(hr)

	for {
		print("What exercise do you wanna check: ")
		var line string
		_, err := fmt.Scan(&line)
		println(hr)

		if err != nil {
			panic(err)
		} else {
			switch {

			case strings.ToLower(line) == "help":
				{
					println("You can pop up this help page again by typing 'help' when you're choosing exercises.")
					println("You can exit the program by typing 'exit' when you're choosing exercises.")
					println("You can choose the exercise you wanna check by typing its number like it was in the book.")
					println("For example you can type '1.1' or '1.12'. You shouldn't add any zeros like '01.01'.")
					println("Also you should use a point '.' not a comma ',' to separate chapter and exercise number.")
					println("The input isn't case-sensetive. And wrong inputs won't crash the program.")
					println("Good luck!")
				}
			case strings.ToLower(line) == "exit":
				{
					println("Thank you for using the program. In hope for a good mark :)")
					return
				}
			case done[line]:
				{
					switch line {
					case "1.1":
						_1_1()
					}
				}

			default:
				{
					println("Sorry, there's no such an exercise or a command...")
				}
			}
		}

		println(hr)
	}
}
