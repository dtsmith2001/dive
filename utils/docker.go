package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// RunDockerCmd runs a given Docker command in the current tty
func RunDockerCmd(cmdStr string, args ...string) error {

	allArgs := cleanArgs(append([]string{cmdStr}, args...))

	cmd := exec.Command("docker", allArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// cleanArgs trims the whitespace from the given set of strings.
func cleanArgs(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, strings.Trim(str, " "))
		}
	}
	return r
}

func DiscoverDockerVersion() string {
	cmd := exec.Command("docker", "version", "--format", "{{.Server.APIVersion}}")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(cmdOutput.Bytes()))
}
