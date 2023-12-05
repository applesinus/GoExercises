package main

import (
	"GoExercises/ch1"
	"GoExercises/ch2"
	"GoExercises/ch3"
	"fmt"
	"strconv"
	"strings"
)

func enterArgs(amount int) []string {
	args := make([]string, 0)
	switch {
	// 0 means that there's undefined amount of arguments
	case amount == 0:
		{
			println("Enter args (to stop enter 'done'): ")
			var arg string
			for i := 0; true; i++ {
				fmt.Scan(&arg)
				if strings.ToLower(arg) == "done" {
					println()
					return args
				}
				args = append(args, arg)
			}
		}
	case amount > 0:
		{
			println("Enter " + strconv.Itoa(amount) + " arg: ")
			var arg string
			for i := 0; i < amount; i++ {
				fmt.Scan(&arg)
				args = append(args, arg)
			}
		}
	default:
		{
			println("ERROR! Amount of arguments is less than 0 in main.go/enterArgs.")
		}
	}
	println()
	return args
}

func main() {
	done := map[string]bool{
		"1.1": true,
		"1.2": true,
		"1.3": true,
		"1.4": true,
		"1.5": true,
		"1.6": true,
		"2.2": true,
		"2.3": true,
		"2.4": true,
		"2.5": true,
		"3.1": true,
	}

	hr := "\n====================\n"

	println("HELLO, THIS IS THE GO EXERCISES FOLDER")
	println("MADE BY DMITRY ROSIN, STUDENT OF KOSYGIN RSU, GROUP MPM-121")
	println("To see the help page, type 'help', to see the list of exercises, type 'list'.")
	println("To exit the program, type 'exit'")
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

			case strings.ToLower(line) == "help" || strings.ToLower(line) == "h":
				{
					println("You can pop up this HELP page again by typing 'help' or 'h' when you're choosing exercises.")
					println("You can EXIT the program by typing 'exit', 'e' or '-1' when you're choosing exercises.")
					println("You can see the list of avaliable exercises by  typing 'list' or 'l' when you're choosing exercises.")
					println("You can choose the EXERCISE you wanna check by typing its number like it was in the book.")
					println("For example you can type '1.1' or '1.12'. You shouldn't add any zeros like '01.01'.")
					println("Also you should use a point '.' not a comma ',' as a separator.")
					println("The input isn't case-sensetive. And wrong inputs won't crash the program.")
					println("Good luck!")
				}
			case strings.ToLower(line) == "exit" || strings.ToLower(line) == "e" || strings.ToLower(line) == "-1":
				{
					println("Thank you for using the program. In hope for a good mark :)")
					return
				}
			case strings.ToLower(line) == "list" || strings.ToLower(line) == "l":
				{
					println("Chatper 1: \n1.1\t1.2\t1.3\t1.4\t1.5\t1.6 \n")
					println("Chapter 2: \n2.2\t2.3\t2.4\t2.5 \n")
					println("Chapter 3: \n3.1 \n")
				}
			case done[line]:
				{
					switch line {
					case "1.1":
						ch1.Ex1()
					case "1.2":
						ch1.Ex2()
					case "1.3":
						ch1.Ex3()
					case "1.4":
						println(".txt files (with extension):")
						ch1.Ex4(enterArgs(0))
					case "1.5":
						println("Name of the output file (without extension):")
						ch1.Ex5(enterArgs(1))
					case "1.6":
						println("Name of the output file (without extension):")
						ch1.Ex6(enterArgs(1))
					case "2.2":
						ch2.Ex2()
					case "2.3", "2.4", "2.5":
						println("Please, for the benchmarks of exercises 2.3-2.5 run the file from cmd.")
					case "3.1":
						ch3.Ex1()
					}
				}

			default:
				{
					println("Sorry, there's no such exercise or command...")
				}
			}
		}

		println(hr)
	}
}
