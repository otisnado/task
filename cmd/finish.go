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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// finishCmd represents the finish command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish a task",
	Long:  `Finish a task with the ID provided`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("finish called")

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
		update := bson.M{"$set": bson.M{"done": true}}
		result, err := tasks.UpdateMany(context.Background(), filter, update)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(result)

	},
}

func init() {
	rootCmd.AddCommand(finishCmd)
}

func IdsToObjectId(ids []string) (oids []primitive.ObjectID, err error) {

	for _, v := range ids {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return nil, err
		}
		oids = append(oids, id)
	}

	return oids, nil

}
