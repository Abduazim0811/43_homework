package wordcount

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"hello world", map[string]int{"hello": 1, "world": 1}},
		{"", map[string]int{}},
		{"a b a b c 2 f", map[string]int{"a": 2, "b": 2, "c": 1, "2": 1, "f": 1}},
		{"test test test", map[string]int{"test": 3}},
		{"Go is fun is Go", map[string]int{"Go": 2, "is": 2, "fun": 1}},
	}
	for _, test := range tests {
		resault := Wordcount(test.input)
		if !reflect.DeepEqual(resault, test.expected) {
			t.Errorf("WordCount(%q) = %v; want %v", test.input, resault, test.expected)
		}
	}
}
