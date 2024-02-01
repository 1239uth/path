package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func printPath(path []string, grep string) {
	slices.Sort(path)
	if grep != "" {
		var new_path []string
		for _, p := range path {
			if strings.Contains(p, grep) {
				new_path = append(new_path, p)
			}
		}
		path = new_path
	}

	fmt.Println(strings.Join(path, "\n"))
}

func main() {
	grepPtr := flag.String("grep", "", "String to search for in the path")
	flag.Parse()
	grep := *grepPtr

	pathEnv := os.Getenv("PATH")
	path := strings.Split(pathEnv, ":")
	printPath(path, grep)
}
