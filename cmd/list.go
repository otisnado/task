/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/otisnado/task/db"
	"github.com/otisnado/task/models"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks saved",
	Long:  `list all tasks saved in database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all task stored in database")

		tasks, err := db.GetCollection("tasks")
		if err != nil {
			log.Fatalln(err)
		}

		curs, err := tasks.Find(context.Background(), bson.D{})
		if err != nil {
			log.Fatalln(err)
		}

		defer curs.Close(context.Background())

		for curs.Next(context.Background()) {
			var task models.Task

			err = curs.Decode(&task)
			if err != nil {
				log.Fatalln(err)
			}

			// fmt.Println(task)
			color.New().Println(task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
