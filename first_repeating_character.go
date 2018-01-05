// module to return first repeating character in a string

package main

import "fmt"

// declaring a dict || hash type
var test map[string] int


func find_repeated_numbers(s string) {

    // print statement and formatting
    fmt.Printf("input type %T and value is %v\n", s, s)
    

    // initializing the map of given type
    test = make(map[string] int)

    // for loop declaration
    for i := 0; i < len(s); i++ {

	// current character pointer
	current_character := string(s[i])
	
	// if statements in Go can include both an initialization statement and condition
	if value, ok := test[current_character]; ok {
	    fmt.Printf("%v repeated more than %v", current_character, value)
	    return 
	}
	
	// updating the dict
	test[current_character] = 1
	
    }
    
    fmt.Println("None")
    return 

	
}

func main() {     
    find_repeated_numbers("DBCAL")
}
