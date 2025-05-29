package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	testBinary := "test-ad-wc"
	build := exec.Command("go", "build", "-o", testBinary)
	if err := build.Run(); err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}

	defer os.Remove(testBinary)

	testCases := []struct {
		name	string
		args 	[]string
		expectedOutput string
		expectedError string
	} {
		{
			name:	"No parameters",
			args: 	[]string{},
			expectedError: "Error: exactyle one parameter required",
		},
		{
			name:	"Open non existing file",
			args: 	[]string{"-c", "nofile.txt"},
			expectedError: "Error when trying to open file: nofile.txt",
		},
		{
			name:	"With c parameter",
			args: 	[]string{"-c", "test.txt"},
			expectedOutput: "342190 test.txt",
		},
		{
			name:	"With full parameter",
			args: 	[]string{"-lcwm", "test.txt"},
			expectedOutput: "7145 58164 339292 342190 test.txt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("./" + testBinary, tc.args...)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			
			err := cmd.Run()

			// Check expected error
			if tc.expectedError != "" {
				if err == nil || !strings.Contains(stderr.String() + stdout.String(), tc.expectedError) {
					t.Errorf("Expected error containing % q, got: %v\nStdout: %s\nStderr: %s", tc.expectedError, err, stdout.String(), stderr.String())
				}
				return
			}

			// Check expected output
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !strings.Contains(stdout.String(), tc.expectedOutput) {
				t.Errorf("Expected output containing %q, got %q", tc.expectedOutput, stdout.String() + stderr.String())
			}
		})
	}
}