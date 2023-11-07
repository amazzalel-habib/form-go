package formgo

import (
	"net/url"
	"testing"
)

type TestData struct {
	StringField string `form:"str_field"`
	IntField    int    `form:"int_field"`
}

func TestEncodeFormValues(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected url.Values
	}{
		{
			name:     "ValidData",
			input:    TestData{StringField: "abc", IntField: 123},
			expected: url.Values{"str_field": []string{"abc"}, "int_field": []string{"123"}},
		},
		{
			name:     "EmptyString",
			input:    TestData{StringField: "", IntField: 123},
			expected: url.Values{"int_field": []string{"123"}},
		},
		{
			name:     "ZeroInt",
			input:    TestData{StringField: "abc", IntField: 0},
			expected: url.Values{"str_field": []string{"abc"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := EncodeFormValues(tc.input)
			if err != nil {
				t.Errorf("EncodeFormValues returned an error: %v", err)
			}

			if !urlValuesEqual(values, tc.expected) {
				t.Errorf("EncodeFormValues result mismatch. Got %v, expected %v", values, tc.expected)
			}
		})
	}
}

func urlValuesEqual(a, b url.Values) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valuesA := range a {
		valuesB, ok := b[key]
		if !ok {
			return false
		}
		if len(valuesA) != len(valuesB) {
			return false
		}
		for i, valueA := range valuesA {
			if valueA != valuesB[i] {
				return false
			}
		}
	}
	return true
}
