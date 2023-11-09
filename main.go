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
    var openAiProvider = &OpenAiProvider{}

    initialPrompt := "You are a helpful AI assistant that helps people write resumes. Resume:" // TODO: Fix
    openAiProvider.Initialize(openAiToken, initialPrompt, nil)

    processResume(YamlResumeParser{}, openAiProvider, "resume_example.yaml", autoCorrect)

    // temp
    file, err := os.ReadFile("resume_example_output.yaml")

    if err != nil {
        fmt.Printf("Error reading file:%v\n ", err)
        return
    }

    fmt.Println(string(file))

    // delete file
    err = os.Remove("resume_example_output.yaml")

    if err != nil {
        return
    }
}

// TODO: refactor
func processResume(parser ResumeParser, aiProvider AiProvider, path string, autoCorrect bool) {
    resume := parser.Parse(path)
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

    outputPath := path + "_output.yaml"
    parser.Write(resume, outputPath)
}
