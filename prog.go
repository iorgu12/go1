package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	p := Person{Name: "Your Name"}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
