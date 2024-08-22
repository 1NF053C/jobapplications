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
            "Status": {
                "Date": "08/22/2024",
                "Kind": "offer",
                "Explanation": "after I followed up I was scheduled for a next day screening and then immediately extended an offer"
            },
            "Events": [
                {
                    "Event": {
                        "Date": "08/19/2024",
                        "Kind": "submitted application",
                        "Explanation": "Filled out form on the hiring platform, clicked submit"
                    }
                },
                {
                    "Event": {
                        "Date": "08/20/2024",
                        "Kind": "follow-up",
                        "Explanation": "Emailed and called and left a message, the recruiter scheduled me in for a 30min screening"
                    }
                },
                {
                    "Event": {
                        "Date": "08/22/2024",
                        "Kind": "screening",
                        "Explanation": "Had screening call w/ team, they had me skip next round of interviewing and extended an offer"
                    }
                },
                {
                    "Event": {
                        "Date": "08/22/2024",
                        "Kind": "offer",
                        "Explanation": "Extended offer immediately after I demo'ed my portfolio of work"
                    }
                }
            ],
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
            "SubmittedDate": "08/11/2024",
            "Status": {
                "Date": "08/22/2024",
                "Kind": "no response",
                "Explanation": "I followed up on 08/18/2024 and am waiting to hear back"
            },
            "Events": [
                {
                    "Event": {
                        "Date": "08/11/2024",
                        "Kind": "submitted application",
                        "Explanation": "Filled out form on the hiring platform, clicked submit"
                    }
                },
                {
                    "Event": {
                        "Date": "08/18/2024",
                        "Kind": "follow-up",
                        "Explanation": "Emailed and called and left a message for the recruiter"
                    }
                }
            ],
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
            "Status": {
                "Date": "08/22/2024",
                "Kind": "offer",
                "Explanation": "after I followed up I was scheduled for a next day screening and then immediately extended an offer"
            },
            "Events": [
                {
                    "Event": {
                        "Date": "08/19/2024",
                        "Kind": "submitted application",
                        "Explanation": "Filled out form on the hiring platform, clicked submit"
                    }
                },
                {
                    "Event": {
                        "Date": "08/20/2024",
                        "Kind": "follow-up",
                        "Explanation": "Emailed and called and left a message, the recruiter scheduled me in for a 30min screening"
                    }
                },
                {
                    "Event": {
                        "Date": "08/22/2024",
                        "Kind": "screening",
                        "Explanation": "Had screening call w/ team, they had me skip next round of interviewing and extended an offer"
                    }
                },
                {
                    "Event": {
                        "Date": "08/22/2024",
                        "Kind": "offer",
                        "Explanation": "Extended offer immediately after I demo'ed my portfolio of work"
                    }
                }
            ],
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
            "SubmittedDate": "08/11/2024",
            "Status": {
                "Date": "08/22/2024",
                "Kind": "no response",
                "Explanation": "I followed up on 08/18/2024 and am waiting to hear back"
            },
            "Events": [
                {
                    "Event": {
                        "Date": "08/11/2024",
                        "Kind": "submitted application",
                        "Explanation": "Filled out form on the hiring platform, clicked submit"
                    }
                },
                {
                    "Event": {
                        "Date": "08/18/2024",
                        "Kind": "follow-up",
                        "Explanation": "Emailed and called and left a message for the recruiter"
                    }
                }
            ],
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
	EXPECTED_RESULT := `CompanyTitle,Details.SubmittedDate,Details.Status.Date,Details.Status.Kind,Details.Status.Explanation,Details.Events,Details.Location,Details.Role,Details.Level,Details.Skills,Details.Remote,Details.Contract,Details.ContractDuration,Details.Platform,Details.Resume.Filename,Details.Resume.Filepath,Details.CoverLetter,Details.Link,Details.JobPostAndDescriptionAlignment.CompanyTitle.Status,Details.JobPostAndDescriptionAlignment.CompanyTitle.Reason,Details.JobPostAndDescriptionAlignment.JobTitle.Status,Details.JobPostAndDescriptionAlignment.JobTitle.Reason,Details.JobPostAndDescriptionAlignment.RequiredSkills.Status,Details.JobPostAndDescriptionAlignment.RequiredSkills.Reason
Company ABC,08/19/2024,08/22/2024,offer,after I followed up I was scheduled for a next day screening and then immediately extended an offer,"[{""Event"":{""Date"":""08/19/2024"",""Kind"":""submitted application"",""Explanation"":""Filled out form on the hiring platform, clicked submit""}},{""Event"":{""Date"":""08/20/2024"",""Kind"":""follow-up"",""Explanation"":""Emailed and called and left a message, the recruiter scheduled me in for a 30min screening""}},{""Event"":{""Date"":""08/22/2024"",""Kind"":""screening"",""Explanation"":""Had screening call w/ team, they had me skip next round of interviewing and extended an offer""}},{""Event"":{""Date"":""08/22/2024"",""Kind"":""offer"",""Explanation"":""Extended offer immediately after I demo'ed my portfolio of work""}}]","ExampleCity, ExampleState",Project Manager,Entry Level,"[""Go"",""Be Awesome""]",true,false,,linkedin,1-pager-2024-08-16.pdf,/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf,,https://www.linkedin.com/jobs/view/#####/?refId=####&trackingId=####&trk=###&lipi=####,ok,only listed once,ok,matches,poor,"much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
Company DEF,08/11/2024,08/22/2024,no response,I followed up on 08/18/2024 and am waiting to hear back,"[{""Event"":{""Date"":""08/11/2024"",""Kind"":""submitted application"",""Explanation"":""Filled out form on the hiring platform, clicked submit""}},{""Event"":{""Date"":""08/18/2024"",""Kind"":""follow-up"",""Explanation"":""Emailed and called and left a message for the recruiter""}}]","ExampleCity, ExampleState",Scrum Master,Intermediate Level,"[""Scrum Methodology"",""Be Awesome""]",true,false,,linkedin,1-pager-2024-08-16.pdf,/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf,,https://www.linkedin.com/jobs/view/#####/?refId=####&trackingId=####&trk=###&lipi=####,ok,only listed once,ok,matches,poor,"much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"
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
