package keysound

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PROJECT = "keysound"
const COLLECTION = "quiz"

var DefaultClient *mongo.Client
var DefaultMusicList MusicList
var DefaultChoices Choices

func init() {
	if os.Getenv("SKIP_DB") != "true" {
		InitDB(os.Getenv("MONGO_URI"))
	}

	var err error
	DefaultMusicList, err = ParseList("list_data.csv")
	if err != nil {
		log.Fatalln(err)
	}

	DefaultChoices = DefaultMusicList.Choices()

	rand.Seed(time.Now().UnixNano())
}

func InitDB(uri string) {
	fmt.Println("Init...")
	var err error
	DefaultClient, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = DefaultClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//
	//// check if db exist
	//databases, err := client.ListDatabaseNames(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Databases:",databases)
	//hasDB := false
	//for _, dbName:= range databases {
	//	if dbName== PROJECT {
	//		hasDB = true
	//	}
	//}
	//
	//// if db doesn't exist
	//if !hasDB {
	//	fmt.Println("Database not found, creating(with collection):",PROJECT,COLLECTION)
	//	err := client.Database(PROJECT).CreateCollection(ctx, COLLECTION)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//} else {
	//	// if db exists
	//	// then check if collection exists
	//	fmt.Println("Database found:",PROJECT)
	//	collections, err := client.Database(PROJECT).ListCollectionNames(ctx, bson.M{})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("Collections:",collections)
	//	hasCo:= false
	//	for _, cName:= range collections {
	//		if cName== COLLECTION {
	//			hasCo = true
	//		}
	//	}
	//
	//	// if collection doesn't exist
	//	if !hasCo {
	//		fmt.Println("Collection not found, creating:",COLLECTION)
	//		err := client.Database(PROJECT).CreateCollection(ctx, COLLECTION)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	} else {
	//		fmt.Println("Collection found:",COLLECTION)
	//	}
	//}

	fmt.Println("Init finished")
}
