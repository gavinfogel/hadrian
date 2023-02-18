package main

import (
	"fmt"
	"hadrian/openai"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	openaiApiKey := os.Getenv("OPENAI_KEY")

	oai := openai.NewOpenaiClient(openaiApiKey)

	completedText, err := oai.Complete("Give the best choice of the following foods: pizza, wings, hotdogs, or salad: ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(completedText)
}
