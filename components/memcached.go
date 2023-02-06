package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Memcached struct {
	Id         string
	Name       string
	Cluster_id string
	Node_type  string
	Num_nodes  string
	Region     string
}

func (m Memcached) CreateComponent() {
	src := "tf/" + m.Region + "/memcached-" + m.Id + ".tf"
	txtFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened memcached.tf")
	defer txtFile.Close()
	body, err := ioutil.ReadAll(txtFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	raw := strings.Split(string(body), "=")
	var final string
	final += raw[0][2:] + "= \"" + m.Cluster_id + "\""
	final += raw[1] + "= \"memcached\""
	final += raw[2] + "= \"" + m.Node_type + "\""
	final += raw[3] + "= " + m.Num_nodes
	final += raw[4] + "= \"default.memcached.4\""
	final += raw[5] + "= 11211" + raw[6] + "= "
	final += raw[7] + "= \"" + m.Name + "\"" + raw[8][:len(raw[8])-2]

	ioutil.WriteFile(src, []byte(final), 0644)
}
