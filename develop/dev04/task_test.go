package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	test := struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		name:  "Test groupAnagrams",
		input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
		expected: map[string][]string{
			"пятак":  {"пятак", "пятка", "тяпка"},
			"листок": {"листок", "слиток", "столик"},
		},
	}

	t.Run(test.name, func(t *testing.T) {
		result := groupAnagrams(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test case %s: expected \"%s\", got \"%s\"", test.name, test.expected, result)
			return
		}
	})
}
