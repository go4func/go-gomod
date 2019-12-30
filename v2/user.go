package doer

import (
	"log"

	doer "github.com/nthlongtma/go-gomock/v2"
)

type (
	User struct {
		Doer doer.Doer
	}
)

func (u User) AnotherUse() {
	log.Println("")
}
