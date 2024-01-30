package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Naam           string   `json:"naam"`
	Leerlingnummer string   `json:"Leerlingnummer"`
	Cijfers        []Cijfer `json:"Cijfers"`
}

type Cijfer struct {
	Vak    string  `json:"Vak"`
	Cijfer float64 `json:"Cijfer"`
}

func main() {
	// Read the JSON data from the file
	data, err := os.ReadFile("cijferlijst.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of Student structs
	var students []Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		log.Fatal(err)
	}

	// Print the contents of the students slice
	fmt.Println("Before adding a new student:")
	for _, student := range students {
		fmt.Println("Student: ", student.Naam)
	}

	// Prompt the user to enter a full name
	var Firstname, Lastname string
	fmt.Println("Voer voornaam:")
	fmt.Scanln(&Firstname)

	fmt.Println("Voer Achternaam:")
	fmt.Scanln(&Lastname)

	fullName := Firstname + " " + Lastname

	// Find the student with the matching full name
	found := false
	for _, student := range students {
		if fullName == student.Naam {
			fmt.Println("Student found:")
			fmt.Println("Name:", student.Naam)
			fmt.Println("Student Number:", student.Leerlingnummer)
			fmt.Println("Grades:")
			for _, grade := range student.Cijfers {
				fmt.Printf("%s: %.1f\n", grade.Vak, grade.Cijfer)
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Println("No student found with that name.")
	}
}
