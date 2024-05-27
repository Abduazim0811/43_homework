package factorial

import "testing"

func TestFactorial(t *testing.T){
	tests:=[]struct{
		input int
		expected int
	}{
		{0, 1},
        {1, 1},
        {2, 2},
        {3, 6},
        {5, 120},
	}
	
	for _,test:=range tests{
		resault:=Factorial(test.input)
		if resault!=test.expected{
			t.Errorf("Factorial(%d)=%d; want %d", test.input,resault,test.expected)
		}
	}
}
