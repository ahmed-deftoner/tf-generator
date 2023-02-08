package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type RDS struct {
	Id               string
	Region           string
	AllocatedStorage string
	DBName           string
	Engine           string
	EngineVersion    string
	Instance         string
	Username         string
	Password         string
}

func (r RDS) CreateComponent() {
	src := "tf/" + r.Region + "/rds-" + r.Id + ".tf"
	txtFile, err := os.Open(src)
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
	identifier := strings.Split(raw[0][2:], "\"")
	final += identifier[0] + " \"" + identifier[1] + "\" \"rds_" + r.Id + "\"" +
		identifier[4] + "= " + r.AllocatedStorage
	final += raw[0][2:] + "= " + r.AllocatedStorage
	final += raw[1] + "= \"" + r.DBName + "\""
	final += raw[2] + "= \"" + r.Engine + "\""
	final += raw[3] + "= \"" + r.EngineVersion + "\""
	final += raw[4] + "= \"" + r.Instance + "\""
	final += raw[5] + "= \"" + r.Username + "\""
	final += raw[6] + "= \"" + r.Password + "\""
	final += raw[7] + "= true" + raw[8][:len(raw[8])-2]

	ioutil.WriteFile(src, []byte(final), 0644)
}
