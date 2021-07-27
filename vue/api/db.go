package handler

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dbPLR struct {
	Quiz PieceListResponse `json:"quiz",bson:"quiz"`
}

func MongoNewQuiz(plr PieceListResponse) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := client.Database(PROJECT).Collection(COLLECTION).
		InsertOne(ctx,
			dbPLR{
				Quiz: plr,
			})
	if err != nil {
		return "", err
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("unknown InsertedID type")
	}
	return id.Hex(), nil
}

func MongoGetQuiz(id string) (PieceListResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result := client.Database(PROJECT).Collection(COLLECTION).
		FindOne(ctx, bson.M{"_id": objectID})
	dbplr := dbPLR{}
	err = result.Decode(&dbplr)
	if err != nil {
		return nil, err
	}
	return dbplr.Quiz, nil
}
