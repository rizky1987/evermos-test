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

type ProductHandler struct {
	Helper         helper.HTTPHelper
	Config         env.Config
	ProductRepository interfaces.ProductInterface
}

// @Tags Product
// @Description Create Product
// @ID Create Product
// @Accept  json
// @Produce  json
// @Param CreateProduct body request.CreateProductRequest true "create product info"
// @Success 200 {object} response.ProductSuccessResponse
// @Failure 400 {object} response.ProductFailedResponse
// @Failure 404 {object} response.ProductFailedResponse
// @Router /product [post]
func (_h *ProductHandler) CreateProduct(c echo.Context) error {

	var (
		errResults []string
		err        error
		input      request.CreateProductRequest
	)

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {

		return _h.Helper.SendValidationError(c, err.(validator.ValidationErrors))
	}

	// Begin Add Your Additional Logic Here

	// End Add Your Additional Logic Here


	//begin save to DB
	var newEntityToSave entity.Product
	errResults = newEntityToSave.ValidateBeforeCreate(input)

	if len(errResults) > 0 {
		return _h.Helper.SendBadRequest(c, errResults)
	}
	_, err = _h.ProductRepository.Create(&newEntityToSave)

	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)

	}

	// End Save To DB

	return _h.Helper.SendSuccess(c, nil)

}

// @Tags Product
// @Description Update Product
// @ID Update Product
// @Accept  json
// @Produce  json
// @Param CreateProduct body request.UpdateProductRequest true "update product info"
// @Success 200 {object} response.ProductSuccessResponse
// @Failure 400 {object} response.ProductFailedResponse
// @Failure 404 {object} response.ProductFailedResponse
// @Router /product/{id} [put]
func (_h *ProductHandler) UpdateProduct(c echo.Context) error {

	var (
		errResults []string
		err   error
		input request.UpdateProductRequest
	)

	id := c.Param("id")

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c,errResults)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {
		return _h.Helper.SendValidationError(c, err.(validator.ValidationErrors))
	}

	//begin save to DB
	entityToUpdate, err := _h.ProductRepository.FindById(id)
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

	_, err = _h.ProductRepository.Update(id, entityToUpdate)

	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}

	// End Save To DB

	result, errStr := _h.ProductFindById(id)

	if result == nil || errStr != "" {

		errResults = append(errResults, errStr)
		return _h.Helper.SendBadRequest(c, errResults)
	}

	return _h.Helper.SendSuccess(c, result)

}

// @Tags Product
// @Description Find All Product
// @Accept  json
// @Produce  json
// @Param Searchuser body request.SearchParamProductRequest true "search product info"
// @Success 200 {object} response.ProductSuccessWithPagingResponse
// @Failure 400 {object} response.ProductFailedResponse
// @Router /product/find-all [post]
func (_h *ProductHandler) FindAll(c echo.Context) error {

	var (
		errResults []string
		err   error
		input request.SearchParamWithPagingProductRequest
	)

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}


	if err = _h.Helper.Validate.Struct(input); err != nil {
		return _h.Helper.SendValidationError(c, err.(validator.ValidationErrors))
	}

	entities, err, totalRecords := _h.ProductRepository.FindAll(input)
	errResults = helper.NotFoundValidationForSearching(err)
	if len(errResults) > 0 {

		return _h.Helper.SendBadRequest(c, errResults)
	}

	// begin parsing from entity to response
	result := &response.ProductSearchResponse{}
	userResponseArray := []*response.ProductResponse{}
	for _, entity := range entities {
		userResponse := &response.ProductResponse{}

		userResponse.ParsingEntityToResponse(entity)
		userResponseArray = append(userResponseArray, userResponse)
	}

	result.GeneratePagingResponse(userResponseArray, input.Page, input.Limit, totalRecords)

	return _h.Helper.SendSuccess(c, result)
}

// @Tags Product
// @Description Find a Product
// @ID Find a Product
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ProductSuccessResponse
// @Failure 400 {object} response.ProductFailedResponse
// @Failure 404 {object} response.ProductFailedResponse
// @Router /product/{id} [get]
func (_h *ProductHandler) FindById(c echo.Context) error {

	var errResults []string
	id := c.Param("id")

	result, err := _h.ProductFindById(id)

	if result == nil || err != "" {

		errResults = append(errResults, err)
		return _h.Helper.SendBadRequest(c, errResults)
	}

	return _h.Helper.SendSuccess(c, result)

}

func (_h *ProductHandler) ProductFindById(id string) (*response.ProductResponse, string) {

	entityResult, err := _h.ProductRepository.FindById(id)
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

	result := &response.ProductResponse{}
	result.ParsingEntityToResponse(entityResult)

	return result, ""
}
