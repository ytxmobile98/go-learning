package main_test

import (
	"testing"

	json2 "example.com/json-not-escape-html/json"
	"gotest.tools/v3/assert"
)

type TestData struct {
	Data any
	ExpectedResults
}

type ExpectedResults struct {
	Escaped   string
	Unescaped string
}

func check(t *testing.T, data *any, escapeHTML bool, want *string) {
	result, err := json2.MarshalJSON(*data, escapeHTML)
	assert.NilError(t, err)
	assert.Equal(t, string(result), *want)
}

func TestMarshalJSON(t *testing.T) {
	tests := []TestData{
		{
			Data: map[string]interface{}{
				"a": "<>&",
			},
			ExpectedResults: ExpectedResults{
				Escaped:   `{"a":"\u003c\u003e\u0026"}`,
				Unescaped: `{"a":"<>&"}`,
			},
		},
	}

	for _, test := range tests {
		t.Run("MarshalJSON", func(t *testing.T) {
			check(t, &test.Data, true, &test.Escaped)
			check(t, &test.Data, false, &test.Unescaped)
		})
	}
}
