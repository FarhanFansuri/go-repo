package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	// Array of Object

	var jsonStrings = `[
	{
		"name":"Ronaldo",
		"age":19
	},
		{
		"name":"Messi",
		"age":21
	}
	]`

	var footballers []Person
	var jsonFootbalers = []byte(jsonStrings)
	var err2 = json.Unmarshal(jsonFootbalers, &footballers)

	if err2 != nil {
		fmt.Println(err2.Error())
	}

	fmt.Println(footballers[0])

	// encode
	var ToJsonFootbaler = []Person{
		{
			Name: "Cristiano",
			Age:  21,
		},
		{
			Name: "Messi",
			Age:  21,
		},
	}
	jsonDataFootbaler, err3 := json.Marshal(ToJsonFootbaler)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	fmt.Println(string(jsonDataFootbaler))
	fmt.Println(reflect.TypeOf(jsonDataFootbaler))
}
