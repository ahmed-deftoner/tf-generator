package utils

import "github.com/ahmed-deftoner/tf-generator/components"

func CheckService(service string, id string, n Node) {
	switch service {
	case "DynamoDB":
		MoveFile(id, n.Region, "dynamodb", "db")
		d := components.DynamoDB{
			Id:             id,
			Name:           n.Properties["name"],
			Region:         n.Region,
			Read_capacity:  n.Properties["read_capacity"],
			Write_capacity: n.Properties["write_capacity"],
			Hash_key:       n.Properties["hashKey"],
		}
		components.Component.CreateComponent(d)
	case "RDS":
		MoveFile(id, n.Region, "rds", "db")
		r := components.RDS{
			Id:               id,
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
	case "Redis":
		MoveFile(id, n.Region, "redis", "db")
		r := components.Redis{
			Id:         id,
			Name:       n.Properties["name"],
			Cluster_id: n.Properties["cluster_id"],
			Node_type:  n.Properties["node_type"],
			Num_nodes:  n.Properties["num_cache_nodes"],
			Region:     n.Region,
		}
		components.Component.CreateComponent(r)
	case "Memcached":
		MoveFile(id, n.Region, "memcached", "db")
		m := components.Memcached{
			Id:         id,
			Name:       n.Properties["name"],
			Cluster_id: n.Properties["cluster_id"],
			Node_type:  n.Properties["node_type"],
			Num_nodes:  n.Properties["num_cache_nodes"],
			Region:     n.Region,
		}
		components.Component.CreateComponent(m)
	case "Neptune":
		MoveFile(id, n.Region, "neptune", "db")
		m := components.Neptune{
			Id:             id,
			Name:           n.Properties["name"],
			Cluster_id:     n.Properties["cluster_id"],
			Instances:      n.Properties["instances"],
			Instance_class: n.Properties["instance_class"],
			Region:         n.Region,
		}
		components.Component.CreateComponent(m)
	case "Lambda":
		MoveFile(id, n.Region, "lambda", "compute")
		l := components.Lambda{
			Id:      id,
			Region:  n.Region,
			Name:    n.Properties["name"],
			Runtime: n.Properties["runtime"],
			Handler: n.Properties["handler"],
			Git:     n.Properties["git"],
		}
		components.Component.CreateComponent(l)
	}
}
