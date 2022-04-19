package models

type Show struct {
	ID       string         `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string         `json:"name" bson:"name"`
	Aliases  []string       `json:"aliases" bson:"aliases,omitempty"`
	Synopsis string         `json:"synopsis" bson:"synopsis"`
	Images   map[int]string `json:"images,omitempty" bson:"images,omitempty"`
}

type Episode struct {
	ID                   string `json:"id,omitempty" bson:"_id,omitempty"`
	ShowId               string `json:"showId,omitempty" bson:"showId"`
	Number               int    `json:"number" bson:"number"`
	Key                  string `json:"key" bson:"key"`
	Name                 string `json:"name" bson:"name"`
	IsSpecial            bool   `json:"isSpecial,omitempty" bson:"isSpecial,omitempty"`
	SuggestedPrevEpisode int    `json:"suggestedPrevEpisode,omitempty" bson:"suggestedPrevEpisode,omitempty"`
}

type PageNodeType int

const (
	HEADER = iota
	SUMMARY
	TABLE
)

type Pair struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

type PageNodeContent struct {
	ProfileImages     []Pair     `json:"profileImages,omitempty" bson:"profileImages,omitempty"`
	ProfileAttributes []Pair     `json:"profileAttributes,omitempty" bson:"profileAttributes,omitempty"`
	Text              string     `json:"text,omitempty" bson:"text,omitempty"`
	Table             [][]string `json:"table,omitempty" bson:"table,omitempty"`
}

type PageNode struct {
	ID         string          `json:"id,omitempty" bson:"_id,omitempty"`
	ShowId     string          `json:"showId,omitempty" bson:"showId"`
	SectionID  string          `json:"sectionId,omitempty" bson:"sectionId,omitempty"`
	EpisodeNum int             `json:"episodeNum,omitempty" bson:"episodeNum,omitempty"`
	Type       PageNodeType    `json:"type" bson:"type"`
	Title      string          `json:"title" bson:"title"`
	Content    PageNodeContent `json:"content" bson:"content"`
}

type Page struct {
	ID                string           `json:"id,omitempty" bson:"_id,omitempty"`
	ShowId            string           `json:"showId,omitempty" bson:"showId"`
	Tags              map[int][]string `json:"tags,omitempty" bson:"tags,omitempty"`
	InitialEpisodeNum int              `json:"initialEpisodeNum" bson:"initialEpisodeNum"`
	SectionIDs        []string         `json:"sectionIds" bson:"sectionIds"`
}
