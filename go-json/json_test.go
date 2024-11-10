package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {

	data := Customer{
		Firstname: "Farhan",
		Lastname:  "Fansuri",
	}
	logJson(data)

}

type Customer struct {
	Firstname string
	Lastname  string
}
