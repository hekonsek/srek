package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Unknown command.")
		os.Exit(-1)
	}

	if os.Args[1] == "provision-server" && os.Args[2] == "aws" && os.Args[3] == "master" {
		if len((os.Args)) == 4 {
			fmt.Println("Please specify the name of the master server.")
			os.Exit(-1)
		}

		cacheAnsibleCommand("provision-server-aws-master")

		name := os.Args[4]
		cmd := exec.Command("ansible-playbook","/tmp/srek-provision-server-aws-master.yml",
			fmt.Sprintf("-e instance_name=%s group=%s", name, name))
		cmd.Env = append(os.Environ(), "ANSIBLE_HOST_KEY_CHECKING=False")
		stdoutStderr, err := cmd.CombinedOutput()
		exitOnError(err, string(stdoutStderr))
		fmt.Println(string(stdoutStderr))
	} else if os.Args[1] == "provision" && os.Args[2] == "master" {
		cacheAnsibleCommand("provision-master")

		host := os.Args[3]
		cmd := exec.Command("ansible-playbook", fmt.Sprintf("-i%s,", host), "/tmp/srek-provision-master.yml")
		cmd.Env = append(os.Environ(), "ANSIBLE_HOST_KEY_CHECKING=False")
		stdoutStderr, err := cmd.CombinedOutput()
		exitOnError(err, string(stdoutStderr))
		fmt.Println(string(stdoutStderr))
	} else {
		fmt.Println("Unknown command.")
	}
}

// Ansible

func cacheAnsibleCommand(command string)  {
	yml, err := Asset(fmt.Sprintf("ansible/%s.yml", command))
	exitOnError(err)
	err = ioutil.WriteFile(fmt.Sprintf("/tmp/srek-%s.yml", command), yml, 0644)
	exitOnError(err)
}

// Utils

func exitOnError(err error, context... interface{}) {
	if err != nil {
		fmt.Println(err)
		if len(context) > 0 {
			fmt.Println(context)
		}
		os.Exit(-1)
	}
}