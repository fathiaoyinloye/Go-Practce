package main
import "fmt"

type Person struct {
	Name string
	Age int
}

func aboutMe(person Person) string{
	return fmt.Sprintf("My name is %s and I am %d years old\n", person.Name, person.Age)
	

}
	