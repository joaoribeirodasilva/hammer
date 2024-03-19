package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/joaoribeirodasilva/hammer/models"
)

const (
	PATH_FIRST_NAMES = "./data/first-names.csv"
	PATH_SURNAMES    = "./data/surnames.csv"
	PATH_DOMAINS     = "./data/domains.csv"
)

type FirstName struct {
	name string
}

type Surname struct {
	name string
}

type Domain struct {
	name string
}

type CreatedUsers struct {
	ID          int
	NumArticles int
}

type Users struct {
	first_names []FirstName
	surnames    []Surname
	domains     []Domain
	users       []CreatedUsers
	max         int
	maxArticles int
	created     int
}

func (a *Users) InitUsers() {

	a.loadNames()
	a.loadSurnames()
	a.loadDomains()
	a.created = 0
}

func (a *Users) CreateUser() *models.Users {

	usr := new(models.Users)

	usr.FirstName = a.firstName()
	usr.LastName = a.lastName()
	usr.Email = fmt.Sprintf("%s.%s@%s", usr.FirstName, usr.LastName, a.domain())
	a.created++

	return usr
}

func (a *Users) GetRandomUser() *CreatedUsers {

	userIndex := rand.Intn(len(a.first_names) - 1)
	return &a.users[userIndex]
}

func (a *Users) AddUser(userId int) bool {

	if len(a.users) >= a.max {
		return false
	}

	user := CreatedUsers{ID: userId, NumArticles: 0}
	a.users = append(a.users, user)

	return true
}

func (a *Users) IsMax() bool {

	return len(a.users) >= a.max
}

func (a *Users) IsMaxArticles(userId int) bool {

	for idx := range a.users {
		if a.users[idx].ID == userId && a.users[idx].NumArticles < a.maxArticles {
			return false
		}
	}

	return true
}

func (a *Users) AddUserArticle(userId int) bool {

	for idx := range a.users {
		if a.users[idx].ID == userId && a.users[idx].NumArticles < a.maxArticles {
			a.users[idx].NumArticles++
			return true
		}
	}

	return false
}

func (a *Users) firstName() string {

	nameIndex := rand.Intn(len(a.first_names)-1) + 1
	return a.first_names[nameIndex].name
}

func (a *Users) lastName() string {

	nameIndex := rand.Intn(len(a.surnames)-1) + 1
	return a.surnames[nameIndex].name
}

func (a *Users) domain() string {

	nameIndex := rand.Intn(len(a.domains)-1) + 1
	return a.domains[nameIndex].name
}

func (a *Users) loadNames() error {

	return a.loadCSV("first_names")
}

func (a *Users) loadSurnames() error {

	return a.loadCSV("surnames")
}

func (a *Users) loadDomains() error {

	return a.loadCSV("domains")
}

func (a *Users) loadCSV(fileType string) error {

	filePath := ""
	if fileType == "surnames" {
		filePath = fmt.Sprintf("%s/%s", a.getBinPath(), PATH_SURNAMES)
	} else if fileType == "first_names" {
		filePath = fmt.Sprintf("%s/%s", a.getBinPath(), PATH_FIRST_NAMES)
	} else if fileType == "domains" {
		filePath = fmt.Sprintf("%s/%s", a.getBinPath(), PATH_DOMAINS)
	}

	log.Printf("Reading data file '%s'\n", filePath)

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	for i, line := range records {
		if i > 0 {
			for _, field := range line {
				if fileType == "surnames" {
					surname := Surname{
						name: field,
					}
					a.surnames = append(a.surnames, surname)
				} else if fileType == "first_names" {
					first_name := FirstName{
						name: field,
					}
					a.first_names = append(a.first_names, first_name)
				} else if fileType == "domains" {
					domain := Domain{
						name: field,
					}
					a.domains = append(a.domains, domain)
				}
			}
		}
	}

	if fileType == "surnames" {
		log.Printf("%d surnames read from file\n", len(a.surnames))
	} else if fileType == "first_names" {
		log.Printf("%d first names read from file\n", len(a.first_names))
	} else if fileType == "domains" {
		log.Printf("%d domains read from file\n", len(a.domains))
	}

	return nil
}

func (a *Users) getBinPath() string {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
