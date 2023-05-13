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
	env := object.NewEnvironment()

	return Eval(program, env)
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

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 46;", 46},
		{"return 10; 46;", 10},
		{"return 46; 11; return 92;", 46},
		{"11; return 4 * 6; 92;", 24},
		{
			`
			if (46 > 13) {
				if (13 > 1) {
					return 31;
				}

				erturn 1;
			}
			`,
			31,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input       string
		expectedMsg string
	}{
		{
			"46 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"46 + true; 46;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}

				return 1;
			}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Log(evaluated)
			t.Log(tt.input)
			t.Errorf("no error object returned. Got %T (%+v)", evaluated, evaluated)

			continue
		}

		if errObj.Message != tt.expectedMsg {
			t.Errorf("wrong error message. Got %q, want %q", errObj.Message, tt.expectedMsg)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 46; a;", 46},
		{"let a = 46 * 2; a;", 92},
		{"let a = 46; let b = a; b;", 46},
		{"let a = 46; let b = a; let c = a + b + 5; c;", 97},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. Got %T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. Got %q", fn.Parameters[0])
	}

	expectedBody := "(x+2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. Got %q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(46);", 46},
		{"let identity = fn(x) { return x; }; identity(46);", 46},
		{"let double = fn(x) { x * 2; }; double(46);", 92},
		{"let add = fn(x, y) { x + y; }; add(46, 92);", 138},
		{"let add = fn(x, y) { x + y; }; add(46 + 92, add(46, 92));", 276},
		{"fn(x) { x; }(46)", 46},
		{
			`
			let newAdder = fn(x) { fn(y) { x + y; }; };
			let addTwo = newAdder(13);
			addTwo(46);
			`,
			59,
		},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
