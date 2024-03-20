package classes

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/joaoribeirodasilva/hammer/database"
	"github.com/joaoribeirodasilva/hammer/models"
)

type Accesses struct {
	Db       *database.Database
	Current  int
	Max      int
	serverId int
}

func (a *Accesses) InitAccesses(serverId int, max int, Db *database.Database) {

	a.Db = Db
	a.serverId = serverId
	a.Max = max
}

func (a *Accesses) CreateAccess() {

	acc := new(models.Accesses)

	acc.Ip = a.ip()
	acc.OriginID = uint(a.serverId)

	result := a.Db.Conn.Create(acc)
	if result.Error != nil {
		log.Fatalf("failed to create Access record. ERR: %s", result.Error.Error())
	}

	a.Current++

}

func (a *Accesses) ip() string {

	b1 := 0
	b4 := 0

	for b1 == 0 || b1 == 255 || b1 == 254 {
		b1 = rand.Intn(256-1) + 1
	}

	b2 := rand.Intn(256-1) + 1
	b3 := rand.Intn(256-1) + 1

	for b4 == 0 || b4 == 255 || b4 == 254 {
		b4 = rand.Intn(256-1) + 1
	}

	return fmt.Sprintf("%d.%d.%d.%d", b1, b2, b3, b4)
}
