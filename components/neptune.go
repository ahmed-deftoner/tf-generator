package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Neptune struct {
	Id             string
	Name           string
	Cluster_id     string
	Instances      string
	Instance_class string
	Region         string
}

func (n Neptune) CreateComponent() {
	src := "tf/" + n.Region + "/neptune-" + n.Id + ".tf"
	txtFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened neptune.tf")
	defer txtFile.Close()
	body, err := ioutil.ReadAll(txtFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	raw := strings.Split(string(body), "=")
	var final string
	identifier := strings.Split(raw[0][2:], "\"")
	final += identifier[0] + " \"" + identifier[1] + "\" \"neptune_cluster_" + n.Id + "\"" +
		identifier[4] + "= \"" + n.Cluster_id + "\""
	final += raw[0][2:] + "= \"" + n.Cluster_id + "\""
	final += raw[1] + "= \"neptune\""
	final += raw[2] + "= true"
	final += raw[3] + "= true" + raw[4] + "= "
	final += raw[5] + "= \"" + n.Name + "\""
	identifier = strings.Split(raw[6], "\"")
	final += identifier[0] + " \"" + identifier[1] + "\" \"neptune_instance_" + n.Id + "\"" +
		identifier[4] + "= " + n.Instances
	final += raw[7] + "= aws_neptune_cluster.default.id"
	final += raw[8] + "= \"neptune\""
	final += raw[9] + "= \"" + n.Instance_class + "\""
	final += raw[10] + "= true" + raw[11] + "= "
	final += raw[12] + "= \"" + n.Name + "\"" + raw[13][:len(raw[13])-2]

	ioutil.WriteFile(src, []byte(final), 0644)
}
