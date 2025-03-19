package main

import (
	"bufio"
	"embed"
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

func downloadAndPrintResource(name string) {
	resources, err := tplFS.ReadDir("../../resources")
	if err != nil {
		panic(err)
	}
	for _, resource := range resources {
		if resource.Name() == name+".yaml" {
			resp, err := tplFS.Open("resources/" + resource.Name())
			if err != nil {
				panic(err)
			}
			defer resp.Close()
			scanner := bufio.NewScanner(resp)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}

	}

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
