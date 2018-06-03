package fizzbuzz

import (
	"testing"
)

func Test_input_1_should_say_1(t *testing.T) {
	input := 1
	expected := "1"

	actual := Say(input)

	Assert(t, actual, expected)
}

func Test_input_2_should_say_2(t *testing.T) {
	input := 2
	expected := "2"

	actual := Say(input)

	Assert(t, actual, expected)
}

func Test_input_3_should_say_Fizz(t *testing.T) {
	input := 3
	expected := "Fizz"

	actual := Say(input)

	Assert(t, actual, expected)
}

func Test_input_5_should_say_Buzz(t *testing.T) {
	input := 5
	expected := "Buzz"

	actual := Say(input)

	Assert(t, actual, expected)
}

func Test_input_6_should_say_Fizz(t *testing.T) {
	input := 6
	expected := "Fizz"

	actual := Say(input)

	Assert(t, actual, expected)
}

func Test_input_10_should_say_Buzz(t *testing.T) {
	input := 10
	expected := "Buzz"

	actual := Say(input)

	Assert(t, actual, expected)
}
func Test_input_15_should_say_FizzBuzz(t *testing.T) {
	input := 15
	expected := "FizzBuzz"

	actual := Say(input)

	Assert(t, actual, expected)
}

func Assert(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Expected %s but %s", expected, actual)
	}
}
