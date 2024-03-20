package classes

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

func (a *Articles) InitArticles(maxTitleWords int, maxContentParagraphs int, max int, Db *database.Database) {
	a.Db = Db
	a.Max = max
	a.maxTitleWords = maxTitleWords
	a.maxContentParagraphs = maxContentParagraphs
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

	a.loremIpsumGenerator = loremipsum.New()
	numWords := rand.Intn(a.maxTitleWords-1) + 1
	return a.loremIpsumGenerator.Words(numWords)
}

func (a *Articles) randomContent() string {

	a.loremIpsumGenerator = loremipsum.New()
	numParagraphs := rand.Intn(a.maxContentParagraphs-1) + 1
	return a.loremIpsumGenerator.Paragraphs(numParagraphs)
}
