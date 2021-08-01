package api

import (
	"evermos-test/config/env"
	"evermos-test/database/entity"
	"evermos-test/helper"
	"evermos-test/http/interfaces"
	"evermos-test/http/request"
	"evermos-test/http/response"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type CustomerHandler struct {
	Helper         					helper.HTTPHelper
	Config         					env.Config
	CustomerRepository 				interfaces.CustomerInterface
	InventoryAdjustmentRepository 	interfaces.InventoryAdjustmentInterface
}

// @Tags Customer
// @Description Create Customer
// @ID Create Customer
// @Accept  json
// @Produce  json
// @Param CreateCustomer body request.CreateCustomerRequest true "create customer info"
// @Success 200 {object} response.CustomerSuccessResponse
// @Failure 400 {object} response.CustomerFailedResponse
// @Failure 404 {object} response.CustomerFailedResponse
// @Router /customer [post]
func (_h *CustomerHandler) CreateCustomer(c echo.Context) error {

	var (
		errResults []string
		err        error
		input      request.CreateCustomerRequest
	)

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {

		return _h.Helper.SendInputValidationError(c, err.(validator.ValidationErrors))
	}

	// Begin Add Your Additional Logic Here

	// End Add Your Additional Logic Here



	//begin save to DB
	var newMainEntityToSave entity.Customer
	errResults = newMainEntityToSave.ValidateBeforeCreate(input)

	customerId := helper.GenerateBsonObjectId()
	newMainEntityToSave.Id = customerId
	if len(errResults) > 0 {
		return _h.Helper.SendBadRequest(c, errResults)
	}
	_, err = _h.CustomerRepository.Create(&newMainEntityToSave)

	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)

	}

	// End Save To DB

	return _h.Helper.SendSuccess(c, nil)

}

// @Tags Customer
// @Description Update Customer
// @ID Update Customer
// @Accept  json
// @Produce  json
// @Param CreateCustomer body request.UpdateCustomerRequest true "update customer info"
// @Success 200 {object} response.CustomerSuccessResponse
// @Failure 400 {object} response.CustomerFailedResponse
// @Failure 404 {object} response.CustomerFailedResponse
// @Router /customer/{id} [put]
func (_h *CustomerHandler) UpdateCustomer(c echo.Context) error {

	var (
		errResults []string
		err   error
		input request.UpdateCustomerRequest
	)

	id := c.Param("id")

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c,errResults)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {
		return _h.Helper.SendInputValidationError(c, err.(validator.ValidationErrors))
	}

	//begin save to DB
	entityToUpdate, err := _h.CustomerRepository.FindById(id)
	if entityToUpdate == nil {

		errResults = append(errResults, helper.ErrorNotFound(id))
		return _h.Helper.SendBadRequest(c, errResults)
	}

	if err != nil {
		isNotFoundError := helper.IsNotFoundErrorValidation(err.Error())
		if isNotFoundError {

			errResults = append(errResults, helper.ErrorNotFound(id))
			return _h.Helper.SendBadRequest(c, errResults)
		}

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}

	// Begin Add Your Additional Logic Here

	// End Add Your Additional Logic Here
	entityToUpdate.ValidateBeforeUpdate(input)

	_, err = _h.CustomerRepository.Update(id, entityToUpdate)

	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}

	// End Save To DB

	result, errStr := _h.CustomerFindById(id)

	if result == nil || errStr != "" {

		errResults = append(errResults, errStr)
		return _h.Helper.SendBadRequest(c, errResults)
	}

	return _h.Helper.SendSuccess(c, result)

}

// @Tags Customer
// @Description Find All Customer
// @Accept  json
// @Produce  json
// @Param Searchuser body request.SearchParamCustomerRequest true "search customer info"
// @Success 200 {object} response.CustomerSuccessWithPagingResponse
// @Failure 400 {object} response.CustomerFailedResponse
// @Router /customer/find-all [post]
func (_h *CustomerHandler) FindAll(c echo.Context) error {

	var (
		errResults []string
		err   error
		input request.SearchParamWithPagingCustomerRequest
	)

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}


	if err = _h.Helper.Validate.Struct(input); err != nil {
		return _h.Helper.SendInputValidationError(c, err.(validator.ValidationErrors))
	}

	entities, err, totalRecords := _h.CustomerRepository.FindAll(input)
	errResults = helper.NotFoundValidationForSearching(err)
	if len(errResults) > 0 {

		return _h.Helper.SendBadRequest(c, errResults)
	}

	// begin parsing from entity to response
	result := &response.CustomerSearchResponse{}
	userResponseArray := []*response.CustomerResponse{}
	for _, entity := range entities {
		userResponse := &response.CustomerResponse{}

		userResponse.ParsingEntityToResponse(entity)
		userResponseArray = append(userResponseArray, userResponse)
	}

	result.GeneratePagingResponse(userResponseArray, input.Page, input.Limit, totalRecords)

	return _h.Helper.SendSuccess(c, result)
}

// @Tags Customer
// @Description Find a Customer
// @ID Find a Customer
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CustomerSuccessResponse
// @Failure 400 {object} response.CustomerFailedResponse
// @Failure 404 {object} response.CustomerFailedResponse
// @Router /customer/{id} [get]
func (_h *CustomerHandler) FindById(c echo.Context) error {

	var errResults []string
	id := c.Param("id")

	result, err := _h.CustomerFindById(id)

	if result == nil || err != "" {

		errResults = append(errResults, err)
		return _h.Helper.SendBadRequest(c, errResults)
	}

	return _h.Helper.SendSuccess(c, result)

}

func (_h *CustomerHandler) CustomerFindById(id string) (*response.CustomerResponse, string) {

	entityResult, err := _h.CustomerRepository.FindById(id)
	if entityResult == nil {
		return nil, helper.ErrorNotFound(id)
	}

	if err != nil {
		isNotFoundError := helper.IsNotFoundErrorValidation(err.Error())
		if isNotFoundError {
			return nil, helper.ErrorNotFound(id)
		}

		return nil, err.Error()
	}

	result := &response.CustomerResponse{}
	result.ParsingEntityToResponse(entityResult)

	return result, ""
}
