package mongoUtils

import (
	"context"

	"popit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func AddPage(client *mongo.Client, ctx context.Context, page models.Page, nodes []models.PageNode) (string, error) {
	pagesCollection := client.Database("popit").Collection("pages")

	sectionIds := make([]string, 0)

	for _, v := range nodes {
		sectionIds = append(sectionIds, v.SectionID)
	}

	page.SectionIDs = sectionIds

	insertResult, err := pagesCollection.InsertOne(ctx, page)
	if err != nil {
		return "", err
	}
	pageId := insertResult.InsertedID.(primitive.ObjectID).Hex()
	err = addPageNodes(client, ctx, page.ShowId, nodes)
	if err != nil {
		return pageId, err
	}
	return pageId, nil
}

func GetPageById(client *mongo.Client, ctx context.Context, pageId string) ([]models.Page, error) {
	pagesCollection := client.Database("popit").Collection("pages")

	objId, err := primitive.ObjectIDFromHex(pageId)
	if err != nil {
		return nil, err
	}

	cur, err := pagesCollection.Find(ctx, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var shows []models.Page
	if err = cur.All(ctx, &shows); err != nil {
		return nil, err
	}
	return shows, nil
}
