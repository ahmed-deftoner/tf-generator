package main

import (
	"strings"

	"github.com/ahmed-deftoner/tf-generator/components"
	"github.com/ahmed-deftoner/tf-generator/utils"
)

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
				r := components.RDS{
					Id:               raw[1],
					AllocatedStorage: n.Properties["allocated_storage"],
					Region:           n.Region,
					DBName:           n.Properties["db_name"],
					Engine:           n.Properties["engine"],
					EngineVersion:    n.Properties["engine_version"],
					Instance:         n.Properties["instance_class"],
					Username:         n.Properties["username"],
					Password:         n.Properties["password"],
				}
				components.Component.CreateComponent(r)
			} else {
				utils.MoveFile("", n.Region, "rds")
				r := components.RDS{
					Id:               raw[1],
					AllocatedStorage: n.Properties["allocated_storage"],
					Region:           n.Region,
					DBName:           n.Properties["db_name"],
					Engine:           n.Properties["engine"],
					EngineVersion:    n.Properties["engine_version"],
					Instance:         n.Properties["instance_class"],
					Username:         n.Properties["username"],
					Password:         n.Properties["password"],
				}
				components.Component.CreateComponent(r)
			}
		} else if raw[0] == "Redis" {
			utils.MoveFile(raw[1], n.Region, "redis")
			r := components.Redis{
				Id:         raw[1],
				Name:       n.Properties["name"],
				Cluster_id: n.Properties["cluster_id"],
				Node_type:  n.Properties["node_type"],
				Num_nodes:  n.Properties["num_cache_nodes"],
				Region:     n.Region,
			}
			components.Component.CreateComponent(r)
		}
		utils.EditProvider(n.Region)
	}
}
