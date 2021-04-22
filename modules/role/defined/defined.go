package defined

import "simple-restful-api-go-hexa-arch/business/role/defined"

type IRepository interface {
	Insert(role defined.Role) (*defined.Role, error)
	FindById(id string) (*defined.Role, error)
	UpdateById(id string, name string, desc string) (*defined.Role, error)
	DeleteById(id string) error
	List(skip int, perPage int) ([]defined.Role, error)
}
