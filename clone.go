package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// the following can be overriden via ldflags during compile time
// @see Makefile
var DefaultOwner = "hectron"
var GitRemoteUrl = "github.com"

type Repository struct {
	Owner, Name string
}

func printUsage() {
	fmt.Print(`Usage:

	clone [REPOSITORY]
	clone [ORGANIZATION]/[REPOSITORY]
`)
}

func cloneRepo(url string) {
	cmd := exec.Command("git", "clone", "--progress", url)
	stderr, _ := cmd.StderrPipe()

	cmd.Start()

	scanner := bufio.NewScanner(stderr)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func buildSshUrl(repositoryOwner, repositoryName string) string {
	return fmt.Sprintf("git@%s:%s/%s", GitRemoteUrl, repositoryOwner, repositoryName)
}

func parseInput(input string) *Repository {
	repo := &Repository{}

	if strings.Contains(input, " ") {
		return repo
	}

	inputAsParts := strings.Split(input, "/")
	numberOfPortions := len(inputAsParts)

	if numberOfPortions == 1 {
		repo.Owner = DefaultOwner
		repo.Name = inputAsParts[0]

		return repo
	}

	if numberOfPortions == 2 {
		repo.Owner = inputAsParts[0]
		repo.Name = inputAsParts[1]

		return repo
	}

	return repo
}

func main() {
	argsWithoutProgramName := os.Args[1:]

	url := ""

	if len(argsWithoutProgramName) == 1 {
		repository := parseInput(argsWithoutProgramName[0])

		url = buildSshUrl(repository.Owner, repository.Name)
	}

	if url == "" {
		printUsage()
		os.Exit(1)
	}

	cloneRepo(url)
}
