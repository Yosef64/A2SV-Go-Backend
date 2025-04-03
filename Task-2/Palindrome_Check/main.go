package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Blue   = "\033[34m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the string : ")
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)
	var isPalendrome bool = checkPalindrome(inputString)

	if isPalendrome {
		println(Green + "The string is a palindrome")
	} else {
		println(Green + "The string is not a palindrome")
	}
}

func checkPalindrome(str string) bool {
	for left, right := 0, len(str)-1; left < right; left, right = left+1, right-1 {
		if str[left] != str[right] {
			return false
		}
	}
	return true
}
