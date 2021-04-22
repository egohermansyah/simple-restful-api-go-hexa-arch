package defined

import "time"

type Role struct {
	Id       string
	Name     string
	Desc     string
	Created  time.Time
	Modified time.Time
}

func NewRole(name string, desc string) Role {
	return Role{
		Name: name,
		Desc: desc,
	}
}
