/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/otisnado/task/db"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task stored in database`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		tasks, err := db.GetCollection("tasks")
		if err != nil {
			log.Fatalln(err)
		}

		defer tasks.Database().Client().Disconnect(context.Background())

		ids, err := IdsToObjectId(args)
		if err != nil {
			log.Fatalln(err)
		}

		filter := bson.M{"_id": bson.M{"$in": ids}}
		result, err := tasks.DeleteMany(context.Background(), filter)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(result)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
