package main

import (
    "fmt"
    "os"
    "strings"
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

    inputFileName := "resume_example.yaml"
    resumeUpdated := processResume(YamlResumeParser{}, openAiProvider, inputFileName, autoCorrect)

    outputFileName := strings.Split(inputFileName, ".")[0] + "_output.yaml"

    YamlResumeParser{}.Write(resumeUpdated, outputFileName)

    // temp
    file, err := os.ReadFile(outputFileName)

    if err != nil {
        fmt.Printf("Error reading file:%v\n ", err)
        return
    }

    fmt.Println(string(file))

    // delete file
    err = os.Remove(outputFileName)

    if err != nil {
        fmt.Printf("Error deleting file:%v\n ", err)
        return
    }
}

// TODO: refactor
func processResume(parser ResumeParser, aiProvider AiProvider, path string, autoCorrect bool) Resume {
    resume := parser.Parse(path)

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
            fmt.Printf("Checking highlight: %v\n", highlight)
            correctedHighlight := aiProvider.CheckSpellingGrammar(highlight)

            if correctedHighlight.Corrected == highlight {
                // No need to log corrected if it's already correct
                break // TODO: REMOVE THIS ONCE WE HAVE RATE LIMITING
                continue
            }

            correctedHighlight.Title = experience.Company + " - " + experience.Position
            correctedHighlights = append(correctedHighlights, correctedHighlight)

            if autoCorrect {
                highlight = correctedHighlight.Corrected
            }

            break // TODO: REMOVE THIS ONCE WE HAVE RATE LIMITING
        }
        break // TODO: REMOVE THIS ONCE WE HAVE RATE LIMITING
    }

    return resume
}
