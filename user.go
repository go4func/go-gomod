package doer

import "log"

type User struct {
	Doer Doer
}

func (u User) Use() {
	u.Doer.Do(1, "two")
}

func (u User) AnotherUse() {
	log.Println("version 2")
}
