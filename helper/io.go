package helper

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ReadString(scanner bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func ReadInt(Scanner bufio.Scanner, prompt string) int {
	for {
		fmt.Print(prompt)
		Scanner.Scan()
		value, err := strconv.Atoi(Scanner.Text())
		if err == nil {
			return value
		}

		fmt.Printf("Invalid input: %v", err)
	}
}

func ReadInt64(scanner bufio.Scanner, prompt string) int64 {
	return int64(ReadInt(scanner, prompt))
}

func ReadIntChoice(scanner bufio.Scanner, prompt string, choices []string, values []int) int {
	for i, choice := range choices {
		choices[i] = fmt.Sprintf("%d- %s", i+1, choice)
	}
	prompt = fmt.Sprintf("%s\n%s\nEnter your choice: ", prompt, strings.Join(choices, "\n"))

	for {
		fmt.Print(prompt)

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err == nil && choice >= 1 && choice <= len(choices) {
			return values[choice-1]
		}

		fmt.Println("Invalid choice")
	}
}

func ReadTime(scanner bufio.Scanner, prompt string) time.Time {
	for {
		fmt.Print(prompt)

		scanner.Scan()
		value, err := time.Parse("2006-01-02 15:04:05", scanner.Text())
		if err == nil {
			return value
		}

		fmt.Printf("Invalid input: %v", err)
	}
}
