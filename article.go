package main

import (
	"math/rand"

	"github.com/XANi/loremipsum"
	"github.com/joaoribeirodasilva/hammer/models"
)

type Articles struct {
	loremIpsumGenerator  *loremipsum.LoremIpsum
	maxTitleWords        int
	maxContentParagraphs int
}

func (a *Articles) InitArticles(maxTitleWords int, maxContentParagraphs int) {
	a.loremIpsumGenerator = loremipsum.New()
}

func (a *Articles) CreateArticle() *models.Articles {

	art := new(models.Articles)

	art.Title = a.randomTitle()
	art.ContentText = a.randomContent()
	return art
}

func (a *Articles) randomTitle() string {

	numWords := rand.Intn(a.maxTitleWords-1) + 1
	return a.loremIpsumGenerator.Words(numWords)
}

func (a *Articles) randomContent() string {

	numParagraphs := rand.Intn(a.maxContentParagraphs-1) + 1
	return a.loremIpsumGenerator.Paragraphs(numParagraphs)
}
