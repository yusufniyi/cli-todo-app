/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-todo-app",
	Short: "A lightweight CLI to-do app for managing tasks, setting priorities, and staying organized directly from your terminal.",
	Long: `A simple and efficient CLI-based to-do app that helps you organize tasks, set priorities, and track progress directly from your terminal. 
	
	Perfect for developers and productivity enthusiasts who prefer lightweight tools without distractions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Easyclick CLI todo app")
		readEmailAndPassword()
		// create a user if it doesn't exists
		// login a user if it exists
		// Prevent unauthenticated user from running a command
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-todo-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readEmailAndPassword() {
	email := readInput("email")
	password := readInput("password")
	fmt.Println("aaa", email, password)
}

func readInput(inputName string) string {
	var input string
	var err error
	for {
		fmt.Printf(fmt.Sprintf("Please enter your %s\n", inputName))
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}

		if len(strings.TrimSpace(input)) <= 0 {
			continue
		}

		break
	}
	return input
}
