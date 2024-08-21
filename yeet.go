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
	"fmt"
	"os"
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
	yamlFilepath := "./raw/submitted_applications.yaml"
	jobApplicationList := JobApplicationList{}.FromYamlFile(yamlFilepath)
	jobApplicationList.writeToJsonFile()
	jobApplicationList.writeToCsvFile()
}
