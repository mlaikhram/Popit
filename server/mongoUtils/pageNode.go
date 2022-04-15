package mongoUtils

import (
	"context"
	"errors"

	"popit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addPageNodes(client *mongo.Client, ctx context.Context, showId string, nodes []models.PageNode) error {
	fixedNodes := make([]interface{}, 0)

	for _, v := range nodes {
		v.ShowId = showId
		fixedNodes = append(fixedNodes, v)
	}

	nodesCollection := client.Database("popit").Collection("page_nodes")
	_, err := nodesCollection.InsertMany(ctx, fixedNodes)
	if err != nil {
		return err
	}
	return nil
}

func AddPageNodeForNewEpisode(client *mongo.Client, ctx context.Context, node models.PageNode) error {
	pageNodes, err := GetPageNodesBySectionId(client, ctx, node.SectionID)
	if err != nil {
		return err
	} else if len(pageNodes) <= 0 {
		return errors.New("error: Page node " + node.SectionID + " does not exist")
	}
	for i := range pageNodes {
		if pageNodes[i].EpisodeNum == node.EpisodeNum {
			return errors.New("error: Page node with this episode number already exists")
		}
	}
	err = addPageNodes(client, ctx, node.ShowId, []models.PageNode{node})
	if err != nil {
		return err
	}
	return nil
}

func InsertPageNode(client *mongo.Client, ctx context.Context, pageId string, node models.PageNode, index int) error {
	pagesCollection := client.Database("popit").Collection("pages")

	pageArr, err := getPageById(client, ctx, pageId)
	if err != nil {
		return err
	} else if len(pageArr) < 1 {
		return errors.New("Error: pageId " + pageId + " return 0 results")
	}

	page := pageArr[0]

	sectionIDs := append(page.SectionIDs[:index], append([]string{node.SectionID}, page.SectionIDs[index:]...)...)

	objId, err := primitive.ObjectIDFromHex(pageId)
	if err != nil {
		return err
	}

	updateResult, err := pagesCollection.UpdateOne(
		ctx,
		bson.M{"_id": objId},
		bson.M{"$set": bson.M{"sectionIds": sectionIDs}},
	)

	if err != nil {
		return err
	} else if updateResult.MatchedCount < 1 {
		return errors.New("Error: pageId " + pageId + " return 0 results")
	}

	err = addPageNodes(client, ctx, page.ShowId, []models.PageNode{node})
	if err != nil {
		return err
	}

	return nil
}

func GetPageNodesBySectionId(client *mongo.Client, ctx context.Context, sectionId string) ([]models.PageNode, error) {
	pageNodesCollection := client.Database("popit").Collection("page_nodes")

	opts := options.Find().SetProjection(bson.M{"_id": 0, "showId": 0, "sectionId": 0}).SetSort(bson.M{"episodeNum": 1})
	cur, err := pageNodesCollection.Find(ctx, bson.M{"sectionId": sectionId}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var pageNodes []models.PageNode
	if err = cur.All(ctx, &pageNodes); err != nil {
		return nil, err
	}
	return pageNodes, nil
}
