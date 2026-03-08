package greetings

import (
    "testing"
    "regexp"
)


func TestHelloName(t *testing.T) {
    want := "Hi, Gladys. Welcome!"
    msg, err := Hello("Gladys")
    if msg != want || err != nil{
    	t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)

    
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
