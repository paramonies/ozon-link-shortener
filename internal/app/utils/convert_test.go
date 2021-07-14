package utils

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUtils_Convert(t *testing.T) {
	type inputArgs struct {
		Url string
		id  int
	}

	tests := []struct {
		name           string
		input          inputArgs
		expectedResult string
	}{
		{
			name: "Ok",
			input: inputArgs{
				Url: "http://test.ru",
				id:  1,
			},
			expectedResult: "mFU1iiuKLp",
		},
		{
			name: "Ok",
			input: inputArgs{
				Url: "http://test1.ru",
				id:  1,
			},
			expectedResult: "jwRXfn0ns7",
		},
	}

	for _, test := range tests {
		result := Convert(test.input.id, test.input.Url)
		// t.Logf("!!! %s expected: %v\ngot: %v", test.name, test.expectedResult, result)
		assert.Equal(t, result, test.expectedResult)
	}
}

func TestUtils_IsShortValid(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{
			name:           "Ok - valid",
			input:          "mFU1iiuKLp",
			expectedResult: true,
		},
		{
			name:           "Failure - Not valid",
			input:          "mFU1iiuKLpP",
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsShortValid(test.input)
			assert.Equal(t, result, test.expectedResult)
		})
	}
}
