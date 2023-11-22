package src

import (
	"testing"
)

func ans(input string) any {
	tokens, _ := Tokenize(input)
	st, _ := Parse(tokens)
	ans, _ := Interpret(st)
	return ans
}

func TestAddition(t *testing.T) {
	want := 12.0
	got := ans("6 + 6")
	if want != got {
		t.Fail()
	}

	want = 0.0
	got = ans("6 + -6")
	if want != got {
		t.Fail()
	}

	want = 1.0 
	got = ans("1/3 + 2/3")
	if want != got {
		t.Fail()
	}
}

func TestPemdas(t *testing.T) {
	want := -3.0
	got := ans("1 + 2 - 6")
	if want != got {
		t.Fail()
	}

	want = 4
	got = ans("1 + 6/2") 
	if want != got {
		t.Fail()
	}

	want = (1 + 6) / 2.0
	got = ans("(1 + 6) / 2")
	if want != got {
		t.Fail()
	}

	want = 64
	got = ans("2^3*8")
	if want != got {
		t.Fail()
	}

	want = 16777216.0
	got = ans("2^(3*8)")
	if want != got {
		t.Fail()
	}

	want = 8
	got = ans("(-44 + 46)^3")
	if want != got {
		t.Fail()
	}

	want = 5.0
	got = ans("2+3-4*3/5^2^(9-1)")
	if want != got {
		t.Fail()
	}

	want2 := true
	got = ans("6 == 6")
	if want2 != got {
		t.Fail()
	}

	want2 = false
	got = ans("6 == 2")
	if want2 != got {
		t.Fail()
	}

	want2 = true
	got = ans("9 == (18 / 2)")
	if want2 != got {
		t.Fail()
	}
}

func TestBoolLiteral(t *testing.T) {
	want := true
	got := ans("false == (true == false)")
	if want != got {
		t.Fail()
	}
}

func TestComparision(t *testing.T) {
	want := false
	got := ans("6 > 23")
	if want != got {
		t.Fail()
	}

	want = true
	got = ans("5 >= 5")
	if want != got {
		t.Fail()
	}

	want = true
	got = ans("6 <= 6")
	if want != got {
		t.Fail()
	}
}
