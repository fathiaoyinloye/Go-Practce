package main

import (
    "testing"
	"slices"
)


func TestSortNumbers(t *testing.T) {
    want := [] int{1,2,3,4,5,6,7}
	nums := [] int {5,7,1,3,6,2,4}
    sorted_num := sortNumbers(nums)
    if !slices.Equal(sorted_num, want) {
    	t.Errorf(`sortNumbers(nums)= %v, want match for %v`, sorted_num, want)

    
    }
}
