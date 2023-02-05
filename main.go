package main

import (
	"strings"

	"github.com/ahmed-deftoner/tf-generator/utils"
)

func main() {
	graph := utils.DecodeJSON("graph.json")
	utils.CreateDirs(graph.Nodes)
	//m := []string{}
	for k, n := range graph.Nodes {
		raw := strings.Split(k, "#")
		utils.CheckService(raw[0], raw[1], n)
		utils.EditProvider(n.Region)
	}
}
