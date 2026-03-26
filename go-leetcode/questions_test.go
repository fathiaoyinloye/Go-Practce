package main

import (
    "testing"
	"slices"
)

func TestLengthfWords(t *testing.T) {
    	got := lengthOfLastWord("I am a girl")
    	want := 4
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}

    

    }


	func TestSortSlice(t *testing.T) {
		got := sortSlice([]int{5, 2, 9, 1, 5, 6})
		want := []int{1, 2, 5, 5, 6, 9}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
