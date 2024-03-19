package main

import (
	"fmt"
	"math/rand"

	"github.com/joaoribeirodasilva/hammer/models"
)

type Accesses struct {
	serverId int
}

func InitAccesses(serverId int) *Accesses {

	acc := new(Accesses)
	acc.serverId = serverId
	return acc
}

func (a *Accesses) CreateAccess() *models.Accesses {

	acc := new(models.Accesses)

	acc.Ip = a.ip()
	acc.OriginID = uint(a.serverId)

	return acc

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
