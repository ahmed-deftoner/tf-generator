package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ahmed-deftoner/tf-generator/components"
	"github.com/ahmed-deftoner/tf-generator/utils"
)

func createRDS(node map[string]utils.Node, id string, region string) {
	src := "tf/" + region + "/rds-" + id + ".tf"
	txtFile, err := os.Open(src)
	key := "RDS#" + id
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened rds.tf")
	defer txtFile.Close()
	body, err := ioutil.ReadAll(txtFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	raw := strings.Split(string(body), "=")
	var final string
	final += raw[0][2:] + "= " + node[key].Properties["allocated_storage"]
	final += raw[1] + "= \"" + node[key].Properties["db_name"] + "\""
	final += raw[2] + "= \"" + node[key].Properties["engine"] + "\""
	final += raw[3] + "= \"" + node[key].Properties["engine_version"] + "\""
	final += raw[4] + "= \"" + node[key].Properties["instance_class"] + "\""
	final += raw[5] + "= \"" + node[key].Properties["username"] + "\""
	final += raw[6] + "= \"" + node[key].Properties["password"] + "\""
	final += raw[7] + "= true" + raw[8][:len(raw[8])-2]

	ioutil.WriteFile(src, []byte(final), 0644)
}

func main() {
	graph := utils.DecodeJSON("graph.json")
	utils.CreateDirs(graph.Nodes)
	for k, n := range graph.Nodes {
		raw := strings.Split(k, "#")
		if raw[0] == "DynamoDB" {
			if raw[1] != "" {
				utils.MoveFile(raw[1], n.Region, "dynamodb")
				d := components.DynamoDB{
					Id:             raw[1],
					Name:           n.Properties["name"],
					Region:         n.Region,
					Read_capacity:  n.Properties["read_capacity"],
					Write_capacity: n.Properties["write_capacity"],
					Hash_key:       n.Properties["hashKey"],
				}
				components.Component.CreateComponent(d)
			} else {
				utils.MoveFile("", n.Region, "dynamodb")
				d := components.DynamoDB{
					Id:             "",
					Name:           n.Properties["name"],
					Region:         n.Region,
					Read_capacity:  n.Properties["read_capacity"],
					Write_capacity: n.Properties["write_capacity"],
					Hash_key:       n.Properties["hashKey"],
				}
				components.Component.CreateComponent(d)
			}
		} else if raw[0] == "RDS" {
			if raw[1] != "" {
				utils.MoveFile(raw[1], n.Region, "rds")
				createRDS(graph.Nodes, raw[1], n.Region)
			} else {
				utils.MoveFile("", n.Region, "rds")
				createRDS(graph.Nodes, "", n.Region)
			}
		}
		utils.EditProvider(n.Region)
	}
}
