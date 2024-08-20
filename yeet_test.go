package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func runCmd(t *testing.T, cmdParts ...string) {
	if err := exec.Command(cmdParts[0], cmdParts[1:]...).Run(); err != nil {
		t.Fatalf("cmd failed, exiting tests. Err: %s", err.Error())
	}
}

func TestProcessSubmittedApplicationsFile(t *testing.T) {
	EXPECTED_RESULT := `[{"SubmittedDate":"08/19/2024","Location":"ExampleCity, ExampleState","Role":"Project Manager","Level":"Entry Level","Skills":["Go"],"Remote":true,"Contract":false,"ContractDuration":"","Platform":"linkedin","Resume":{"Filename":"1-pager-2024-08-16.pdf","Filepath":"/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf"},"CoverLetter":null,"Link":"https://www.linkedin.com/jobs/view/#####/?refId=####\u0026trackingId=####\u0026trk=###\u0026lipi=####","JobPostAndDescriptionAlignment":{"CompanyTitle":{"Status":"ok","Reason":"only listed once"},"JobTitle":{"Status":"ok","Reason":"matches"},"RequiredSkills":{"Status":"poor","Reason":"much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"}}},{"SubmittedDate":"08/18/2024","Location":"ExampleCity, ExampleState","Role":"Scrum Master","Level":"Intermediate Level","Skills":["Scrum Methodology"],"Remote":true,"Contract":false,"ContractDuration":"","Platform":"linkedin","Resume":{"Filename":"1-pager-2024-08-16.pdf","Filepath":"/home/a/projects/job-applications/resumes/pdf/1-pager-2024-08-16.pdf"},"CoverLetter":null,"Link":"https://www.linkedin.com/jobs/view/#####/?refId=####\u0026trackingId=####\u0026trk=###\u0026lipi=####","JobPostAndDescriptionAlignment":{"CompanyTitle":{"Status":"ok","Reason":"only listed once"},"JobTitle":{"Status":"ok","Reason":"matches"},"RequiredSkills":{"Status":"poor","Reason":"much more required skills listed in description than on the job posting, and the main skill on job post is not listed in description"}}}]`
	result := processSubmittedApplicationsFile()
	if result != EXPECTED_RESULT {
		t.Fatalf("got: %s, but expected: %s", result, EXPECTED_RESULT)
	}
}

func TestMainProgram(t *testing.T) {
	runCmd(t, "rm", "-f", "./yeet")
	runCmd(t, "go", "build", "yeet.go")

	var EXIT_STATUS_0_STR = "exit status 0"
	var EXIT_STATUS_1_STR = "exit status 1"

	var tests = []struct {
		cmd                   string
		expectedOutputStr     string
		expectedExitStatusStr string
	}{
		{
			"./yeet",
			ErrIncorrectNumberOfArgs.Error(),
			EXIT_STATUS_1_STR,
		},
		{
			"./yeet proces",
			ErrInvalidArg.Error(),
			EXIT_STATUS_1_STR,
		},
		{
			"./yeet process",
			"success",
			EXIT_STATUS_0_STR,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("testname: %s", tt.cmd)
		t.Run(testname, func(t *testing.T) {
			cmdParts := strings.Split(tt.cmd, " ") // "./program arg1, ...argN"
			program := cmdParts[0]
			args := cmdParts[1:]
			out, err := exec.Command(program, args...).Output()
			outstr := string(out)

			// exec.Command:
			//  out is what is printed to stdout, err.Error() will show exit status 1, err is nil when exit status is 0
			if outstr != tt.expectedOutputStr {
				t.Fatalf("output is '%s', expected '%s'", outstr, tt.expectedOutputStr)
			}

			if err != nil {
				errstr := err.Error()
				if errstr != tt.expectedExitStatusStr {
					t.Fatalf("err is '%s', expected '%s'", errstr, tt.expectedExitStatusStr)
				}
			} else {
				if tt.expectedExitStatusStr == EXIT_STATUS_1_STR {
					t.Fatalf("exit status is %s, expected %s", "0", tt.expectedExitStatusStr)
				}
			}
		})
	}
}
