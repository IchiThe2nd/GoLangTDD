package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 6)
	fmt.Println(repeated)
	// Output: cccccc
}
