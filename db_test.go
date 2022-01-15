package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

type User2 struct {
	Name     string `bson:"name"`
	Age      int    `bson:"age"`
	Birthday time.Time
}

func init() {
	db, err := gorm.Open(mysql.Open(MYSQL_DSN))
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	if err != nil {
		return
	}
}

func TestMySQLConnection(t *testing.T) {
	var user = User{Name: "Yxxx", Age: 16, Birthday: time.Now()}
	db, err := gorm.Open(mysql.Open(MYSQL_DSN))
	if err != nil {
		t.Log(err)
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	if err != nil {
		return
	}
	db.Create(&user)
	var user2 User
	db.Where("name = ?", "Yxxx").Find(&user2)
	fmt.Println(user2.Name)
}

func TestMongoDBConnection(t *testing.T) {
	mongoOptions := options.Client().ApplyURI(MONGODB_URI)
	
	client, err := mongo.Connect(context.TODO(), mongoOptions)
	if err != nil {
		t.Log(err)
	}
	err = client.Ping(context.TODO(), nil)
	
	collection := client.Database(MONGO_DATABASE).Collection(MONGO_COLLECTION)
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		
		}
	}(client, context.TODO())
	var user User2
	err = collection.FindOne(context.TODO(), bson.M{"name": "hccc"}).Decode(&user)
	if err != nil {
		t.Log(err)
	}
	t.Log(user.Name, user.Age)
}
