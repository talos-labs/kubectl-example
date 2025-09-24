package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed resources/*
var tplFS embed.FS

func printUsage(failure bool) {
	if failure {
		fmt.Println("Not enough arguments")
	}
	fmt.Println("Usage: example <RESOURCE_NAME>\nResources:")
	files, err := tplFS.ReadDir("resources")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("- %s \n", strings.Split(file.Name(), ".")[0])
	}
}

func downloadAndPrintResource(name string) {
	path := fmt.Sprintf("resources/%s.yaml", name)
	data, err := tplFS.ReadFile(path)
	if err != nil {
		log.Fatal(err)
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
