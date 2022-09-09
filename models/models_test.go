package models_test

import (
	"log"
	"testing"
	"time"

	"github.com/otisnado/task/models"
)

func TestCreateTask(t *testing.T) {
	task1 := models.Task{
		Author:  "Roberto",
		Content: "Limpiar",
		Done:    false,
		Date:    time.Now(),
	}

	result, err := models.CreateTask(task1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	log.Println(result.InsertedID)
}
