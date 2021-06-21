package main

import "testing"

func TestInvalidOperationLastPlus(t *testing.T) {
	var args = []string{"1/3", "+", "3/7", "+"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "+", "3/7", "+"]) expected error message: `)
	}
}

func TestInvalidOperationLastMinus(t *testing.T) {
	var args = []string{"1/3", "+", "3/7", "-"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "+", "3/7", "-"]) expected error message: `)
	}
}

func TestInvalidOperationLastMultiply(t *testing.T) {
	var args = []string{"1/3", "+", "3/7", "*"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "+", "3/7", "*"]) expected error message: `)
	}
}

func TestInvalidOperationLastDivide(t *testing.T) {
	var args = []string{"1/3", "+", "3/7", "/"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "+", "3/7", "/"]) expected error message: `)
	}
}

func TestInvalidOperationDoublePlus(t *testing.T) {
	var args = []string{"1/3", "+", "+", "3/7"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "+", "+", "3/7"]) expected error message: `)
	}
}

func TestInvalidOperationDoubleMinus(t *testing.T) {
	var args = []string{"1/3", "-", "-", "3/7"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "-", "-", "3/7"]) expected error message: `)
	}
}

func TestInvalidOperationDoubleMultiply(t *testing.T) {
	var args = []string{"1/3", "*", "*", "3/7"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "*", "*", "3/7"]) expected error message: `)
	}
}

func TestInvalidOperationDoubleDivide(t *testing.T) {
	var args = []string{"1/3", "/", "/", "3/7"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "/", "/", "3/7"]) expected error message: `)
	}
}

func TestInvalidOperationDoubleMixed(t *testing.T) {
	var args = []string{"1/3", "+", "/", "3/7"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "+", "/", "3/7"]) expected error message: `)
	}
}

func TestInvalidOperationWrongCharacter(t *testing.T) {
	var args = []string{"1/3", "&", "3/7"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["1/3", "&", "3/7"]) expected error message: `)
	}
}

func TestInvalidOperationFirstCharacterOperator(t *testing.T) {
	var args = []string{"+", "1/3", "3/7"}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation(["+", "1/3", "3/7"]) expected error message: `)
	}
}

func TestNoArguments(t *testing.T) {
	var args = []string{}
	_, _, err := calculateOperation(args)
	if err == nil {
		t.Fatalf(`calculateOperation([]) expected error message: `)
	}
}

func TestValidFractionSum(t *testing.T) {
	var args = []string{"1/2", "+", "1/2"}
	sum, _, err := calculateOperation(args)
	if err != nil {
		t.Fatalf(`calculateOperation([]) expected error message: `)
	}

	if sum != 1 {
		t.Fatalf(`calculateOperation([]) expected : 1.0 got: %0.2f `, sum)
	}

}
