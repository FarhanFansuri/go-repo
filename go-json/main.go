package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:name`
	Age  int    `json:age`
}

func main() {
	//JSON
	jsonString := `{
	"name" : "John",
	"age":21
	}`

	var jsonData = []byte(jsonString)

	var john Person
	var err = json.Unmarshal(jsonData, &john)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(john)

}
