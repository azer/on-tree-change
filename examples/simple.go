package main

import (
	"fmt"
	. "github.com/azer/on-tree-change"
	"os"
	"strings"
)

func main() {
	err := OnTreeChange("/Users/azer/dev/go-choo-starter/ui", filter, func(name string) {
		fmt.Println(name + " changed")
	})

	if err != nil {
		panic(err)
	}
}

func filter(name string, info os.FileInfo) bool {
	return info.IsDir() && !strings.Contains(name, "node_modules")
}
