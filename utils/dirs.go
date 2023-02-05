package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func Contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
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

func MoveFile(id string, region string, service string) {
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

func EditProvider(region string) {
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
