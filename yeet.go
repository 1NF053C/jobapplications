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
	"log"
	"os"
	"strings"
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
		processJobApplicationListFile()
	}
}

func processJobApplicationListFile() {
	yamlFilepath := "./raw/submitted_applications.yaml"
	jobApplicationList := JobApplicationList{}.FromYamlFile(yamlFilepath)

	// write job application list to a json file
	outputJsonFilepath := "processed/json/submitted_applications.json"
	jobApplicationList.WriteToJsonFile(outputJsonFilepath)

	// write job application list to a csv file
	outputCsvFilepath := "processed/csv/submitted_applications.csv"
	jobApplicationList.WriteToCsvFile(outputCsvFilepath)
}

func checkErr(e error) {
	if e != nil {
		log.Fatalf("err: %s", e.Error())
		os.Exit(1)
	}
}

func condense(s string) string {
	return strings.Join(strings.Fields(s), "")
}
