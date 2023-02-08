package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Redis struct {
	Id         string
	Name       string
	Cluster_id string
	Node_type  string
	Num_nodes  string
	Region     string
}

func (r Redis) CreateComponent() {
	src := "tf/" + r.Region + "/redis-" + r.Id + ".tf"
	txtFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened redis.tf")
	defer txtFile.Close()
	body, err := ioutil.ReadAll(txtFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	raw := strings.Split(string(body), "=")
	var final string
	identifier := strings.Split(raw[0][2:], "\"")
	final += identifier[0] + " \"" + identifier[1] + "\" \"redis_" + r.Id + "\"" +
		identifier[4] + "= \"" + r.Cluster_id + "\""
	final += raw[1] + "= \"redis\""
	final += raw[2] + "= \"" + r.Node_type + "\""
	final += raw[3] + "= " + r.Num_nodes
	final += raw[4] + "= \"default.redis3.2\""
	final += raw[5] + "= \"3.2.10\""
	final += raw[6] + "= 6379" + raw[7] + "= "
	final += raw[8] + "= \"" + r.Name + "\"" + raw[9][:len(raw[9])-2]

	ioutil.WriteFile(src, []byte(final), 0644)
}
