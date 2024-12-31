package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Question struct {
	question     string
	options      []string
	correctIndex int
}

func main() {
	questions := []Question{
		{"In what year did the Great October Socialist Revolution take place?", []string{"1. 1917", "2. 1923", "3. 1914", "4. 1920"}, 1},
		{"What is the largest lake in the world?", []string{"1. Caspian Sea", "2. BaikalLake", "3. Superior", "4. Ontario"}, 2},
		{"Which planet in the solar system is known as the “Red Planet”?", []string{"1. Venus", "2. Earth", "3. Mars", "4. Jupiter"}, 3},
	}

	var score int
	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("Enter 'exit' to quit anytime.")
	fmt.Println("\n")

	for i, q := range questions {
		fmt.Printf("Q%d: %s\n", i+1, q.question)
		for _, option := range q.options {
			fmt.Println(option)
		}

		for {
			fmt.Print("Enter your answer: ")
			var input string
			fmt.Scanln(&input)

			input = strings.TrimSpace(input)
			if input == "exit" {
				fmt.Println("Exiting... Thank you!")
				goto End
			}
			answer, err := strconv.Atoi(input)
			if err != nil || answer < 1 || answer > len(q.options) {
				fmt.Println("Invalid input. Try again.")
				continue
			}
			if answer == q.correctIndex {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong! Correct answer:", q.options[q.correctIndex])
			}
			break
		}
	}

End:
	fmt.Println("=====================================")
	fmt.Printf("Quiz done! Your score: %d/%d\n", score, len(questions))

	if score == len(questions) {
		fmt.Println("Excellent!")
	} else if score >= len(questions)/2 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Needs Improvement.")
	}

	fmt.Println("Thanks for taking the quiz!")
}
