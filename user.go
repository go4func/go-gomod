package doer

type User struct {
	Doer Doer
}

func (u User) Use() {
	u.Doer.Do(1, "two")
}
