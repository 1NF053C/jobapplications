package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v3"
)

type JobApplicationList []JobApplication
type JobApplicationListCsvEncodingReady []JobApplicationCsvEncodingReady

func (j JobApplicationList) FromYamlFile(filepath string) JobApplicationList {
	yamlData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file %s", filepath)
	}
	var jobApplicationList JobApplicationList
	if err := yaml.Unmarshal([]byte(yamlData), &jobApplicationList); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	return jobApplicationList
}

func (j JobApplicationList) writeToJsonFile() {
	jsonBytes, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	outputJsonFilepath := "processed/json/submitted_applications.json"
	if err := os.WriteFile(outputJsonFilepath, jsonBytes, 0644); err != nil {
		log.Fatalf("Error writing to file %s", outputJsonFilepath)
	}
}

func (j JobApplicationList) writeToCsvFile() {
	jobApplicationsCsvList := j.ToCsvEncodingReady()
	csvBytes, err := csvutil.Marshal(jobApplicationsCsvList)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	outputCsvFilepath := "processed/csv/submitted_applications.csv"
	if err := os.WriteFile(outputCsvFilepath, csvBytes, 0644); err != nil {
		log.Fatalf("Error writing to file %s", outputCsvFilepath)
	}
}

func (jl JobApplicationList) ToCsvEncodingReady() JobApplicationListCsvEncodingReady {
	jlCsvEncReady := make(JobApplicationListCsvEncodingReady, len(jl))
	for i, j := range jl {
		jlCsvEncReady[i] = j.ToCsvEncodingReady()
	}
	return jlCsvEncReady
}

type JobApplicationCsvEncodingReady struct {
	JobApplication
	Details JobApplicationDetailsCsvEncodingReady `csv:"-"`
}

type JobApplicationDetailsCsvEncodingReady struct {
	JobApplicationDetails
	Skills string `csv:"skills"`
}

func (j JobApplication) ToCsvEncodingReady() JobApplicationCsvEncodingReady {
	return JobApplicationCsvEncodingReady{
		JobApplication: j,
		Details:        j.Details.ToCsvEncodingReady(),
	}
}

func (j JobApplicationDetails) ToCsvEncodingReady() JobApplicationDetailsCsvEncodingReady {
	return JobApplicationDetailsCsvEncodingReady{
		JobApplicationDetails: j,
		Skills:                strings.Join(j.Skills, ", "),
	}
}

type JobApplication struct {
	CompanyTitle string                `csv:"company_title"`
	Details      JobApplicationDetails `yaml:"-" csv:"-"`
	// todo: try embedded struct here
}

type JobApplicationDetails struct {
	SubmittedDate                  string        `yaml:"submittedDate" csv:"submitted_date"`
	Location                       string        `yaml:"location" csv:"location"`
	Role                           string        `yaml:"role" csv:"role"`
	Level                          string        `yaml:"level" csv:"level"`
	Skills                         []string      `yaml:"skills" csv:"-"`
	Remote                         bool          `yaml:"remote" csv:"remote"`
	Contract                       bool          `yaml:"contract,omitempty" csv:"contract"`
	ContractDuration               string        `yaml:"contractDuration,omitempty" csv:"contract_duration"`
	Platform                       string        `yaml:"platform" csv:"platform"`
	Resume                         Resume        `yaml:"resume" csv:"-"` // Nested struct, handled separately
	CoverLetter                    interface{}   `yaml:"coverLetter" csv:"cover_letter"`
	Link                           string        `yaml:"link" csv:"link"`
	JobPostAndDescriptionAlignment AlignmentItem `yaml:"jobPostAndDescriptionAlignment" csv:"-"` // Nested struct, handled separately
}

func (j *JobApplication) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode || len(value.Content) != 2 {
		log.Fatal("expected a mapping node with two key-value pairs")
		os.Exit(1)
	}
	j.CompanyTitle = value.Content[0].Value
	return value.Content[1].Decode(&j.Details)
}

type Resume struct {
	Filename string `yaml:"filename" csv:"resume_filename"`
	Filepath string `yaml:"filepath" csv:"resume_filepath"`
}

type AlignmentItem struct {
	CompanyTitle   *AlignmentDetail `yaml:"companyTitle,omitempty" csv:"company_title_status,company_title_reason"`
	JobTitle       *AlignmentDetail `yaml:"jobTitle,omitempty" csv:"job_title_status,job_title_reason"`
	RequiredSkills *AlignmentDetail `yaml:"requiredSkills,omitempty" csv:"required_skills_status,required_skills_reason"`
}

type AlignmentDetail struct {
	Status string `yaml:"status" csv:"status"`
	Reason string `yaml:"reason" csv:"reason"`
}
