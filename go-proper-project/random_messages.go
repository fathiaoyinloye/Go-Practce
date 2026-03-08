package main
import (
	"fmt"
	"errors"
)

func hellos(names []string) (map[string] string, error){
	messages := make (map[string]string)
	
    	for _, name := range names {
    	if name == ""{
    		return nil , errors.New("Name cannot be an empty string")
    	
    	}
    	
	    message := fmt.Sprintf(randomFormat(), name)
		messages[name] = message
	
	}
	return messages, nil

}
