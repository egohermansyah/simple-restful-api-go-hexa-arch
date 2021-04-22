package role

import (
	"github.com/labstack/echo/v4"
	controllerRoleDefined "simple-restful-api-go-hexa-arch/api/app/v1/role/defined"
	"simple-restful-api-go-hexa-arch/api/utils/custom_validator"
	"simple-restful-api-go-hexa-arch/api/utils/response"
	"simple-restful-api-go-hexa-arch/api/utils/skip_perpage_generator"
	"simple-restful-api-go-hexa-arch/business/role"
	serviceRoleDefined "simple-restful-api-go-hexa-arch/business/role/defined"
	"net/http"
)

type Controller struct {
	service role.IService
}

func NewController(service role.IService) *Controller {
	return &Controller{service}
}

func (controller *Controller) Insert(c echo.Context) error {
	bodyRequest := new(controllerRoleDefined.InsertRequest)
	if err := c.Bind(bodyRequest); err != nil {
		errors := response.BuildErrors(err, err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	if err := c.Validate(bodyRequest); err != nil {
		errors := custom_validator.BuildErrorBodyRequestValidator(err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	data := serviceRoleDefined.NewRole(
		bodyRequest.Name,
		bodyRequest.Desc)
	result, err := controller.service.Insert(data)
	if err != nil {
		errors := response.BuildErrors(err, err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	resultResponse := controllerRoleDefined.NewDefaultResponse(result)
	return c.JSON(http.StatusCreated, response.NewResponse(response.Mapping["created"], response.Errors{}, "", resultResponse))
}

func (controller *Controller) FindById(c echo.Context) error {
	id := c.Param("id")
	result, err := controller.service.FindById(id)
	if err != nil {
		errors := response.BuildErrors(err, err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	resultResponse := controllerRoleDefined.NewDefaultResponse(result)
	return c.JSON(http.StatusOK, response.NewResponse(response.Mapping["ok"], response.Errors{}, "", resultResponse))
}

func (controller *Controller) UpdateById(c echo.Context) error {
	id := c.Param("id")
	bodyRequest := new(controllerRoleDefined.InsertRequest)
	if err := c.Bind(bodyRequest); err != nil {
	}
	result, err := controller.service.UpdateById(id, bodyRequest.Name, bodyRequest.Desc)
	if err != nil {
		errors := response.BuildErrors(err, err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	resultResponse := controllerRoleDefined.NewDefaultResponse(result)
	return c.JSON(http.StatusOK, response.NewResponse(response.Mapping["ok"], response.Errors{}, "", resultResponse))
}

func (controller *Controller) DeleteById(c echo.Context) error {
	id := c.Param("id")
	err := controller.service.DeleteById(id)
	if err != nil {
		errors := response.BuildErrors(err, err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	resultResponse := controllerRoleDefined.DefaultResponse{}
	return c.JSON(http.StatusOK, response.NewResponse(response.Mapping["ok"], response.Errors{}, "", resultResponse))
}

func (controller *Controller) List(c echo.Context) error {
	page := c.QueryParam("page")
	perPage := c.QueryParam("per_page")
	skipPerPage, err := skip_perpage_generator.NewSkipPerPageGenerator(page, perPage)
	if err != nil {
		errors := response.BuildErrors(err, err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	result, err := controller.service.List(skipPerPage.Skip, skipPerPage.PerPage)
	if err != nil {
		errors := response.BuildErrors(err, err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(response.Mapping["badRequest"], errors, "", nil))
	}
	resultResponse := controllerRoleDefined.NewDefaultResponseList(result)
	return c.JSON(http.StatusOK, response.NewResponse(response.Mapping["ok"], response.Errors{}, "", resultResponse))
}
