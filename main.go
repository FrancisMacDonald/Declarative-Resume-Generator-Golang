package main

import (
	"fmt"
	"os"
)

type CorrectedText struct {
	Title     string
	Original  string
	Corrected string
	Changes   string
}

func main() {
	autoCorrect := true

	openAiToken := os.Getenv("OPENAI_TOKEN")

	openAiProvider := OpenAiProvider{}
	openAiProvider.Initialize(openAiToken)

	processResume(YamlResumeParser{}, openAiProvider, autoCorrect)

	// delete file
	err := os.Remove("resume_example_output.yaml")

	if err != nil {
		return
	}
}

func processResume(parser ResumeParser, aiProvider AiProvider, autoCorrect bool) {
	resume := parser.Parse("resume_example.yaml")
	fmt.Println(resume)

	// Correct summary
	correctedSummary := aiProvider.CheckSpellingGrammar(resume.Summary)
	correctedSummary.Title = "Summary"

	if autoCorrect {
		resume.Summary = correctedSummary.Corrected
	}

	var correctedHighlights []CorrectedText

	// Correct experience highlights
	for _, experience := range resume.Experience {
		for _, highlight := range experience.Highlights {
			correctedHighlight := aiProvider.CheckSpellingGrammar(highlight)

			if correctedHighlight.Corrected == highlight {
				// No need to log corrected if it's already correct
				continue
			}

			correctedHighlight.Title = experience.Company + " - " + experience.Position
			correctedHighlights = append(correctedHighlights, correctedHighlight)

			if autoCorrect {
				highlight = correctedHighlight.Corrected
			}
		}
	}

	parser.Write(resume, "resume_example_corrected.yaml")
}
