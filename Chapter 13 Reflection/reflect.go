package main

import "reflect"

// walk accepts an interface{} and func(input string). we are passing it a struct{} and
// func(input string) [an empty func with a string input declaration]
// so fn := func(input string) // as passed from TestWalk and as defined here
// how does fn(return "Chris") modify input.

/*
The test passes in a function that looks like this.  It's probably the confusing part.
func(input string) {
				got = append(got, input)
			}
*/
func walk(x interface{}, fn func(input string)) {

	// get concrete values for the interface, assign to val
	val := reflect.ValueOf(x)

	// for each field in val
	for i := 0; i < val.NumField(); i++ {
		
		//field is now chris or london depending on i

		// Set field to the val iterators currently indexed position
		field := val.Field(i)
		// how is below assining things to input? what is this doing?
		//test case below shoul now be fn("Chris") for i=0. is input being
		//redeclared/assigned inside the fn() call. I think im missing a key concept
		// around the value of input

		// Take the current field as a string and use it as a parameter to our anonymous function
		fn(field.String())

		// The function we passed in appends got to the given input, which should be the field we passed in when calling the function.
	}
}
