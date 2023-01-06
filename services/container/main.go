package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var bashPath = "C:\\Program Files\\Git\\git-bash.exe"

func start(name string) error {
	command := fmt.Sprintf("docker start %s", name)
	cmd := exec.Command(bashPath, command)
	return cmd.Run()
}

func stop(name string) error {
	command := fmt.Sprintf("docker stop %s", name)
	cmd := exec.Command(bashPath, command)
	return cmd.Run()
}

func create(name, token string) error {
	command := fmt.Sprintf(`
		docker run --name %s -e 
	`)
	cmd := exec.Command(bashPath, command)
	return cmd.Run()
}

func main() {
	var command string
	var name string
	var token string

	flag.StringVar(&name, "name", "", "name of the container")
	flag.StringVar(&token, "token", "", "token of the bot")
	flag.StringVar(&command, "command", "", "create/start/stop")
	flag.Parse()

	var err error
	switch command {
	case "stop":
		err = stop(name)
		break
	case "start":
		err = start(name)
		break
	case "create":
		err = create(name, token)
		break
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
