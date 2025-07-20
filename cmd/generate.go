/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"infraGen/utils"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Укажи путь до проекта")
			return
		}

		path := args[0]

		files, err := utils.ScanFiles(path)
		if err != nil {
			fmt.Println("Ошибка при сканировании:", err)
			return
		}

		detected := utils.DetectLanguages(files)
		choice := utils.PromptUser(detected)

		switch choice {
			case "Go - suitable for a Go server application":
				utils.RunGoQuestions(path)

			case "Python - suitable for a Python server application":
				// TODO: runPythonQuestions(path)

			case "other":
				fmt.Println("🛠 Пока реализована только генерация для Go. Скоро будет больше.")

			case "quit":
				fmt.Println("👋 Выход.")
				os.Exit(0)

			default:
				fmt.Println("❗ Неизвестный выбор.")
			}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
