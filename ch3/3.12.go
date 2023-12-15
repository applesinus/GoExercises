package ch3

import (
	"fmt"
)

func Ex12() {
	var str1, str2 string
	print("Enter the first string: ")
	fmt.Scan(&str1)
	print("Enter the second string: ")
	fmt.Scan(&str2)

	slc1 := make([]rune, 0, len(str1))
	slc2 := make([]rune, 0, len(str2))

	for _, run := range str1 {
		slc1 = append(slc1, run)
	}
	for _, run := range str2 {
		slc2 = append(slc2, run)
	}

	if len(slc1) == len(slc2) {
		for i := range str1 {
			if slc1[i] != slc2[len(slc2)-i-1] {
				println("Not an anagram")
				return
			}
		}
	} else {
		println("Not an anagram")
		return
	}

	println("The words are anagrams for each other")
	return
}
