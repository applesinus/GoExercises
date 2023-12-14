package ch1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Ex1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func Ex2() {
	for index, arg := range os.Args[0:] {
		fmt.Printf("%d: %s\n", index+1, arg)
	}
}

func Ex3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("Time to execute strings.Join(): %vns\n\n", time.Since(start).Nanoseconds())

	start = time.Now()
	for _, arg := range os.Args[0:] {
		fmt.Println(arg + " ")
	}
	fmt.Printf("Time to execute unefficient way: %vns\n", time.Since(start).Nanoseconds())
	fmt.Println("If there's no difference, well, I've tried...")
}
