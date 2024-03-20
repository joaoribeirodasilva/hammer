package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joaoribeirodasilva/hammer/classes"
	"github.com/joaoribeirodasilva/hammer/database"
	"github.com/joho/godotenv"
)

type Limits struct {
	Max     int
	Current int
}

type Main struct {
	Db           *database.Database
	DsnMaster    string
	DsnClient    string
	EngineMaster string
	EngineClient string
	IsClient     bool
	IsHelp       bool
	ServerID     int
	Users        Limits
	Articles     Limits
	Accesses     Limits
}

func main() {

	m := &Main{IsClient: false}
	m.readEnv()
	m.cmdArgs()
	m.run()
}

func (m *Main) readEnv() {

	fmt.Println("\nKafka Replication Hammer v0.1")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	m.DsnMaster = os.Getenv("DSN_MASTER")
	m.DsnClient = os.Getenv("DSN_CLIENT")
	m.EngineMaster = os.Getenv("ENGINE_MASTER")
	m.EngineClient = os.Getenv("ENGINE_CLIENT")

	if m.DsnMaster == "" && !m.IsClient {
		log.Fatal("the program is running as master but no DSN_MASTER was found inside the .env file")
	} else if m.DsnClient == "" && m.IsClient {
		log.Fatal("the program is running as client but no DSN_CLIENT was found inside the .env file")
	}

	if m.EngineMaster == "" && !m.IsClient {
		log.Fatal("the program is running as master but no ENGINE_MASTER was found inside the .env file")
	} else if m.EngineClient == "" && m.IsClient {
		log.Fatal("the program is running as client but no ENGINE_CLIENT was found inside the .env file")
	}

	m.Db = database.New()

	if !m.IsClient {
		m.Db.Connect(m.DsnMaster, m.EngineMaster)
	} else {
		m.Db.Connect(m.DsnClient, m.EngineClient)
	}
}

func (m *Main) run() {

	users := &classes.Users{}
	articles := &classes.Articles{}
	accesses := &classes.Accesses{}

	if !m.IsClient {
		articles.InitArticles(5, 10, m.Articles.Max, m.Db)
		users.InitUsers(m.Users.Max, m.Db, articles)
	}

	accesses.InitAccesses(m.ServerID, m.Accesses.Max, m.Db)

	notFinish := true

	for notFinish {

		if accesses.Current < accesses.Max {

			accesses.CreateAccess()
			m.Accesses.Current++

		}

		if !m.IsClient {

			var usr *classes.CreatedUsers

			if users.Current < users.Max {

				usr = users.CreateUser()
				users.CreateUserArticle(usr)

			} else {

				usr = users.GetNextUser()
				users.CreateUserArticle(usr)

			}

			users.CreateUserArticle(usr)

			notFinish = accesses.Current < accesses.Max || users.Current < users.Max || users.TotalArticles < m.Articles.Max*m.Users.Max
		} else {
			notFinish = accesses.Current < accesses.Max
		}

		fmt.Printf("\rAdded %d accesses, %d users, %d articles", accesses.Current, users.Current, users.TotalArticles)
	}
	fmt.Println("")

}

func (m *Main) cmdArgs() {

	flag.BoolVar(&m.IsClient, "client", false, "producer is a client producer")
	flag.IntVar(&m.ServerID, "c", 0, "client producer id")
	flag.IntVar(&m.Users.Max, "u", 10, "maximum number of users to create. Default(10)")
	flag.IntVar(&m.Articles.Max, "apu", 10, "maximum number of articles per user to create. Default(10)")
	flag.IntVar(&m.Accesses.Max, "acc", 100, "maximum number of accesses to create. Default(100)")
	flag.BoolVar(&m.IsHelp, "help", false, "prints this help")

	flag.Parse()

	if !m.IsClient && m.ServerID > 0 {
		log.Fatalf("the program is not running as client therefore server id can't be %d", m.ServerID)
	} else if m.IsClient && m.ServerID == 0 {
		log.Fatalf("the program is running as master therefore server id can't be %d", m.ServerID)
	} else if m.Users.Max == 0 {
		log.Fatalf("the number of users to create can't be 0")
	} else if m.Articles.Max == 0 {
		log.Fatalf("the number of articles per user to create can't be 0")
	}

	if m.IsHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

}
