package main

import (
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_output(t *testing.T) {

	// drwxr-xr-x のモードの出力。
	// -l が正しくできていることの判断に使用。
	const modeRegex = "[d-][rwx-]{9}.*"

	testCases := []struct {
		name          string
		params        Params
		expectedExit  int
		checkResponse func(t *testing.T, result string)
	}{
		{
			name: "WithNoDirectoryName",
			params: Params{
				IsHelp:  false,
				IsColor: false,
				// When no Arg is passed, the '.' is set to Name.
				ShowList: false,
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
				require.Contains(t, result, "main.go")
				require.Contains(t, result, "examples")
				require.NotContains(t, result, ".github")
				mode, _ := regexp.MatchString(modeRegex, result)
				require.False(t, mode)
			},
		},
		{
			name: "WithNoDirectoryNameAndAllOption",
			params: Params{
				IsHelp:     false,
				IsColor:    false,
				ShowHidden: true,
				// When no Arg is passed, the '.' is set to Name.
				ShowList: false,
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
				require.Contains(t, result, "main.go")
				require.Contains(t, result, "examples")
				require.Contains(t, result, ".github")
				mode, _ := regexp.MatchString(modeRegex, result)
				require.False(t, mode)
			},
		},
		{
			name: "WithSpecificDirectoryName",
			params: Params{
				IsHelp:   false,
				IsColor:  false,
				Args:     []string{"examples"},
				ShowList: false,
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
				require.Contains(t, result, "apple")
				require.Contains(t, result, "banana")
				require.Contains(t, result, "dir_ex")
				require.NotContains(t, result, ".secret")
				mode, _ := regexp.MatchString(modeRegex, result)
				require.False(t, mode)
			},
		},
		{
			name: "WithSpecificDirectoryNameAndAllOption",
			params: Params{
				IsHelp:     false,
				IsColor:    false,
				ShowHidden: true,
				ShowList:   false,
				Args:       []string{"examples"},
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
				require.Contains(t, result, "apple")
				require.Contains(t, result, "banana")
				require.Contains(t, result, "dir_ex")
				require.Contains(t, result, ".secret")

			},
		},
		{
			name: "MultipleDirectoryNames",
			params: Params{
				IsHelp:   false,
				IsColor:  false,
				ShowList: false,
				Args:     []string{"examples", "examples0", "examples2"},
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
			},
		},
		{
			name: "WithHelpOption",
			params: Params{
				IsHelp:   true,
				IsColor:  false,
				ShowList: false,
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
				require.Contains(t, result, "ls – list directory contents")
				require.Contains(t, result, "The following options are available:")
				mode, _ := regexp.MatchString(modeRegex, result)
				require.False(t, mode)
			},
		},
		{
			name: "WithWrongDirectoryName",
			params: Params{
				IsHelp:   false,
				IsColor:  false,
				ShowList: false,
				Args:     []string{"not_existing_directory"},
			},
			expectedExit: 1,
			checkResponse: func(t *testing.T, result string) {
				require.Contains(t, result, "No such file or directory")
				mode, _ := regexp.MatchString(modeRegex, result)
				require.False(t, mode)
			},
		},
		{
			name: "WithColor",
			params: Params{
				IsHelp:   false,
				IsColor:  true,
				ShowList: false,
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
				mode, _ := regexp.MatchString(modeRegex, result)
				require.False(t, mode)
			},
		},
		{
			name: "WithListOption",
			params: Params{
				IsHelp:   false,
				IsColor:  false,
				ShowList: true,
			},
			expectedExit: 0,
			checkResponse: func(t *testing.T, result string) {
				mode, _ := regexp.MatchString(modeRegex, result)
				require.True(t, mode)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			oldExit := osExit
			defer func() {
				osExit = oldExit
			}()
			t.Log(tc.params.Args)
			t.Log(len(tc.params.Args))

			var statusCode int
			exit := func(code int) {
				statusCode = code
			}
			osExit = exit

			// Act
			result := GetStringFromStdOutput(execute, tc.params)

			// Assert
			t.Log(result)
			tc.checkResponse(t, result)
			require.Equal(t, tc.expectedExit, statusCode)
		})
	}
}

// Get Standard Outputs for a function call!
func GetStringFromStdOutput(fun func(Params), params Params) string {
	stdOut := os.Stdout

	defer func() {
		os.Stdout = stdOut
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w

	fun(params)

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	return output
}
