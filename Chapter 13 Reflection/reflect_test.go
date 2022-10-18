package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pouintters to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slicese dices",
			[]Profile{
				{33, "London"},
				{34, "Reyk"},
			},
			[]string{"London", "Reyk"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reyk"},
			},
			[]string{"London", "Reyk"},
		},
		{
			"maps",
			map[string]string{
				"foo": "bar",
				"baz": "boz",
			},
			[]string{"bar", "boz"},
		},
	}

	for _, test := range cases {
		t.Run("with Maps", func(t *testing.T) {
			aMap := map[string]string{
				"foo": "bar",
				"baz": "boz",
			}
			var got []string
			walk(aMap, func(input string) {
				got = append(got, input)
			})
			assertContains(t, got, "bar")
			assertContains(t, got, "boz")
		})

		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				//fmt.Println(input)

				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})

		t.Run("with Channels", func(t *testing.T) {
			aChannel := make(chan Profile)
			go func() {
				aChannel <- Profile{33, "Berlin"}
				aChannel <- Profile{34, "Kato"}
				close(aChannel)
			}()

			var got []string
			want := []string{"Berlin", "Kato"}

			walk(aChannel, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})

		t.Run("with function", func(t *testing.T) {
			aFunction := func() (Profile, Profile) {
				return Profile{33, "Berlin"}, Profile{34, "Kato"}
			}
			var got []string
			want := []string{"Berlin", "Kato"}

			walk(aFunction, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v wanted %v", got, want)
			}

		})

	}
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didnt", haystack, needle)
	}

}
