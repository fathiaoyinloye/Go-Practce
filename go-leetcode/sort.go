package main
import (
	"fmt"
)
func sortSlice(slice []int) []int {
	for index := 0; index < len(slice); index++ {
		smallest := slice[index]

		for index2 := index + 1; index2 < len(slice); index2++ {
			if smallest > slice[index2] {
				smallest = slice[index2]
			}
			if smallest < slice[index] {
				slice[index2] = slice[index];
				slice[index] = smallest;

		}
	}

}
	fmt.Println(slice)
	return slice
}