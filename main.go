package main

import (

	//"fmt"

	"strings"

	"github.com/ahmed-deftoner/tf-generator/utils"
)

/*
// PowerShell struct
type PowerShell struct {
	powerShell string
}

// New create new session
func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

var (
	// Below command will enable the HyperV module
	clone = `git clone https://github.com/ahmed-deftoner/ec2-benchmark.git`
)*/

func main() {
	//	posh := New()

	// Scenario 1
	// stdOut, stdErr, err := posh.execute(elevateProcessCmds)
	// fmt.Printf("ElevateProcessCmds:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s", strings.TrimSpace(stdOut), stdErr, err)
	// ========= Above working and invoke a publisher permission dialog and Admin shell ================

	// Scenario 2
	// stdOut, stdErr, err := posh.execute(enableHyperVCmd)
	// fmt.Printf("\nEnableHyperV:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s", strings.TrimSpace(stdOut), stdErr, err)
	// ========= Behavior(expected one): StdErr: 'Enable-WindowsOptionalFeature : The requested operation requires elevation.

	// Scenario 3 : Both scenario 1 and 2 combined
	//	stdOut, stdErr, err := posh.execute(clone)
	//	fmt.Printf("\ngit clone:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s", strings.TrimSpace(stdOut), stdErr, err)
	// ========= Above suppose to open a permission dialog, on click of "yes" should
	// ========= run the hyperv enable command and once done ask for restart operation
	graph := utils.DecodeJSON("graph.json")
	utils.CreateDirs(graph.Nodes)
	m := []string{}
	for k, n := range graph.Nodes {
		if !utils.Contains(m, n.Region) {
			m = append(m, n.Region)
		}
		raw := strings.Split(k, "#")
		utils.CheckService(raw[0], raw[1], n)
	}
	/*	for _, v := range m {
		utils.EditProvider(v)
	}*/
}
