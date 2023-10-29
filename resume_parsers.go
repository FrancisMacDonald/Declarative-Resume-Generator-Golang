package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type ResumeParser interface {
	Parse(path string) Resume
	Write(resume Resume, path string)
}

type YamlResumeParser struct{}
type JsonResumeParser struct{}

func (parser YamlResumeParser) Parse(path string) Resume {
	filename, _ := filepath.Abs(path)
	yamlContents, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error reading file:%v\n ", err)

		return Resume{}
	}

	resume := Resume{}
	err = yaml.Unmarshal(yamlContents, &resume)

	if err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)

		return Resume{}
	}

	return resume
}

func (parser JsonResumeParser) Parse(path string) Resume {
	panic("JsonResumeParser not implemented")
}

func (parser YamlResumeParser) Write(resume Resume, path string) {
	filename, _ := filepath.Abs(path)
	yamlContents, err := yaml.Marshal(resume)

	if err != nil {
		fmt.Printf("Error marshalling YAML: %v\n", err)
		return
	}

	err = os.WriteFile(filename, yamlContents, 0644)

	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}
