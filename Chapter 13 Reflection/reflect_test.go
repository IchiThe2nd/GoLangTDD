package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	// cases is a slice of structs with slice 0 being the test named struct with on string. Input is struct{Name field} == chris and .ExpectedCalls == []string with chris @ [0]
	// it has a struct interface Name string .Name is chris and .ExepectedCalls is chris
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct A with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with 2 string field ",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				// func(input string) passed to walk as second argument
				//final line of walk is fn("chris") which assigned "chris" to input
				//
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
