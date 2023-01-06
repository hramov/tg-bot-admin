package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func start(name string) (string, error) {
	cmd := exec.Command("docker", "start", name)
	out, err := cmd.Output()
	return string(out), err
}

func startAll() (string, error) {
	namesStr, err := getAll()
	names := strings.Split(namesStr, "\n")
	if err != nil {
		return "", err
	}

	for _, n := range names {
		cmd := exec.Command("docker", "start", n)
		_, err := cmd.Output()
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func stop(name string) (string, error) {
	cmd := exec.Command("docker", "stop", name)
	out, err := cmd.Output()
	return string(out), err
}

func stopAll() (string, error) {
	namesStr, err := getAll()
	names := strings.Split(namesStr, "\n")
	if err != nil {
		return "", err
	}

	for _, n := range names {
		cmd := exec.Command("docker", "stop", n)
		_, err := cmd.Output()
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func getStarted() (string, error) {
	cmd := exec.Command("docker", "ps", "--format", "{{.Names}}")
	out, err := cmd.Output()
	return string(out), err
}

func getAll() (string, error) {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}")
	out, err := cmd.Output()
	return string(out), err
}

func create(name, token string) error {
	cmd := exec.Command("docker", "")
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
	var out string
	switch command {
	case "stop":
		out, err = stop(name)
		break
	case "stopall":
		out, err = stopAll()
		break
	case "start":
		out, err = start(name)
		break
	case "startall":
		out, err = startAll()
		break
	case "create":
		err = create(name, token)
		break
	case "started":
		out, err = getStarted()
		break
	case "all":
		out, err = getAll()
	}

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(out)
}
