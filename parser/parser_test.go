package parser

import (
	"fmt"
	"testing"

	"github.com/bamchoh/json_parser/lexer"
)

func TestParser(t *testing.T) {
	input := `{
		"string": 1234,
		"string2": 2345,
		"true": true,
		"false": false,
		"null": null,
	}`

	l := lexer.New(input)
	p := New(l)
	data := p.Parse()

	obj, ok := data.(map[string]interface{})
	if !ok {
		t.Fatalf("Root type is not ast.Object got=%T", data)
	}

	expected := []struct {
		Key  string
		Val  interface{}
		Type string
	}{
		{"string", float64(1234), "NUMBER"},
		{"string2", float64(2345), "NUMBER"},
		{"true", nil, "TRUE"},
		{"false", nil, "FALSE"},
		{"null", nil, "NULL"},
	}

	for _, tt := range expected {
		val, ok := obj[tt.Key]
		if !ok {
			t.Fatalf("'string' key does not exist in a object")
		}

		switch tt.Type {
		case "NUMBER":
			checkNumber(t, val, tt.Val.(float64))
		case "TRUE":
			checkBool(t, val, true)
		case "FALSE":
			checkBool(t, val, false)
		case "NULL":
			checkNull(t, val)
		}
	}
}

func TestParserArray(t *testing.T) {
	input := `[
		1234,
		2345,
		"string",
		true,
		false,
		null,
	]`

	l := lexer.New(input)
	p := New(l)
	data := p.Parse()

	obj, ok := data.([]interface{})
	if !ok {
		t.Fatalf("Root type is not ast.Object got=%T", data)
	}

	expected := []struct {
		Val  interface{}
		Type string
	}{
		{float64(1234), "NUMBER"},
		{float64(2345), "NUMBER"},
		{"string", "STRING"},
		{true, "TRUE"},
		{false, "FALSE"},
		{nil, "NULL"},
	}

	for i, tt := range expected {
		val := obj[i]

		fmt.Println(tt, val)

		switch tt.Val.(type) {
		case float64:
			checkNumber(t, val, tt.Val.(float64))
		case string:
			checkString(t, val, tt.Val.(string))
		case bool:
			checkBool(t, val, tt.Val.(bool))
		case nil:
			checkNull(t, val)
		}
	}
}

func TestParserComplex(t *testing.T) {
	input := `[
		{"string":1234},
		{"number":[1,2,3]},
		{"boolean":[true, true, false, false]},
		{"null": null},
		{"array of array": [[1,2],[3,4]]}
	]`

	l := lexer.New(input)
	p := New(l)
	data := p.Parse()

	obj, ok := data.([]interface{})
	fmt.Println(data)
	if !ok {
		t.Fatalf("Root type is not []interface{} got=%T", data)
	}

	expected := []map[string]interface{}{
		{"string": float64(1234)},
		{"number": []float64{1, 2, 3}},
		{"boolean": []bool{true, true, false, false}},
		{"null": nil},
		{"array of array": [][]int{[]int{1, 2}, []int{3, 4}}},
	}

	if len(obj) != len(expected) {
		t.Fatalf("len is different between obj and expected len(obj)=%v, len(expected)=%v", len(obj), len(expected))
	}

	for i, map2 := range expected {
		map1, ok := obj[i].(map[string]interface{})
		if !ok {
			t.Fatalf("obj[i] is not map[string]interface{} got=%T", obj[i])
		}
		for k, cmp2 := range map2 {
			cmp1, ok := map1[k]
			if !ok {
				t.Fatalf("map1 has no key '%v'", k)
			}
			if fmt.Sprintf("%v", cmp1) != fmt.Sprintf("%v", cmp2) {
				t.Fatalf("obj is not equal to expected, \nobj=%v\nexp=%v", cmp1, cmp2)
			}
		}
	}
}

func checkNumber(t *testing.T, val interface{}, expected float64) {
	if num, ok := val.(float64); !ok {
		t.Fatalf("Type of 'val' is not float64 got=%T", val)
	} else {
		if num != expected {
			t.Fatalf("'val' is not %v got=%v", expected, num)
		}
	}
}

func checkString(t *testing.T, val interface{}, expected string) {
	if str, ok := val.(string); !ok {
		t.Fatalf("Type of 'val' is not string got=%T", val)
	} else {
		if str != expected {
			t.Fatalf("'val' is not %v got=%v", expected, str)
		}
	}
}

func checkBool(t *testing.T, val interface{}, expected bool) {
	if v, ok := val.(bool); !ok {
		t.Fatalf("Type of 'val' is not *ast.True got=%T", val)
	} else {
		if v != expected {
			t.Fatalf("'val' is not %v got=%v", expected, v)
		}
	}
}

func checkNull(t *testing.T, val interface{}) {
	if val != nil {
		t.Fatalf("Type of 'val' is not nil got=%T", val)
	}

}
