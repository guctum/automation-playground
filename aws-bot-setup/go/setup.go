package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Now creating the AWS infrastructure for the Discord bot")
	// run cdk synth and cdk deploy in the infra-setup directory
	cmd := exec.Command("cdk deploy")
	cmd.Dir = "../infra-setup"
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else  {
		fmt.Printf("%s", out)
	}
	fmt.Println("AWS Infrastructure successfully created")
	// create a git repo as well, see if we can share it on GitHub too
	fmt.Println("Now creating git repository")
	repoSetup()
}
