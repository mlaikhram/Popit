package mongoUtils

import (
	"context"
	"fmt"

	"popit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddEpisodes(client *mongo.Client, ctx context.Context, showId string, eps []models.Episode) error {
	fixedEps := make([]interface{}, 0)

	for _, v := range eps {
		v.ShowId = showId
		fixedEps = append(fixedEps, v)
	}
	fmt.Println(fixedEps)

	episodesCollection := client.Database("popit").Collection("episodes")
	_, err := episodesCollection.InsertMany(ctx, fixedEps)
	if err != nil {
		return err
	}
	return nil
}

func GetEpisodesByShowId(client *mongo.Client, ctx context.Context, showId string) ([]models.Episode, error) {
	epsiodesCollection := client.Database("popit").Collection("episodes")

	opts := options.Find().SetProjection(bson.M{"_id": 0, "showId": 0}).SetSort(bson.M{"number": 1})
	cur, err := epsiodesCollection.Find(ctx, bson.M{"showId": showId}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var episodes []models.Episode
	if err = cur.All(ctx, &episodes); err != nil {
		return nil, err
	}
	return episodes, nil
}
