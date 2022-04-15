package endpoints

import (
	"context"
	"net/http"

	"popit/mongoUtils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetShowById(c *gin.Context, client *mongo.Client, ctx context.Context, showId string) {
	show, err := mongoUtils.GetShowById(client, ctx, showId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, show)
	}
}

func GetShowByTitle(c *gin.Context, client *mongo.Client, ctx context.Context, title string) {
	show, err := mongoUtils.GetShowByTitle(client, ctx, title)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, show)
	}
}

func GetEpisodesByShowId(c *gin.Context, client *mongo.Client, ctx context.Context, showId string) {
	eps, err := mongoUtils.GetEpisodesByShowId(client, ctx, showId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, eps)
	}
}
