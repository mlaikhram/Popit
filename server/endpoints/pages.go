package endpoints

import (
	"context"
	"net/http"

	"popit/mongoUtils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPageNodesBySectionId(c *gin.Context, client *mongo.Client, ctx context.Context, sectionId string) {
	eps, err := mongoUtils.GetPageNodesBySectionId(client, ctx, sectionId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, eps)
	}
}
func GetPageByEpisode(c *gin.Context, client *mongo.Client, ctx context.Context, pageId string, episodeNum int) {
	pageNodes, err := mongoUtils.GetPageByEpisode(client, ctx, pageId, episodeNum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, pageNodes)
	}
}