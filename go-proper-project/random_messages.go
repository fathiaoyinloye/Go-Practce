package main
import (
	"fmt"
	"math/rand"
	//"errors"
)

funtion Hellos([] string name) map[string] string{
	messages := make (map[string]string)
	
	for count:= 0; count < len(name); count++{
	    message := fmt.Sprintf(randomFormat(), name[count])
		mesages[name] = message
	
	}
	return messages;


}
