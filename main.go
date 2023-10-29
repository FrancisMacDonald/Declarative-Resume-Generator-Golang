package main

import (
	"fmt"
	"os"
)

type CorrectedText struct {
	Title     string
	Original  string
	Corrected string
}

func main() {
	openAiToken := os.Getenv("OPENAI_TOKEN")

	openAiProvider := OpenAiProvider{}
	openAiProvider.Initialize(openAiToken)

	processResume(YamlResumeParser{}, openAiProvider)

	// delete file
	err := os.Remove("resume_example_output.yaml")

	if err != nil {
		return
	}
}

func processResume(parser ResumeParser, aiProvider AiProvider) {
	resume := parser.Parse("resume_example.yaml")
	fmt.Println(resume)

	checkGrammarAndSpelling(aiProvider, resume)

	parser.Write(resume, "resume_example_output.yaml")
}

func checkGrammarAndSpelling(aiProvider AiProvider, resume Resume) {
	correctedSummary := aiProvider.CheckSpellingGrammar(resume.Summary)
	correctedSummary.Title = "Summary"

	var correctedHighlights []CorrectedText

	// Correct spelling and grammar on all experience highlights
	for _, experience := range resume.Experience {
		for _, highlight := range experience.Highlights {
			correctedHighlight := aiProvider.CheckSpellingGrammar(highlight)
			correctedHighlight.Title = experience.Company + " - " + experience.Position
			correctedHighlights = append(correctedHighlights, correctedHighlight)
		}
	}

}
