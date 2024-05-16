package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Course struct {
	Name     string `json:"coursename"`
	Password string `json:"-"`
	Price    int
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	courses := []Course{
		{"AI/ML", "abc123", 2000, []string{"ai", "genai"}},
		{"MERN", "pqr123", 4000, nil},
		{"Devops", "abc123", 1000, []string{"docker", "k8s"}},
	}
	// body, err := json.MarshalIndent(courses, "", "\t")
	body, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	var stringbuilder strings.Builder
	stringbuilder.Write(body)
	fmt.Println(stringbuilder.String())

	stringbuilder.Reset()

}
