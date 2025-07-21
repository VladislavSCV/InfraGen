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

	files := []string{
		"Dockerfile.tmpl",
		"compose.yaml.tmpl",
		".dockerignore.tmpl",
		"deployment.yaml.tmpl",
		"service.yaml.tmpl",
		"configmap.yaml.tmpl",
		"ingress.yaml.tmpl",
		".gitlab-ci.yml.tmpl",
	}

	for _, name := range files {
		data, err := Asset(name)
		if err != nil {
			return err
		}

		tmpl, err := template.New(name).Parse(string(data))
		if err != nil {
			return err
		}

		outPath := filepath.Join(infraPath, name[:len(name)-5]) // без .tmpl
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
