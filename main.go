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

	vec, err := oai.Embed("hello I am gavin the idiot")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vec)

	fmt.Println("I am a golang project")
}
