package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	pl := map[string]Person{}
	pl["test"] = Person{"John", "Doe", 42}

	fmt.Println(pl)
}

func toto(test string) (res string, err error) {
	return test, nil
}
