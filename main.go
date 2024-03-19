package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Limits struct {
	Max     int
	Current int
}

type Main struct {
	DsnMaster string
	DsnClient string
	IsClient  bool
	ServerID  int
	Users     Limits
	Articles  Limits
	Accesses  Limits
}

func main() {

	m := &Main{IsClient: false}
	m.readEnv()
	m.cmdArgs()

}

func (m *Main) readEnv() {

	fmt.Println("\nKafka Replication Hammer v0.1")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	m.DsnMaster = os.Getenv("DSN_MASTER")
	m.DsnClient = os.Getenv("DSN_CLIENT")

	if m.DsnMaster == "" && !m.IsClient {
		log.Fatal("the program is running as master but no DSN_MASTER was found inside the .env file")
	} else if m.DsnClient == "" && m.IsClient {
		log.Fatal("the program is running as client but no DSN_CLIENT was found inside the .env file")
	}

	m.run()
}

func (m *Main) run() {

}

func (m *Main) cmdArgs() {

	flag.BoolVar(&m.IsClient, "client", false, "producer is a client producer")
	flag.IntVar(&m.ServerID, "c", 0, "client producer id")
	flag.IntVar(&m.Users.Max, "u", 10, "maximum number of users to create. Default(10)")
	flag.IntVar(&m.Articles.Max, "apu", 10, "maximum number of articles per user to create. Default(10)")
	flag.IntVar(&m.Accesses.Max, "acc", 100, "maximum number of accesses to create. Default(100)")

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

}
