package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrIncorrectNumberOfArgs = fmt.Errorf("incorrect number of args")
var ErrInvalidArg = fmt.Errorf("invalid arg")

func main() {
	defer os.Exit(0)
	processArgs(os.Args)
	fmt.Print("success")
}

func processArgs(args []string) {
	if len(args) != 2 {
		fmt.Print(ErrIncorrectNumberOfArgs)
		os.Exit(1)
	}

	arg := os.Args[1]

	if arg != "process" {
		fmt.Print(ErrInvalidArg)
		os.Exit(1)
	}

	if arg == "process" {
		processSubmittedApplicationsFile()
	}
}

type JobApplicationDetails struct {
	SubmittedDate                  string        `yaml:"submittedDate"`
	Location                       string        `yaml:"location"`
	Role                           string        `yaml:"role"`
	Level                          string        `yaml:"level"`
	Skills                         []string      `yaml:"skills"`
	Remote                         bool          `yaml:"remote"`
	Contract                       bool          `yaml:"contract,omitempty"`
	ContractDuration               string        `yaml:"contractDuration,omitempty"`
	Platform                       string        `yaml:"platform"`
	Resume                         Resume        `yaml:"resume"`
	CoverLetter                    interface{}   `yaml:"coverLetter"`
	Link                           string        `yaml:"link"`
	JobPostAndDescriptionAlignment AlignmentItem `yaml:"jobPostAndDescriptionAlignment"`
}

type Resume struct {
	Filename string `yaml:"filename"`
	Filepath string `yaml:"filepath"`
}

type AlignmentItem struct {
	CompanyTitle   *AlignmentDetail `yaml:"companyTitle,omitempty"`
	JobTitle       *AlignmentDetail `yaml:"jobTitle,omitempty"`
	RequiredSkills *AlignmentDetail `yaml:"requiredSkills,omitempty"`
}

type AlignmentDetail struct {
	Status string `yaml:"status"`
	Reason string `yaml:"reason"`
}

func processSubmittedApplicationsFile() {
	filepath := "./raw/submitted_applications.yaml"
	yamlData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file %s", filepath)
	}

	var jobApplications []JobApplicationDetails
	if err := yaml.Unmarshal([]byte(yamlData), &jobApplications); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	jsonBytes, err := json.MarshalIndent(jobApplications, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	outputJsonFilepath := "processed/json/submitted_applications.json"
	if err := os.WriteFile(outputJsonFilepath, jsonBytes, 0644); err != nil {
		log.Fatalf("Error writing to file %s", outputJsonFilepath)
	}
}
