package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fmt.Print(Yellow + "Enter Your Name: " + Reset)
	studentName, _ := reader.ReadString('\n')
	studentName = strings.TrimSpace(studentName)

	fmt.Print(Blue + "Enter how many subjects you took: " + Reset)
	var numberOfSubjects, total int
	fmt.Scan(&numberOfSubjects)

	fmt.Println(Green + "\n Now, enter the subject and the grade. Example: English 85" + Reset)

	subjects := make(map[string]int)

	for i := 0; i < numberOfSubjects; {
		fmt.Printf(Yellow+"Enter subject No. %d: "+Reset, i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		sAndgPair := strings.Fields(input)
		if len(sAndgPair) < 2 {
			fmt.Println(Red + "Invalid input! Please enter in format: Subject Grade" + Reset)
			continue
		}

		subject, gradeStr := sAndgPair[0], sAndgPair[1]
		if subjects[subject] != 0 {
			fmt.Println(Red, "You cannot have multiple grades for the same subject!")
			continue

		}
		grade, err := strconv.Atoi(strings.TrimSpace(gradeStr))
		if err != nil {
			fmt.Println(Red+"Invalid grade! Must be a number."+Reset, err)
			continue
		}

		subjects[subject] = grade
		i++
	}

	fmt.Println(Green + "\n Summary of Grades for " + studentName + ":" + Reset)
	fmt.Println("────────────────────────────────────")
	for subject, grade := range subjects {
		total += grade
		fmt.Printf(Blue+" %-15s : %3d\n"+Reset, subject, grade)
	}
	fmt.Println("────────────────────────────────────")
	fmt.Printf(Blue+" %-15s : %3d\n"+Reset, "Average", total/numberOfSubjects)

	fmt.Println("────────────────────────────────────")
	fmt.Println(Green + "Thank you for using the grade recorder!" + Reset)
}
