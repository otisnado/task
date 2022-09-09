package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cfg Config
)

func init() {
	LoadConfig()
}

func GetConnection() (client *mongo.Client, err error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", cfg.DB_user, cfg.DB_password, cfg.DB_host, cfg.DB_name)
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetCollection(name string) (col *mongo.Collection, err error) {
	client, err := GetConnection()
	if err != nil {
		return nil, err
	}

	return client.Database("task-admin").Collection(name), nil
}

type Config struct {
	DB_name     string `mapstructure: "DB_NAME"`
	DB_host     string `mapstructure: "DB_HOST"`
	DB_user     string `mapstructure: "DB_USER"`
	DB_password string `mapstructure: "DB_PASSWORD"`
}

func LoadConfig() {
	vi := viper.New()

	path, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
		return
	}

	vi.AddConfigPath(path + "/conftask")
	vi.SetConfigType("env")
	vi.SetConfigName("app")

	vi.AutomaticEnv()

	err = vi.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return
	} else {
		log.Println("Read config: ", vi.ConfigFileUsed())
	}

	vi.Unmarshal(&cfg)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
