package doer

import (
	"log"
)

type (
	User struct {
		Doer Doer
	}
)

func (u User) AnotherUse() {
	log.Println("")
}
