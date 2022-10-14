package main

import "reflect"

// walk accepts an interface{} and func(input string). we are passing it a struct{} and
// func(input string) [an empty func with a string input declaration]
// so fn := func(input string) // as passed from TestWalk and as defined here
// how does fn(return "Chris") modify input.
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		//field is now chris or london depending on i
		field := val.Field(i)
		// how is below assining things to input? what is this doing?
		//test case below shoul now be fn("Chris") for i=0. is input being
		//redeclared/assigned inside the fn() call. I think im missing a key concept
		// around the value of input
		fn(field.String())
	}
}
