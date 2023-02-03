package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Node struct {
	Region     string            `json:"region"`
	Properties map[string]string `json:"properties"`
}

type Graph struct {
	Adjacency map[string][]string `json:"adjacency"`
	Nodes     map[string]Node     `json:"nodes"`
	Edges     map[string]Edge     `json:"edges"`
}

func DecodeJSON(fileNmae string) Graph {
	jsonFile, err := os.Open(fileNmae)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result Graph
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result.Adjacency["ALB"][0])
	return result
}
