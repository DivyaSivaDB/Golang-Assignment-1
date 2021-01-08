package main

import(
	"fmt"
	"encoding/json"
)

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Email string
	Age int
	HeightInMeters float64
	IsMale bool
}

func main() {
	
	// define `Ram` struct
	Ram := Student{
		FirstName: "Ram",
		lastName: "Kumar",
		Age: 25,
		HeightInMeters: 1.75,
		IsMale: true,
	}

	// encode `Ram` as JSON
	RamJSON, _ := json.Marshal( Ram )

	// print JSON string
	fmt.Println( string(RamJSON) )
}
