package evaluator

import (
	"monkey-language/lexer"
	"monkey-language/object"
	"monkey-language/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-46", -46},
		{"-123", -123},
		{"5 + 5 + 5 + 5 - 46", -26},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"46 > 10", true},
		{"10 > 46", false},
		{"46 < 10", false},
		{"10 < 46", true},
		{"46 == 46", true},
		{"46 != 46", false},
		{"46 == 10", false},
		{"46 != 10", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(46 > 10) == true", true},
		{"(46 > 10) == false", false},
		{"(46 < 10) == true", false},
		{"(46 < 10) == false", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj is not an Integer. Got %T (%+v)", obj, obj)

		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. Got %d, want %d", result.Value, expected)

		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj is not an Integer. Got %T (%+v)", obj, obj)

		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. Got %t, want %t", result.Value, expected)

		return false
	}

	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!46", false},
		{"!!true", true},
		{"!!false", false},
		{"!!46", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input  string
		expect interface{}
	}{
		{"if (true) { 46 }", 46},
		{"if (false) { 46 }", nil},
		{"if (1) { 46 }", 46},
		{"if (1 < 2) { 46 }", 46},
		{"if (1 > 2) { 46 }", nil},
		{"if (1 > 2) { 46 } else { 92 }", 92},
		{"if (1 < 2) { 46 } else { 92 }", 46},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expect.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("obj is not NULL. Got %T (%+v)", obj, obj)

		return false
	}

	return true
}
