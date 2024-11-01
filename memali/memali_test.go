package memali_test

import (
	"memali/memali"
	"slices"
	"testing"
)

func TestParseCode(t *testing.T) {
	testCases := []struct {
		name             string
		code             string
		expectedElements []string
	}{
		{
			name: "flat code",
			code: `type s struct {
				a int
				c []string
				d bool
			}`,
			expectedElements: []string{"a int", "c []string", "d bool"},
		},
		{
			name: "flat code with comments",
			code: `// Some stuct 
				type s struct{
				// Some field
				a int
				// And another field
				// and long comm
				c int
			}`,
			expectedElements: []string{"a int", "c int"},
		},
		{
			name: "flat code with comments and spaces",
			code: `// Some struct
				type s struct {


				// Comment
				a bool


				// Comment
				
				// long

				q string			
	}`,
			expectedElements: []string{"a bool", "q string"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fields, err := memali.FindField(tc.code)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}
			if !slices.Equal(fields, tc.expectedElements) {
				t.Errorf("got: %v\nexpected: %v", fields, tc.expectedElements)
			}
		})
	}

}
