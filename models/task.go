package models

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/otisnado/task/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Author  string             `bson:"author"`
	Content string             `bson:"content"`
	Done    bool               `bson:"done"`
	Date    time.Time          `bson:"date"`
}

func CreateTask(t Task) (resultOne *mongo.InsertOneResult, err error) {
	tasks, err := db.GetCollection("tasks")
	if err != nil {
		return nil, err
	}
	result, err := tasks.InsertOne(context.Background(), t)
	if err != nil {
		return nil, err
	}

	defer tasks.Database().Client().Disconnect(context.Background())
	return result, nil
}

func (t Task) String() string {
	done := color.New(color.FgGreen).SprintFunc()
	if !t.Done {
		done = color.New(color.FgRed).SprintFunc()
	}

	green := color.New(color.FgGreen).SprintFunc()

	return fmt.Sprintf("ID: %-20s \t Author: %-20s \t Task: %-20s \t Done: %-20s \t", t.ID, green(t.Author), t.Content, done(t.Done))
}
