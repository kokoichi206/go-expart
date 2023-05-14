package evaluator

import (
	"monkey-language/object"
	"testing"
)

func TestBuiltinFunctionsLen(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world!")`, 12},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}

func TestBuiltinFunctionString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      string
	}{
		{`string(4)`, "4", ""},
		{`string("1")`, "", "argument to `string` not supported, got STRING"},
		{`string("one", "two")`, "", "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		if tt.err == "" {
			testStringObject(t, evaluated, tt.expected)
		} else {
			errObj, ok := evaluated.(*object.Error)
			t.Log(evaluated)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Message != tt.err {
				t.Errorf("wrong error message. expected=%q, got=%q", tt.err, errObj.Message)
			}
		}
	}
}
