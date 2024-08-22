/*
 * job_application_list.test.go
 *
 * Test code for job_application_list.go
 *
 * Author: 1nf053c
 * Created: 2024-08-20
 * Last modified: 2024-08-20
 *
 * Version: 0.1
 */

package main

import (
	"os"
	"testing"
)

func FailIfErr(t *testing.T, e error) {
	if e != nil {
		t.Fatalf("err: %s", e.Error())
	}
}

func TestYamlFileToJson(t *testing.T) {
	yamlFilepath := "./raw/submitted_applications.yaml"
	jobApplicationList := JobApplicationList{}.FromYamlFile(yamlFilepath)
	jsonBytes := jobApplicationList.ToJson()
	jsonstr := condense(string(jsonBytes))

	EXPECTED_RESULT := condense(`[
            {
                "CompanyTitle": "Company ABC",
                "Details": {
                    "SubmittedDate": "08/19/2024",
                    "Location": "ExampleCity, ExampleState",
                    "Role": "Project Manager",
                    "Level": "Entry Level",
                    "Skills": [
                        "Go",
                        "Be Awesome"
                    ],
                    "Remote": true,
                    "Contract": false,
                    "ContractDuration": "",
                    "Platform": "linkedin",
                    "Resume": {
                        "Filename": "1-pager-2024-08-16.pdf",
                        "Filepath": "/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf"
                    },
                    "CoverLetter": null,
                    "Link": "https://www.linkedin.com/jobs/view/#####/?refId=####\u0026trackingId=####\u0026trk=###\u0026lipi=####",
                    "JobPostAndDescriptionAlignment": {
                        "CompanyTitle": {
                            "Status": "ok",
                            "Reason": "only listed once"
                        },
                        "JobTitle": {
                            "Status": "ok",
                            "Reason": "matches"
                        },
                        "RequiredSkills": {
                            "Status": "poor",
                            "Reason": "much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
                        }
                    }
                }
            },
            {
                "CompanyTitle": "Company DEF",
                "Details": {
                    "SubmittedDate": "08/18/2024",
                    "Location": "ExampleCity, ExampleState",
                    "Role": "Scrum Master",
                    "Level": "Intermediate Level",
                    "Skills": [
                        "Scrum Methodology",
                        "Be Awesome"
                    ],
                    "Remote": true,
                    "Contract": false,
                    "ContractDuration": "",
                    "Platform": "linkedin",
                    "Resume": {
                        "Filename": "1-pager-2024-08-16.pdf",
                        "Filepath": "/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf"
                    },
                    "CoverLetter": null,
                    "Link": "https://www.linkedin.com/jobs/view/#####/?refId=####\u0026trackingId=####\u0026trk=###\u0026lipi=####",
                    "JobPostAndDescriptionAlignment": {
                        "CompanyTitle": {
                            "Status": "ok",
                            "Reason": "only listed once"
                        },
                        "JobTitle": {
                            "Status": "ok",
                            "Reason": "matches"
                        },
                        "RequiredSkills": {
                            "Status": "poor",
                            "Reason": "much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
                        }
                    }
                }
            }
        ]`)

	if jsonstr != EXPECTED_RESULT {
		t.Fatalf("err: received %s, expected %s", jsonstr, EXPECTED_RESULT)
	}
}

func TestJsonToFile(t *testing.T) {
	EXPECTED_RESULT := condense(`[
    {
        "CompanyTitle": "Company ABC",
        "Details": {
            "SubmittedDate": "08/19/2024",
            "Location": "ExampleCity, ExampleState",
            "Role": "Project Manager",
            "Level": "Entry Level",
            "Skills": [
                "Go",
                "Be Awesome"
            ],
            "Remote": true,
            "Contract": false,
            "ContractDuration": "",
            "Platform": "linkedin",
            "Resume": {
                "Filename": "1-pager-2024-08-16.pdf",
                "Filepath": "/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf"
            },
            "CoverLetter": null,
            "Link": "https://www.linkedin.com/jobs/view/#####/?refId=####\u0026trackingId=####\u0026trk=###\u0026lipi=####",
            "JobPostAndDescriptionAlignment": {
                "CompanyTitle": {
                    "Status": "ok",
                    "Reason": "only listed once"
                },
                "JobTitle": {
                    "Status": "ok",
                    "Reason": "matches"
                },
                "RequiredSkills": {
                    "Status": "poor",
                    "Reason": "much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
                }
            }
        }
    },
    {
        "CompanyTitle": "Company DEF",
        "Details": {
            "SubmittedDate": "08/18/2024",
            "Location": "ExampleCity, ExampleState",
            "Role": "Scrum Master",
            "Level": "Intermediate Level",
            "Skills": [
                "Scrum Methodology",
                "Be Awesome"
            ],
            "Remote": true,
            "Contract": false,
            "ContractDuration": "",
            "Platform": "linkedin",
            "Resume": {
                "Filename": "1-pager-2024-08-16.pdf",
                "Filepath": "/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf"
            },
            "CoverLetter": null,
            "Link": "https://www.linkedin.com/jobs/view/#####/?refId=####\u0026trackingId=####\u0026trk=###\u0026lipi=####",
            "JobPostAndDescriptionAlignment": {
                "CompanyTitle": {
                    "Status": "ok",
                    "Reason": "only listed once"
                },
                "JobTitle": {
                    "Status": "ok",
                    "Reason": "matches"
                },
                "RequiredSkills": {
                    "Status": "poor",
                    "Reason": "much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
                }
            }
        }
    }
]`)

	yamlFilepath := "./raw/submitted_applications.yaml"
	jobApplicationList := JobApplicationList{}.FromYamlFile(yamlFilepath)

	outputFilepath := "./processed/json/submitted_applications.json"
	jobApplicationList.WriteToJsonFile(outputFilepath)

	filecontentbytes, err := os.ReadFile(outputFilepath)
	FailIfErr(t, err)

	filecontent := condense(string(filecontentbytes))

	if filecontent != EXPECTED_RESULT {
		t.Fatalf("got: %s, but expected: %s", filecontent, EXPECTED_RESULT)
	}
}

func TestJobApplicationListToCsvFile(t *testing.T) {
	EXPECTED_RESULT := `CompanyTitle,Details.SubmittedDate,Details.Location,Details.Role,Details.Level,Details.Skills,Details.Remote,Details.Contract,Details.ContractDuration,Details.Platform,Details.Resume.Filename,Details.Resume.Filepath,Details.CoverLetter,Details.Link,Details.JobPostAndDescriptionAlignment.CompanyTitle.Status,Details.JobPostAndDescriptionAlignment.CompanyTitle.Reason,Details.JobPostAndDescriptionAlignment.JobTitle.Status,Details.JobPostAndDescriptionAlignment.JobTitle.Reason,Details.JobPostAndDescriptionAlignment.RequiredSkills.Status,Details.JobPostAndDescriptionAlignment.RequiredSkills.Reason
Company ABC,08/19/2024,"ExampleCity, ExampleState",Project Manager,Entry Level,"[""Go"",""Be Awesome""]",true,false,,linkedin,1-pager-2024-08-16.pdf,/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf,,https://www.linkedin.com/jobs/view/#####/?refId=####&trackingId=####&trk=###&lipi=####,ok,only listed once,ok,matches,poor,"much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
Company DEF,08/18/2024,"ExampleCity, ExampleState",Scrum Master,Intermediate Level,"[""Scrum Methodology"",""Be Awesome""]",true,false,,linkedin,1-pager-2024-08-16.pdf,/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf,,https://www.linkedin.com/jobs/view/#####/?refId=####&trackingId=####&trk=###&lipi=####,ok,only listed once,ok,matches,poor,"much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
`
	yamlFilepath := "./raw/submitted_applications.yaml"
	jobApplicationList := JobApplicationList{}.FromYamlFile(yamlFilepath)

	outputFilepath := "processed/csv/submitted_applications.csv"
	jobApplicationList.WriteToCsvFile(outputFilepath)

	filecontentbytes, err := os.ReadFile(outputFilepath)
	filecontent := string(filecontentbytes)
	FailIfErr(t, err)

	if filecontent != EXPECTED_RESULT {
		t.Fatalf("got: %s, but expected: %s", filecontent, EXPECTED_RESULT)
	}
}
