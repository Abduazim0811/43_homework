package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests:=[]struct{
		input string
		expected string
	}{
		{"", ""},
        {"a", "a"},
        {"ba", "ab"},
        {"hello", "olleh"},
        {"world", "dlrow"},
	}

	for _,test:=range tests{
		resault:=Reverse(test.input)
		if resault != test.expected{
			t.Errorf("Reverse(%q)=%q; want %q", test.input,resault,test.expected)
		}
	}
}
