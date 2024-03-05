package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/L0RAX-26/fizzbuzz"
)

func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("the value %v should not have returned true", 1)
		t.Fail()
	}

	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("the value %v should have returned true", 3)
		t.Fail()
	}

	if result != "fizz" {
		t.Log("Result should have been fizz")
		t.Fail()
	}
}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	// output:fizz
}
