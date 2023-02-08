package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type DynamoDB struct {
	Id             string
	Name           string
	Read_capacity  string
	Write_capacity string
	Hash_key       string
	Region         string
}

func (d DynamoDB) CreateComponent() {
	src := "tf/" + d.Region + "/dynamodb-" + d.Id + ".tf"
	txtFile, err := os.Open(src)
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
	identifier := strings.Split(raw[0][2:], "\"")
	final += identifier[0] + " \"" + identifier[1] + "\" \"dynamodb_" + d.Id + "\"" +
		identifier[4] + "= \"" + d.Name + "\""
	final += raw[1] + "= " + d.Read_capacity
	final += raw[2] + "= " + d.Write_capacity
	final += raw[3] + "= \"" + d.Hash_key + "\""
	final += raw[4] + "= \"" + d.Hash_key + "\""
	final += raw[5] + "= \"S\"" + raw[6][:len(raw[6])-2]

	ioutil.WriteFile(src, []byte(final), 0644)
}
