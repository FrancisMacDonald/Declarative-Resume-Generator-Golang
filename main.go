package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func main() {
	resume := parseResumeFromYaml("resume_example.yaml")
	fmt.Println(resume)
	writeResumeToYaml(resume, "resume_example_output.yaml")

	// delete file
	err := os.Remove("resume_example_output.yaml")

	if err != nil {
		return
	}
}

func writeResumeToYaml(resume Resume, path string) {
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

func parseResumeFromYaml(path string) Resume {
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
