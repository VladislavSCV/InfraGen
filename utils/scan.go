package utils

import (
	"fmt"
	"path/filepath"
	"os"
)

func ScanFiles(projectPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			files = append(files, ext)
		}
		return nil
	})

	return files, err
}

func DetectLanguages(files []string) []string {
	langs := map[string]string{
		".go":     "Go - suitable for a Go server application",
		".py":     "Python - suitable for a Python server application",
		".js":     "Node - suitable for a Node server application",
		".ts":     "Node - suitable for a Node server application",
		".rs":     "Rust - suitable for a Rust server application",
		".csproj": "ASP.NET Core - suitable for an ASP.NET Core application",
		".php":    "PHP with Apache - suitable for a PHP web application",
		".java":   "Java - suitable for a Java application",
		"pom.xml": "Java - uses Maven",
	}

	found := make(map[string]string)
	for _, f := range files {
		if val, ok := langs[f]; ok {
			found[f] = val
		}
	}

	var results []string
	for _, v := range found {
		results = append(results, v)
	}

	return results
}

func PromptUser(choices []string) string {
	fmt.Println("🧠 Detected possible languages / stacks:")
	for i, choice := range choices {
		fmt.Printf("%d) %s\n", i+1, choice)
	}
	fmt.Printf("%d) Other - general purpose\n", len(choices)+1)
	fmt.Printf("%d) Quit\n", len(choices)+2)

	fmt.Print("Choose your stack: ")
	var selection int
	fmt.Scanln(&selection)

	if selection <= 0 || selection > len(choices)+2 {
		return "invalid"
	}

	if selection == len(choices)+1 {
		return "other"
	}
	if selection == len(choices)+2 {
		return "quit"
	}

	return choices[selection-1]
}
