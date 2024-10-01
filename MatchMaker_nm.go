package main

import (
	"fmt"
	"math"
)

const (
	TrueLoveScore = 67
	FriendsScore  = 34
	//RunawayScore has no usages
)

// Default Questions
var questions = []string{
	"I enjoy spending time outdoors and being active",
	"Family is the most important aspect in my life",
	"Open and honest communication is the most important part of a relationship",
	"I would prefer a partner who enjoys the same hobbies as me",
	"I am willing to relocate for the right relationship",
}

var desiredValues = []int{5, 5, 5, 5, 5}

func main() {
	userAnswers := make([]int, len(questions))
	var totalPercentage int
	//Display Header
	fmt.Print("💖💖💖Welcome to MatchMaker by Nicole💖💖💖\n" +
		"Find out if you and your potential match are truly compatible!\n" +
		"Please respond to the following questions with a number from 1-5:\n" +
		"1 = Strongly Disagree\t2 = Disagree\t3 = Neutral\t4 = Agree\t5 = Strongly Agree\n")
	fmt.Println("==========================================================================================")

	for i, question := range questions {
		userAnswers[i] = validateInput(question)
		compatibilityPercentage := calculateCompatibility(userAnswers[i], desiredValues[i])
		totalPercentage += userAnswers[i]

		//Display each question's compatability
		fmt.Printf("\n💬Question %d: %s\n", i+1, question)
		fmt.Printf("💞Your compatibility for this question: %.2f%%💞\n", 100-compatibilityPercentage)
		//fmt.Println(totalPercentage)
		fmt.Println("------------------------------------------------------------------------------------------")
	}

	percentage := (float64(totalPercentage) / 25) * 100
	fmt.Printf("💗Your final compatibility score: %.2f%%💗\n", percentage)
	finalScore := int(percentage)
	printResult(finalScore)
}

// Ensuring input is a number 1-5 or error message
func validateInput(question string) int {
	var answer int
	for {
		fmt.Printf("%s (1-5): ", question)
		_, inError := fmt.Scan(&answer)
		if inError != nil {
			fmt.Println("❗️Invalid input. Please enter a number 1-5.")
			continue
		}
		if answer < 1 || answer > 5 {
			fmt.Println("❗️Invalid input. Please enter a number 1-5.")
			continue
		}
		return answer
	}
}

// Calculating absolute difference
func calculateCompatibility(userAnswer, desiredValue int) float64 {
	diff := math.Abs(float64(userAnswer - desiredValue))
	return diff * 25
}

// Evaluating final score and displaying closing message
func printResult(finalScore int) {
	if finalScore >= TrueLoveScore {
		fmt.Println("\n💓True Love! You are a perfect match!💓")
	} else if finalScore >= FriendsScore {
		fmt.Println("\n👯‍♀️You might make good friends.👯‍♀️")
	} else {
		fmt.Println("\n🚩🚩Run away.🚩🚩")
	}
	fmt.Println("💖💖💖Thank you for using MatchMaker by Nicole💖💖💖")
}
