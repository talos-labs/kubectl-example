package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
)

//go:embed resources/*
var tplFS embed.FS

func printUsage(failure bool) {
	if failure {
		fmt.Println("Not enough arguments")
	}
	fmt.Println("Usage: example <RESOURCE_NAME>")
	if failure {
		os.Exit(1)
	}
}

//go:embed resources/*.yaml
var resourceFS embed.FS

func downloadAndPrintResource(name string) {
	path := fmt.Sprintf("resources/%s.yaml", name)
	data, err := resourceFS.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage(true)
		return
	}

	param := args[1]
	if param == "-h" || param == "--help" {
		printUsage(false)
		return
	}

	resource, err := resourceName(param)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	downloadAndPrintResource(resource)
}
