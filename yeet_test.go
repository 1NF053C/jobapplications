/*
 * yeet_test.go
 *
 * Test code for yeet.go
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
	"os/exec"
	"strings"
	"testing"
)

func TestMainProgram(t *testing.T) {
	runCmd(t, "rm", "-f", "./yeet")
	runCmd(t, "go", "build")

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

func runCmd(t *testing.T, cmdParts ...string) {
	if err := exec.Command(cmdParts[0], cmdParts[1:]...).Run(); err != nil {
		t.Fatalf("cmd failed, exiting tests. Err: %s", err.Error())
	}
}
