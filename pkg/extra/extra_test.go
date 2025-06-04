package extra

import "testing"

func TestSum(t *testing.T) {
	if Sum(2, 3) != 5 {
		t.Fatalf("sum")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(2, 3) != 6 {
		t.Fatalf("mul")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(5, 3) != 2 {
		t.Fatalf("sub")
	}
}

func TestDiv(t *testing.T) {
	if Div(6, 3) != 2 {
		t.Fatalf("div")
	}
	if Div(4, 0) != 0 {
		t.Fatalf("div zero")
	}
}
