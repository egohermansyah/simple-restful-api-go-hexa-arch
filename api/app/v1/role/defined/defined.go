package defined

import (
	"simple-restful-api-go-hexa-arch/business/role/defined"
	"time"
)

type InsertRequest struct {
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc"`
}

type DefaultResponse struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewDefaultResponse(role *defined.Role) *DefaultResponse {
	return &DefaultResponse{
		role.Id,
		role.Name,
		role.Desc,
		role.Created,
		role.Modified}
}

func NewDefaultResponseList(roles []defined.Role) []*DefaultResponse {
	var result []*DefaultResponse
	for _, role := range roles {
		result = append(result, NewDefaultResponse(&role))
	}
	return result
}
