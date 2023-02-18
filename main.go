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

	completedText, err := oai.Complete("You are a genius at reading and writing, and an expert on Roman private life. Understanding all of the nuance of private life, your colegue and friend who teaches on the topic wants to provide a sample of a very good answer to his essay for his students. He asks you to answer the following prompt: \n First Writing Assignment: Critical Fabulation of 'A Day in the Life' \n For the first writing assignment, you are tasked with writing a fictional account of the experience of an ancient individual living in Ancient Roman Italy (e.g. the city of Rome, or Pompeii). This individual’s identity is up to you—this could be a wife of a prominent city official, a soldier, an enslaved individual, a gladiator, etc. The narrative that you tell from that individual’s perspective must be informed by archaeological, artistic, and literary evidence discussed in our class lectures and the assigned readings. Moreover, consider the ways that the multiple identities of your individual (gender, sex, age, citizenship status, ethnicity or cultural background, economic status, occupation, etc.) would inform how you navigate the Roman world as we have explored it so far in class. \n But despite the valuable sources shared in class, there will still be 'unknowns' that our evidence cannot speak to. This is where critical fabulation comes in. Historian Saidiya Hartman, professor of English and comparative literature at Columbia University, coined the term ‘critical fabulation’ in her essay “Venus in Two Acts.” Critical fabulation is the combining of historical and archival research with critical theory and speculative narrative to fill in the blanks left in the historical record. It is both a fleshing out and a problematizing of history and the stories we can tell or recover with our limited evidence from the past. While Hartman studies those individuals implicated in, consumed by, and ultimately lost to the transatlantic slave trade, we can appropriate this methodology (while crediting Hartman and acknowledging the radically different contexts in which she works) to, “imagine what might have happened or might have been said or might have been done” in a past time otherwise lost to us.")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(completedText)
}
