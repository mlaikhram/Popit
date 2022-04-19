package mongoUtils

import (
	"context"
	"errors"

	"popit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func AddShow(client *mongo.Client, ctx context.Context, show models.Show, eps []models.Episode) (string, error) {
	showsCollection := client.Database("popit").Collection("shows")
	insertResult, err := showsCollection.InsertOne(ctx, show)
	if err != nil {
		return "", err
	}
	showId := insertResult.InsertedID.(primitive.ObjectID).Hex()
	err = AddEpisodes(client, ctx, showId, eps)
	if err != nil {
		return showId, err
	}
	return showId, nil
}

func GetShowByTitle(client *mongo.Client, ctx context.Context, name string) ([]models.Show, error) {
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

	var shows []models.Show
	if err = cur.All(ctx, &shows); err != nil {
		return nil, err
	}
	return shows, nil
}

func GetShowById(client *mongo.Client, ctx context.Context, showId string) ([]models.Show, error) {
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

	var shows []models.Show
	if err = cur.All(ctx, &shows); err != nil {
		return nil, err
	}
	return shows, nil
}

func EditShow(client *mongo.Client, ctx context.Context, showId string, show models.Show) error {
	showsCollection := client.Database("popit").Collection("shows")

	objId, err := primitive.ObjectIDFromHex(showId)
	if err != nil {
		return err
	}

	// show.ID = showId

	replaceResult, err := showsCollection.ReplaceOne(ctx, bson.M{"_id": objId}, show)
	if err != nil {
		return err
	} else if replaceResult.MatchedCount < 1 {
		return errors.New("Error: pageId " + showId + " return 0 results")
	}
	return nil
}
