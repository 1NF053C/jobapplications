/*
 * yeet.go
 *
 * Program that processes submitted job application details.
 *
 * Author: 1nf053c
 * Created: 2024-08-20
 * Last modified: 2024-08-20
 *
 * Version: 0.1
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jszwec/csvutil"
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

func processSubmittedApplicationsFile() {
	jobApplications := yamlToJobApplications()
	jobApplicationsToJsonFile(jobApplications)
	jobApplicationsToCsvFile(jobApplications)
}

func yamlToJobApplications() []JobApplication {
	filepath := "./raw/submitted_applications.yaml"
	yamlData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file %s", filepath)
	}
	var jobApplications []JobApplication
	if err := yaml.Unmarshal([]byte(yamlData), &jobApplications); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	return jobApplications
}

func jobApplicationsToJsonFile(jobApplications []JobApplication) {
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

func jobApplicationsToCsvFile(jobApplications []JobApplication) {
	jobApplicationCsvs := jobApplicationsToJobApplicationsCsv(jobApplications)
	csvBytes, err := csvutil.Marshal(jobApplicationCsvs)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	outputCsvFilepath := "processed/csv/submitted_applications.csv"
	if err := os.WriteFile(outputCsvFilepath, csvBytes, 0644); err != nil {
		log.Fatalf("Error writing to file %s", outputCsvFilepath)
	}
}
