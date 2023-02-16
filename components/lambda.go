package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Lambda struct {
	Id      string
	Region  string
	Name    string
	Runtime string
	Handler string
	Git     string
}

func (l Lambda) CreateComponent() {
	src := "tf/" + l.Region + "/lambda-" + l.Id + ".tf"
	txtFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened lambda.tf")
	defer txtFile.Close()
	body, err := ioutil.ReadAll(txtFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	resources := strings.Split(string(body), "resource")
	var final string
	vars := strings.Split(resources[1], "=")
	identifier := strings.Split(vars[0], "\"")
	final += "resource \"" + identifier[1] + "\" \"" + l.Name +
		"_lambda_exec\"" + identifier[4] + "= \"" + l.Name +
		"_lambda\"" + vars[1] + "= " + vars[2]

	vars = strings.Split(resources[2], "=")
	identifier = strings.Split(vars[0], "\"")
	final += "resource \"" + identifier[1] + "\" \"" + l.Name +
		"_lambda_policy\"" + identifier[4] + "= aws_iam_role." +
		l.Name + "_lambda_exec.name" + vars[1] + "= " + vars[2]

	vars = strings.Split(resources[3], "=")
	identifier = strings.Split(vars[0], "\"")
	final += "resource \"" + identifier[1] + "\" \"" + l.Name +
		"\"" + identifier[4] + "= \"" + l.Name + "\"" +
		vars[1] + "= \"" + l.Runtime + "\"" + vars[2] + "= \"" + l.Handler +
		"\"" + vars[3] + "= data.archive_file.lambda_" + l.Name + ".output_path" +
		vars[4] + "= data.archive_file.lambda_" + l.Name + ".output_base64sha256" +
		vars[5] + "= aws_iam_role." + l.Name + "_lambda_exec.arn"

	identifier = strings.Split(vars[6], "\"")
	final += identifier[0] + "\"" + identifier[1] + "\" \"lambda_" + l.Name +
		"\"" + identifier[4] + "= " + vars[7] + "= \"../${path.module}/hello\"" +
		vars[8] + "= \"../${path.module}/hello.zip\"\n}"
	fmt.Println(final)
	/*final += identifier[0] + " \"" + identifier[1] + "\" \"" +
		l.Name + "_lambda_" + l.Id + "\"" +
		identifier[4] + "= \"" + l.Name + "-lambda" + "\""
	final += raw[0][2:] + "= " + raw[1]
	final += raw[1] + "= \"memcached\""
	final += raw[2] + "= \"" + l.Node_type + "\""
	final += raw[3] + "= " + l.Num_nodes
	final += raw[4] + "= \"default.memcached.4\""
	final += raw[5] + "= 11211" + raw[6] + "= "
	final += raw[7] + "= \"" + l.Name + "\"" + raw[8][:len(raw[8])-2]
	*/
	ioutil.WriteFile(src, []byte(final), 0644)
}
