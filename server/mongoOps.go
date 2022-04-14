package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func getClient() (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"))
}

func addShow(client *mongo.Client, ctx context.Context, show Show, eps []Episode) (string, error) {
	showsCollection := client.Database("popit").Collection("shows")
	insertResult, err := showsCollection.InsertOne(ctx, show)
	if err != nil {
		return "", err
	}
	showId := insertResult.InsertedID.(primitive.ObjectID).Hex()
	err = addEpisodes(client, ctx, showId, eps)
	if err != nil {
		return showId, err
	}
	return showId, nil
}

func addEpisodes(client *mongo.Client, ctx context.Context, showId string, eps []Episode) error {
	fixedEps := make([]interface{}, 0)

	for _, v := range eps {
		v.ShowId = showId
		fixedEps = append(fixedEps, v)
	}
	fmt.Println(fixedEps)

	showsCollection := client.Database("popit").Collection("episodes")
	_, err := showsCollection.InsertMany(ctx, fixedEps)
	if err != nil {
		return err
	}
	return nil
}

func getShows(client *mongo.Client, ctx context.Context, name string) ([]Show, error) {
	showsCollection := client.Database("popit").Collection("shows")

	nameRegex := bson.M{
		"$regex": primitive.Regex{
			Pattern: name,
			Options: "i",
		},
	}

	cur, err := showsCollection.Find(ctx, bson.M{"$or": bson.A{bson.M{"name": nameRegex}, bson.M{"aliases": nameRegex}}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var shows []Show
	if err = cur.All(ctx, &shows); err != nil {
		return nil, err
	}
	return shows, nil
}

func getShowById(client *mongo.Client, ctx context.Context, showId string) ([]Show, error) {
	showsCollection := client.Database("popit").Collection("shows")

	objId, err := primitive.ObjectIDFromHex(showId)
	if err != nil {
	  return nil, err
	}

	cur, err := showsCollection.Find(ctx, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var shows []Show
	if err = cur.All(ctx, &shows); err != nil {
		return nil, err
	}
	return shows, nil
}

func getEpisodes(client *mongo.Client, ctx context.Context, showId string) ([]Episode, error) {
	epsiodesCollection := client.Database("popit").Collection("episodes")

	opts := options.Find().SetProjection(bson.M{"_id": 0, "showId": 0}).SetSort(bson.M{"number": 1})
	cur, err := epsiodesCollection.Find(ctx, bson.M{"showId": showId}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var episodes []Episode
	if err = cur.All(ctx, &episodes); err != nil {
		return nil, err
	}
	return episodes, nil
}
