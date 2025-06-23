package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name       string	`json:"Name"`
	Occupation string	`json:"Occupation"`
	Age        int		`json:"Age"`
}

func main() {
	Trevor := Person{Name: "Trevor", Occupation: "SRE", Age: 23}
	jsonTrevor, err := json.Marshal(Trevor)
	if err != nil {
		log.Fatal("no good json")
	}

	fmt.Printf("%s\n", jsonTrevor)
}
