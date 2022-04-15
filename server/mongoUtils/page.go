package mongoUtils

import (
	"context"
	"errors"
	"log"

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

func getPageById(client *mongo.Client, ctx context.Context, pageId string) ([]models.Page, error) {
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

	var pages []models.Page
	if err = cur.All(ctx, &pages); err != nil {
		return nil, err
	}
	return pages, nil
}

func GetPageByEpisode(client *mongo.Client, ctx context.Context, pageId string, episodeNum int) ([]models.PageNode, error) {
	pages, err := getPageById(client, ctx, pageId)
	if err != nil {
		return nil, err
	}
	page := pages[0]
	log.Println("page: " + page.ID)
	pageNodesCollection := client.Database("popit").Collection("page_nodes")

	cur, err := pageNodesCollection.Aggregate(ctx, bson.A{
		bson.M{"$match": bson.M{
			"episodeNum": bson.M{"$lte": episodeNum},
			"sectionId":  bson.M{"$in": page.SectionIDs},
		}},
		bson.M{"$sort": bson.M{"episodeNum": -1}},
		bson.M{"$group": bson.M{"_id": "$sectionId", "doc_with_max_episodeNum": bson.M{"$first": "$$ROOT"}}},
		bson.M{"$replaceWith": "$doc_with_max_episodeNum"},
		bson.M{"$addFields": bson.M{"__order": bson.M{"$indexOfArray": bson.A{page.SectionIDs, "$sectionId"}}}},
		bson.M{"$sort": bson.M{"__order": 1}},
	})
	if err != nil {
		return nil, err
	}

	var pageNodes []models.PageNode
	if err = cur.All(ctx, &pageNodes); err != nil {
		return nil, err
	}
	return pageNodes, nil
}

func RevalidatePageSections(client *mongo.Client, ctx context.Context, pageId string) error {
	pagesCollection := client.Database("popit").Collection("pages")

	objId, err := primitive.ObjectIDFromHex(pageId)
	if err != nil {
		return err
	}

	cur, err := pagesCollection.Find(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	var pages []models.Page
	if err = cur.All(ctx, &pages); err != nil {
		return err
	}

	page := pages[0]
	newSectionIDs := make([]string, 0)
	for _, v := range page.SectionIDs {
		pageNodes, err := GetPageNodesBySectionId(client, ctx, v)
		if err != nil || pageNodes == nil || len(pageNodes) <= 0 {
			log.Println("page node " + v + " does not exist; removing")
		} else {
			newSectionIDs = append(newSectionIDs, v)
		}
	}

	updateResult, err := pagesCollection.UpdateOne(
		ctx,
		bson.M{"_id": objId},
		bson.M{"$set": bson.M{"sectionIds": newSectionIDs}},
	)

	if err != nil {
		return err
	} else if updateResult.MatchedCount < 1 {
		return errors.New("Error: pageId " + pageId + " return 0 results")
	}

	return nil
}
