package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with string field",
			struct {
				Name   string
				Second string
			}{"hello", "world"},
			[]string{"hello", "world"},
		},
		{
			"struct with int field",
			struct {
				Name int
			}{123},
			[]string{"123"},
		},
		{
			"struct with string field",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"hello", struct {
				Age  int
				City string
			}{20, "wuhan"}},
			[]string{"hello", "20", "wuhan"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("want %v, got %v", test.ExpectedCalls, got)
			}
		})
	}
}
