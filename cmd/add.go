/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/otisnado/task/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to database selected`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Author: ", viper.GetString("author"))
		fmt.Println("Task: ", viper.GetString("task"))

		task := models.Task{
			Author:  viper.GetString("author"),
			Content: viper.GetString("task"),
			Done:    false,
			Date:    time.Now(),
		}

		result, err := models.CreateTask(task)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("ID: ", result.InsertedID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("task", "t", "", "task to be added")
	addCmd.MarkFlagRequired("task")
	viper.BindPFlag("task", addCmd.Flags().Lookup("task"))

	addCmd.Flags().StringP("author", "a", "", "task's author")
	addCmd.MarkFlagRequired("author")
	viper.BindPFlag("author", addCmd.Flags().Lookup("author"))
}
