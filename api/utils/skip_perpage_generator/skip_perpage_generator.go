package skip_perpage_generator

import (
	"strconv"
)

type SkipPerPage struct {
	Skip    int
	PerPage int
}

type ISkipPerPageGenerator interface {
	SkipPerPageGenerator(page string, perPage string) (SkipPerPage, error)
}

func NewSkipPerPageGenerator(page string, perPage string) (SkipPerPage, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}
	perPageInt, err := strconv.Atoi(perPage)
	if err != nil || perPageInt <= 0 {
		perPageInt = 10
	}
	skip := (pageInt * perPageInt) - perPageInt
	return SkipPerPage{Skip: skip, PerPage: perPageInt}, nil
}
