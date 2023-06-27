package DB

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkMongoWorker struct {
	Client        *mongo.Client
	UrlCollection *mongo.Collection
}

func NewLinkWorker(c *mongo.Client) *LinkMongoWorker {
	lmw := &LinkMongoWorker{
		Client: c,
	}
	lmw.UrlCollection = lmw.Client.Database("URL_Shortener").Collection("URLs")
	return lmw
}

func (lmw *LinkMongoWorker) AddRecordToURLCol(link, shortenedURL string) {
	linkBSON := bson.D{{Key: "link", Value: link}, {Key: "ShortendLink", Value: shortenedURL}}
	result, err := lmw.UrlCollection.InsertOne(context.TODO(), linkBSON)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	log.Println(result.InsertedID)
}

func (lmw *LinkMongoWorker) Findlink(link string) bool {
	var result bson.M
	err := lmw.UrlCollection.FindOne(context.TODO(), bson.D{{"link", link}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the title %s\n", link)
		return false
	}
	log.Println("55555555555555555")
	if err != nil {
		panic(err)
	}
	log.Println("6666666666666666666")

	jsonData, err := json.MarshalIndent(result, "", "    ")
	log.Println("444444444444444444")

	if err != nil {
		panic(err)
	}
	log.Printf("%s\n!!!!!!!!!!!!!!!!!!!!!!", jsonData)
	return true
}
