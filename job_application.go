package main

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// TODO FIX THIS
// current issue is with csv encoding.
// company_title is recognized as a key and 2 company titles show up as expected
// but no other columns show up.
// Details is handled similarily to skills, but skills is simpler because there is one layer and it resolves to a string
// where-as JobApplication.Details => JobApplication => JobApplicationCsv => JobApplicationCsv.Details
type JobApplication struct {
	CompanyTitle string                `csv:"company_title"`
	Details      JobApplicationDetails `yaml:"-" csv:"-"`
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

type JobApplicationCsv struct {
	JobApplication
	Details JobApplicationDetailsCsv `csv:"-"`
}

type JobApplicationDetailsCsv struct {
	JobApplicationDetails
	Skills string `csv:"skills"`
}

func (ja *JobApplication) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode || len(value.Content) != 2 {
		log.Fatal("expected a mapping node with two key-value pairs")
		os.Exit(1)
	}
	ja.CompanyTitle = value.Content[0].Value
	return value.Content[1].Decode(&ja.Details)
}

// JobApplicationCsv is the same but its Details field must be converted to JobApplicationDetailsCsv
// so that Skills can be represented as a string
// The JobApplication field within is flattened, and the Details field should override same named field
func (j JobApplication) ToJobApplicationCsv() JobApplicationCsv {
	return JobApplicationCsv{
		JobApplication: j,
		Details:        j.Details.ToCsv(),
	}
}

// Skills must be represeted as a string so it can fit in a single csv column,
// The JobApplicationDetails field is flattened, Skills field should override same named field
func (j JobApplicationDetails) ToCsv() JobApplicationDetailsCsv {
	return JobApplicationDetailsCsv{
		JobApplicationDetails: j,
		Skills:                strings.Join(j.Skills, ", "),
	}
}

// alternatively make JobApplicationList a named type
func jobApplicationsToJobApplicationsCsv(jobApplications []JobApplication) []JobApplicationCsv {
	jobApplicationCsvs := make([]JobApplicationCsv, len(jobApplications))
	for i, app := range jobApplications {
		jobApplicationCsvs[i] = app.ToJobApplicationCsv()
	}
	return jobApplicationCsvs
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
