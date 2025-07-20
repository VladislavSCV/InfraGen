package utils

import (
	"os"
	"path/filepath"
	"text/template"
)

func GenerateManifestsForGo(answers GoAnswers, projectPath string) error {
	infraPath := filepath.Join(projectPath, "infra")
	err := os.MkdirAll(infraPath, os.ModePerm)
	if err != nil {
		return err
	}

	templatesDir := "templates/go"
	files := []string{
		"Dockerfile",
		"compose.yaml",
		".dockerignore",

		"deployment.yaml",
		"service.yaml",
		"configmap.yaml",
		"ingress.yaml",

		".gitlab-ci.yml",
	}

	for _, name := range files {
		tmplPath := filepath.Join(templatesDir, name+".tmpl")
		outPath := filepath.Join(infraPath, name)

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			return err
		}

		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		err = tmpl.Execute(outFile, answers)
		if err != nil {
			return err
		}
	}

	return nil
}
