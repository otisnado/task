package db_test

import (
	"context"
	"log"
	"testing"

	"github.com/otisnado/task/db"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetConn(t *testing.T) {
	client, err := db.GetConnection()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	listdb, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	log.Println(listdb)
}

func TestGetCollection(t *testing.T) {
	tasks, err := db.GetCollection("tasks")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	result := tasks.FindOne(context.Background(), bson.D{})
	var doc bson.D
	err = result.Decode((&doc))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = tasks.Database().Client().Disconnect(context.Background())
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	log.Println(doc)
}

func TestConfig(t *testing.T) {
	db.LoadConfig()
}
