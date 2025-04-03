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
	fmt.Print(Yellow + "Enter a string: " + Reset)
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)
	result := calculateFrequency(inputString)
	fmt.Printf(Green+"%-10s%-10s\n", "Character", "Frequency")
	fmt.Println("----------------------")

	for char, count := range result {
		fmt.Printf(Blue+"%-10s%-10d\n", char, count)
	}

}

func calculateFrequency(str string) map[string]int {
	frequecies := make(map[string]int)
	for _, char := range str {
		if chr := string(char); chr != "!" {
			frequecies[chr]++
		}
	}
	return frequecies
}
