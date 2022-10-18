package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}

var dog []struct {
	Name string
}{ "Sisu"}

func main() {
	walk(dog.Name)
}
