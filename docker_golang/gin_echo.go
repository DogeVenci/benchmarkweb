package main

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type result struct {
	ID                bson.ObjectId `bson:"_id,omitempty"`
	Id                string        `bson:"id"`
	Name              string        `bson:"name"`
	Slug              string        `bson:"slug"`
	Created_at        string        `bson:"created_at"`
	Released_at       string        `bson:"released_at"`
	Views             string        `bson:"views"`
	Interests         string        `bson:"interests"`
	Poster_url        string        `bson:"poster_url"`
	Cover_url         string        `bson:"cover_url"`
	Is_hard_subtitled string        `bson:"is_hard_subtitled,omitempty"`
	Brand             string        `bson:"brand,omitempty"`
	Duration_in_ms    string        `bson:"duration_in_ms,omitempty"`
	Is_censored       string        `bson:"is_censored,omitempty"`
	Rating            string        `bson:"rating,omitempty"`
	Likes             string        `bson:"likes,omitempty"`
	Dislikes          string        `bson:"dislikes,omitempty"`
	Downloads         string        `bson:"downloads,omitempty"`
	Monthly_rank      string        `bson:"monthly_rank,omitempty"`
	Brand_id          string        `bson:"brand_id,omitempty"`
	Is_banned_in      string        `bson:"is_banned_in,omitempty"`
	Created_at_unix   string        `bson:"created_at_unix,omitempty"`
	Released_at_unix  string        `bson:"released_at_unix,omitempty"`
	Filename          string        `bson:"filename,omitempty"`
	Extension         string        `bson:"extension,omitempty"`
	CreatedAt         string        `bson:"raticreatedAtng,omitempty"`
	UpdateAt          string        `bson:"updateAt,omitempty"`
}

func main() {
	session, err := mgo.Dial("mongodb://root:example@mongo:27017/")
	if err != nil {
		panic(err)
	}
	c := session.DB("test").C("items")
	results := []result{}
	// item := result{}
	err = c.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// time.Sleep(time.Duration(500) * time.Millisecond)
		c.JSON(200,
			results,
		)
	})
	gin.SetMode(gin.ReleaseMode)
	r.Run() // listen and serve on 0.0.0.0:8080
}
