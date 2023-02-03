package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func CreateDirs(nodes map[string]Node) {
	var path string

	for _, n := range nodes {
		path = "tf/" + n.Region
		_, err := os.Stat(path)
		if err == nil {
			continue
		}
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func moveFile(id string, region string, service string) {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := ioutil.ReadFile("templates/" + service + ".tf")
	if err != nil {
		log.Fatal(err)
	}
	dst := "tf/" + region + "/" + service + "-" + id + ".tf"
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func createDynamoDB(node map[string]Node, id string, region string) {
	src := "tf/" + region + "/dynamodb-" + id + ".tf"
	txtFile, err := os.Open(src)
	key := "DynamoDB#" + id
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened dynamodb.tf")
	defer txtFile.Close()
	body, err := ioutil.ReadAll(txtFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	raw := strings.Split(string(body), "=")
	var final string
	final += raw[0][2:] + "= \"" + node[key].Properties["name"] + "\""
	final += raw[1] + "= " + node[key].Properties["read_capacity"]
	final += raw[2] + "= " + node[key].Properties["write_capacity"]
	final += raw[3] + "= \"" + node[key].Properties["hashKey"] + "\""
	final += raw[4] + "= \"" + node[key].Properties["hashKey"] + "\""
	final += raw[5] + "= \"S\"" + raw[6][:len(raw[6])-2]

	ioutil.WriteFile(src, []byte(final), 0644)
}

func createRDS(node map[string]Node, id string, region string) {
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

func editProvider(region string) {
	prov := "\n\nmodule \"" + region + "\"" + "{\n  source = \"./" +
		region + "\"\n  providers = {\n    aws = aws." + region + "\n   }\n}"
	file, err := os.OpenFile("tf/provider.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal("Could not open provider.tf")
		return
	}

	defer file.Close()

	_, err2 := file.WriteString(prov)

	if err2 != nil {
		log.Fatal("Could not write text to provider.tf")
	} else {
		log.Println("Operation successful! Text has been appended to provider.tf")
	}
}

func main() {
	graph := DecodeJSON("graph.json")
	CreateDirs(graph.Nodes)
	for k, n := range graph.Nodes {
		raw := strings.Split(k, "#")
		if raw[0] == "DynamoDB" {
			if raw[1] != "" {
				moveFile(raw[1], n.Region, "dynamodb")
				createDynamoDB(graph.Nodes, raw[1], n.Region)
			} else {
				moveFile("", n.Region, "dynamodb")
				createDynamoDB(graph.Nodes, "", n.Region)
			}
			editProvider(n.Region)
		} else if raw[0] == "RDS" {
			if raw[1] != "" {
				moveFile(raw[1], n.Region, "rds")
				createRDS(graph.Nodes, raw[1], n.Region)
			} else {
				moveFile("", n.Region, "rds")
				createRDS(graph.Nodes, "", n.Region)
			}
		}
	}
}
