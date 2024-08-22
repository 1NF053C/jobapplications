/*
 * job_application_list.go
 *
 * Code that defines a JobApplicationList and related funcs
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

	"log"
	"os"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v3"
)

type JobApplicationList []JobApplication

type JobApplication struct {
	CompanyTitle string
	Details      JobApplicationDetails `yaml:"job_application_details"`
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
	Resume                         Resume        `yaml:"resume"` // Nested struct, handled separately
	CoverLetter                    interface{}   `yaml:"coverLetter"`
	Link                           string        `yaml:"link"`
	JobPostAndDescriptionAlignment AlignmentItem `yaml:"jobPostAndDescriptionAlignment"` // Nested struct, handled separately
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

func (j JobApplicationList) FromYamlFile(filepath string) JobApplicationList {
	yamlData, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading file %s", filepath)
		os.Exit(1)
	}

	var jobApplicationList JobApplicationList
	if err := yaml.Unmarshal(yamlData, &jobApplicationList); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
		os.Exit(1)
	}
	return jobApplicationList
}

func (jal *JobApplicationList) UnmarshalYAML(value *yaml.Node) error {
	// UnmarshalYAML will receive a SequenceNode for arrays
	if value.Kind != yaml.SequenceNode {
		log.Fatalf("expected a sequence node")
		os.Exit(1)
	}

	// value.Content is the array element, it should be size 2 when two companies are in the list
	// I have a company name as the key, and the value is another node
	for _, item := range value.Content {
		// item should be a MappingNode, which is a single key value pair
		if item.Kind != yaml.MappingNode || len(item.Content) != 2 {
			log.Fatalf("expected a mapping with a single key-value pair")
			os.Exit(1)
		}

		// key is item.Content[0]
		companyNode := item.Content[0]
		// value is item.Content[1], which is the rest of the fields in JobApplication except CompanyTitle
		detailsNode := item.Content[1]

		var ja JobApplication
		// CompanyTitle value can be taken from the key
		ja.CompanyTitle = companyNode.Value

		// no custom yaml unmarshallers should be needed for the rest of JobApplication
		// so this Decode func should handle the rest and update ja
		if err := detailsNode.Decode(&ja.Details); err != nil {
			return err
		}

		// append each decoded JobApplicationDetails to this JobApplicationList
		*jal = append(*jal, ja)
	}

	// success
	return nil
}

func (j JobApplicationList) ToJson() []byte {
	jsonBytes, err := json.MarshalIndent(j, "", "    ")
	checkErr(err)
	return jsonBytes
}

func (j JobApplicationList) WriteToJsonFile(filepath string) {
	bytes := j.ToJson()
	err := os.WriteFile(filepath, bytes, 0644)
	checkErr(err)
}

func (j JobApplicationList) ToCsv() string {
	csvString, err := gocsv.MarshalString(&j)
	checkErr(err)
	return csvString
}

func (j JobApplicationList) WriteToCsvFile(filepath string) {
	csvString := j.ToCsv()
	err := os.WriteFile(filepath, []byte(csvString), 0644)
	checkErr(err)
}
