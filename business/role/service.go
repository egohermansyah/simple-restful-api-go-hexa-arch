package role

import (
	roleServiceDefined "simple-restful-api-go-hexa-arch/business/role/defined"
	roleRepositoryDefined "simple-restful-api-go-hexa-arch/modules/role/defined"
)

type IService interface {
	Insert(role roleServiceDefined.Role) (*roleServiceDefined.Role, error)
	FindById(id string) (*roleServiceDefined.Role, error)
	UpdateById(id string, name string, desc string) (*roleServiceDefined.Role, error)
	DeleteById(id string) error
	List(skip int, perPage int) ([]roleServiceDefined.Role, error)
}

type Service struct {
	Repository roleRepositoryDefined.IRepository
}

func NewService(repository roleRepositoryDefined.IRepository) IService {
	return &Service{
		Repository: repository,
	}
}

func (service *Service) Insert(role roleServiceDefined.Role) (*roleServiceDefined.Role, error) {
	data := roleServiceDefined.NewRole(
		role.Name,
		role.Desc,
	)
	return service.Repository.Insert(data)
}

func (service *Service) FindById(id string) (*roleServiceDefined.Role, error) {
	return service.Repository.FindById(id)
}

func (service *Service) UpdateById(id string, name string, desc string) (*roleServiceDefined.Role, error) {
	return service.Repository.UpdateById(id, name, desc)
}

func (service *Service) DeleteById(id string) error {
	return service.Repository.DeleteById(id)
}

func (service *Service) List(skip int, perPage int) ([]roleServiceDefined.Role, error) {
	return service.Repository.List(skip, perPage)
}
