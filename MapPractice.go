package main
import (
    "fmt"
    //"maps"
)
func main() {
	//To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)
	//Set key/value pairs using typical name[key] = val syntax.

  	m["k1"] = 7
  	m["k2"] = 13
	//Printing a map with e.g. fmt.Println will show all of its key/value pairs.

    	fmt.Println("map:", m)
    	
    	v1 := m["k1"]
    	fmt.Println("v1:", v1)
    	
    	v3 := m["k3"]
    	fmt.Println("v3:", v3)
	//The builtin len returns the number of key/value pairs when called on a map.

    	fmt.Println("len:", len(m))
    	
    	delete(m, "k2")
    	fmt.Println("map:", m)
	//To remove all key/value pairs from a map, use the clear builtin.

    	//clear(m)
    	//fmt.Println("map:", m)
	//The optional second return value when getting a value from a map indicates if the key was 			present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.

   	 p, prs := m["k1"]
   	 
    	fmt.Println("prs:", prs, "p",p)
    
    	//n := map[string]int{"foo": 1, "bar": 2}
    	//fmt.Println("map:", n)
	//The maps package contains a number of useful utility functions for maps.

    	//n2 := map[string]int{"foo": 7, "bar": 2}
    	//if maps.Equal(n, n2) {
        //	fmt.Println("n == n2")
   	 //}
    	
}
