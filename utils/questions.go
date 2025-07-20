package utils

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

type GoAnswers struct {
	AppName        string
	Namespace      string
	Image          string
	GoVersion      string
	DockerRegistry string
	MainDir        string
	EnableDeploy   bool
	ListenPort     int
}

func RunGoQuestions(projectPath string) {
	var answers GoAnswers

	// Строковые вопросы
	survey.AskOne(&survey.Input{
		Message: "✅ Application name:",
		Default: "my-go-app",
	}, &answers.AppName)

	survey.AskOne(&survey.Input{
		Message: "✅ Kubernetes namespace:",
		Default: "default",
	}, &answers.Namespace)

	survey.AskOne(&survey.Input{
		Message: "✅ Docker image (e.g. myrepo/myimage:tag):",
		Default: "myrepo/myimage:latest",
	}, &answers.Image)

	survey.AskOne(&survey.Input{
		Message: "✅ Go version:",
		Default: "1.22.1",
	}, &answers.GoVersion)

	survey.AskOne(&survey.Input{
		Message: "✅ Docker registry URL:",
		Default: "registry.gitlab.com/myproject",
	}, &answers.DockerRegistry)

	survey.AskOne(&survey.Input{
		Message: "📂 Relative directory of main package:",
		Default: ".",
	}, &answers.MainDir)

	// Логический вопрос (через input + парсинг)
	var deployStr string
	survey.AskOne(&survey.Input{
		Message: "🚀 Enable deployment manifest generation? (true/false):",
		Default: "true",
	}, &deployStr)
	answers.EnableDeploy = deployStr == "true" || deployStr == "1"

	// Целочисленный вопрос (порт)
	var portStr string
	survey.AskOne(&survey.Input{
		Message: "🔌 Server listen port:",
		Default: "8080",
	}, &portStr)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080 // дефолт, если ввели что-то не то
	}
	answers.ListenPort = port

	fmt.Println("\n🚀 Генерация файлов началась...")

	err = GenerateManifestsForGo(answers, projectPath)
	if err != nil {
		fmt.Println("❌ Ошибка генерации:", err)
	} else {
		fmt.Println("✅ Все манифесты успешно сгенерированы в папке infra.")
	}
}
