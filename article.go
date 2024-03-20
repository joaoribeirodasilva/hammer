package main

import (
	"log"
	"math/rand"

	"github.com/XANi/loremipsum"
	"github.com/joaoribeirodasilva/hammer/database"
	"github.com/joaoribeirodasilva/hammer/models"
)

type Articles struct {
	Db                   *database.Database
	Max                  int
	loremIpsumGenerator  *loremipsum.LoremIpsum
	maxTitleWords        int
	maxContentParagraphs int
}

func (a *Articles) InitArticles(maxTitleWords int, maxContentParagraphs int, Db *database.Database) {
	a.Db = Db
	a.loremIpsumGenerator = loremipsum.New()
}

func (a *Articles) CreateArticle(userId uint) {

	art := new(models.Articles)

	art.UserID = userId
	art.Title = a.randomTitle()
	art.ContentText = a.randomContent()

	result := a.Db.Conn.Save(art)
	if result.Error != nil {
		log.Fatalf("failed to create Article record. ERR: %s", result.Error.Error())
	}
}

func (a *Articles) randomTitle() string {

	numWords := rand.Intn(a.maxTitleWords-1) + 1
	return a.loremIpsumGenerator.Words(numWords)
}

func (a *Articles) randomContent() string {

	numParagraphs := rand.Intn(a.maxContentParagraphs-1) + 1
	return a.loremIpsumGenerator.Paragraphs(numParagraphs)
}
