package main

import (
	"fmt"
	"context"
	"log"
	"net/http"

	// "github.com/google/uuid"
	"github.com/mlaikhram/popit/server/mongoUtils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func getShow(c *gin.Context, client *mongo.Client, ctx context.Context) {
	log.Println("show param " + c.Param("showId"))
	show, err := getShowById(client, ctx, c.Param("showId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, show)
	}
}

func searchShow(c *gin.Context, client *mongo.Client, ctx context.Context) {
	show, err := getShows(client, ctx, c.Param("term"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, show)
	}
}

func getEpisodesByShowId(c *gin.Context, client *mongo.Client, ctx context.Context) {
	eps, err := getEpisodes(client, ctx, c.Param("showId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, eps)
	}
}

func getPageNodesBySectionId(c *gin.Context, client *mongo.Client, ctx context.Context) {

}

func main() {
	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	router := gin.Default()
	router.GET("/shows/show/:showId", func(c *gin.Context) {
		getShow(c, client, ctx)
	})

	router.GET("/shows/search/:term", func(c *gin.Context) {
		searchShow(c, client, ctx)
	})

	router.GET("/shows/episodes/:showId", func(c *gin.Context) {
		getEpisodesByShowId(c, client, ctx)
	})

	// show := Show{
	// 	Name: "Attack on Titan",
	// 	Aliases: []string{"Shingeki no Kyojin"},
	// }

	// episode := Episode{
	// 	Number: 1,
	// 	Key: "S1E1",
	// 	Name: "To You, in 2000 Years: The Fall of Shiganshina, Part 1",
	// }

	// addShow(client, ctx, show, episode)

	// eps := []Episode{
	// 	{
	// 		Number: 2,
	// 		Key: "S1E2",
	// 		Name: "That Day: The Fall of Shiganshina, Part 2",
	// 	},
	// 	{
	// 		Number: 3,
	// 		Key: "S1E3",
	// 		Name: "A Dim Light Amid Despair: Humanity's Comeback, Part 1",
	// 	},
	// 	{
	// 		Number: 4,
	// 		Key: "S1E4",
	// 		Name: "The Night of the Closing Ceremony: Humanity's Comeback, Part 2",
	// 	},
	// 	{
	// 		Number: 5,
	// 		Key: "S1E5",
	// 		Name: "First Battle: The Struggle for Trost, Part 1",
	// 	},
	// }

	// err = addEpisodes(client, ctx, "625798abc15bd223b8a9eeae", eps)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// page := Page {
	// 	ShowId: "625798abc15bd223b8a9eeae",
	// 	Tags: map[int][]string {
	// 		1: {"Characters"},
	// 	},
	// 	InitialEpisodeNum: 1,
	// }

	// nodes := []PageNode {
	// 	{
	// 		SectionID: uuid.New().String(),
	// 		EpisodeNum: 1,
	// 		Type: HEADER,
	// 		Title: "Eren Yeager",
	// 		Content: PageNodeContent {
	// 			ProfileImages: []Pair {
	// 				{
	// 					Key: "Young Eren",
	// 					Value: "test_img",
	// 				},
	// 			},
	// 			ProfileAttributes: []Pair {
	// 				{
	// 					Key: "Gender",
	// 					Value: "Male",
	// 				},
	// 				{
	// 					Key: "Hometown",
	// 					Value: "Shiganshina",
	// 				},
	// 				{
	// 					Key: "Status",
	// 					Value: "Alive",
	// 				},
	// 			},
	// 			Text: "Eren Yeager is the main character. I don't feel like writing any more than that.",
	// 		},
	// 	},
	// 	{
	// 		SectionID: uuid.New().String(),
	// 		EpisodeNum: 1,
	// 		Type: SUMMARY,
	// 		Title: "Appearance",
	// 		Content: PageNodeContent {
	// 			Text: "He's short with short, black hair and green eyes.",
	// 		},
	// 	},
	// 	{
	// 		SectionID: uuid.New().String(),
	// 		EpisodeNum: 1,
	// 		Type: SUMMARY,
	// 		Title: "Personality",
	// 		Content: PageNodeContent {
	// 			Text: "He angry.",
	// 		},
	// 	},
	// 	{
	// 		SectionID: uuid.New().String(),
	// 		EpisodeNum: 1,
	// 		Type: SUMMARY,
	// 		Title: "Childhood",
	// 		Content: PageNodeContent {
	// 			Text: "A titan ate his mom lol.",
	// 		},
	// 	},
	// }

	// id, err := addPage(client, ctx, page, nodes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("id: " + id)

	// pageArr, err := getPageById(client, ctx, id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(pageArr[0])

	shows, err := getShows(client, ctx, "shingeki no Kyojin")
	// shows, err := getShows(client, ctx, "attack on titan")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shows)

	episodes, err := getEpisodes(client, ctx, "625798abc15bd223b8a9eeae")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(episodes)

	router.Run("localhost:8080")
}
