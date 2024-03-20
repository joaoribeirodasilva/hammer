package classes

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/joaoribeirodasilva/hammer/database"
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
	user        *models.Users
	NumArticles int
}

type Users struct {
	Db            *database.Database
	Current       int
	Max           int
	TotalArticles int
	Articles      *Articles
	currentUser   int
	first_names   []FirstName
	surnames      []Surname
	domains       []Domain
	users         []CreatedUsers
}

func (a *Users) InitUsers(max int, db *database.Database, articles *Articles) {

	a.Db = db
	a.Articles = articles
	a.Max = max
	a.TotalArticles = 0
	a.users = make([]CreatedUsers, 0)
	a.loadNames()
	a.loadSurnames()
	a.loadDomains()
	a.Current = 0
}

func (a *Users) CreateUser() *CreatedUsers {

	if a.IsMax() {
		return nil
	}

	cUser := &CreatedUsers{
		user:        nil,
		NumArticles: 0,
	}

	usr := new(models.Users)

	usr.FirstName = a.firstName()
	usr.LastName = a.lastName()
	usr.Email = strings.ToLower(fmt.Sprintf("%s.%s@%s", usr.FirstName, usr.LastName, a.domain()))

	result := a.Db.Conn.Save(usr)
	if result.Error != nil {
		log.Fatalf("failed to create User record. ERR: %s", result.Error.Error())
	}

	cUser.user = usr
	a.users = append(a.users, *cUser)
	a.Current++

	return cUser
}

func (a *Users) GetNextUser() *CreatedUsers {

	user := &a.users[a.currentUser]
	a.currentUser++
	if a.currentUser >= len(a.users) {
		a.currentUser = 0
	}

	return user
}

func (a *Users) CreateUserArticle(user *CreatedUsers) {

	if a.IsMaxArticles(user) {
		return
	}

	a.Articles.CreateArticle(user.user.ID)
	user.NumArticles++
	a.TotalArticles++
}

func (a *Users) IsMax() bool {

	return len(a.users) >= a.Max
}

func (a *Users) IsMaxArticles(user *CreatedUsers) bool {

	return user.NumArticles >= a.Articles.Max
}

func (a *Users) firstName() string {

	nameIndex := rand.Intn(len(a.first_names)-1) + 1
	return a.first_names[nameIndex].name
}

func (a *Users) lastName() string {

	nameIndex := rand.Intn(len(a.surnames)-1) + 1
	return strings.ToLower(a.surnames[nameIndex].name)
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
		filePath = fmt.Sprintf("%s", PATH_SURNAMES)
	} else if fileType == "first_names" {
		filePath = fmt.Sprintf("%s", PATH_FIRST_NAMES)
	} else if fileType == "domains" {
		filePath = fmt.Sprintf("%s", PATH_DOMAINS)
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
