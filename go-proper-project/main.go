package main

import (
    "fmt"
    "log"
)

func main() {
    log.SetFlags(0)
    log.SetPrefix("greetings: ")



    message, err := Hello("dfghjkkkklllllllllllllj")
    


    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
    
    
    
    log.SetPrefix("calculate: ")
    division, err2 := calculate(2, 7)
     if err2 != nil {
        log.Fatal(err2)
    }
    fmt.Println(division)
    
    log.SetPrefix("HelloFormatted: ")
    message2, err3 := HelloFormatted("Fathia");
    if err3 != nil {
        log.Fatal(err3)
    }
    fmt.Println(message2)
    
    
   log.SetPrefix("messages: ")
  names := []string{"fathua", "toba", "omotemmy"} 
  messages, err4 := hellos(names)
  if messages == nil{
  	log.Fatal(err4)
  }
  fmt.Println(messages);
    
    
}
