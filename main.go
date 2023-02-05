package main

import (
	"strings"

	"github.com/ahmed-deftoner/tf-generator/utils"
)

func main() {
	graph := utils.DecodeJSON("graph.json")
	utils.CreateDirs(graph.Nodes)
	m := []string{}
	for k, n := range graph.Nodes {
		if !utils.Contains(m, n.Region) {
			m = append(m, n.Region)
		}
		raw := strings.Split(k, "#")
		utils.CheckService(raw[0], raw[1], n)
	}
	for _, v := range m {
		utils.EditProvider(v)
	}
}
