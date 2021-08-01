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

type CartHandler struct {
	Helper         					helper.HTTPHelper
	CartRepository 					interfaces.CartInterface
	Config         					env.Config
}

// @Tags Cart
// @Description Create Cart
// @ID Create Cart
// @Accept  json
// @Produce  json
// @Param CreateCart body request.CreateCartRequest true "create cart info"
// @Success 200 {object} response.CartSuccessResponse
// @Failure 400 {object} response.CartFailedResponse
// @Failure 404 {object} response.CartFailedResponse
// @Router /cart [post]
func (_h *CartHandler) CreateCart(c echo.Context) error {

	var (
		errResults []string
		err        error
		input      request.CreateCartRequest
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
	var newMainEntityToSave entity.Cart
	errResults = newMainEntityToSave.ValidateBeforeCreate(input)

	if len(errResults) > 0 {
		return _h.Helper.SendBadRequest(c, errResults)
	}
	_, err = _h.CartRepository.Create(&newMainEntityToSave)

	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)

	}

	// End Save To DB

	return _h.Helper.SendSuccess(c, nil)

}

// @Tags Cart
// @Description Update Cart
// @ID Update Cart
// @Accept  json
// @Produce  json
// @Param CreateCart body request.UpdateCartRequest true "update cart info"
// @Success 200 {object} response.CartSuccessResponse
// @Failure 400 {object} response.CartFailedResponse
// @Failure 404 {object} response.CartFailedResponse
// @Router /cart/{id} [put]
func (_h *CartHandler) UpdateCart(c echo.Context) error {

	var (
		errResults []string
		err   error
		input request.UpdateCartRequest
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
	entityToUpdate, err := _h.CartRepository.FindById(id)
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

	_, err = _h.CartRepository.Update(id, entityToUpdate)

	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}

	// End Save To DB

	result, errStr := _h.CartFindById(id)

	if result == nil || errStr != "" {

		errResults = append(errResults, errStr)
		return _h.Helper.SendBadRequest(c, errResults)
	}

	return _h.Helper.SendSuccess(c, result)

}

// @Tags Cart
// @Description Find All Cart
// @Accept  json
// @Produce  json
// @Param Searchuser body request.SearchParamCartRequest true "search cart info"
// @Success 200 {object} response.CartSuccessWithPagingResponse
// @Failure 400 {object} response.CartFailedResponse
// @Router /cart/find-all [post]
func (_h *CartHandler) FindAll(c echo.Context) error {

	var (
		errResults []string
		err   error
		input request.SearchParamWithPagingCartRequest
	)

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}


	if err = _h.Helper.Validate.Struct(input); err != nil {
		return _h.Helper.SendInputValidationError(c, err.(validator.ValidationErrors))
	}

	entities, err, totalRecords := _h.CartRepository.FindAll(input)
	errResults = helper.NotFoundValidationForSearching(err)
	if len(errResults) > 0 {

		return _h.Helper.SendBadRequest(c, errResults)
	}

	// begin parsing from entity to response
	result := &response.CartSearchResponse{}
	userResponseArray := []*response.CartResponse{}
	for _, entity := range entities {
		userResponse := &response.CartResponse{}

		userResponse.ParsingEntityToResponse(entity)
		userResponseArray = append(userResponseArray, userResponse)
	}

	result.GeneratePagingResponse(userResponseArray, input.Page, input.Limit, totalRecords)

	return _h.Helper.SendSuccess(c, result)
}

// @Tags Cart
// @Description Find a Cart
// @ID Find a Cart
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CartSuccessResponse
// @Failure 400 {object} response.CartFailedResponse
// @Failure 404 {object} response.CartFailedResponse
// @Router /cart/{id} [get]
func (_h *CartHandler) FindById(c echo.Context) error {

	var errResults []string
	id := c.Param("id")

	result, err := _h.CartFindById(id)

	if result == nil || err != "" {

		errResults = append(errResults, err)
		return _h.Helper.SendBadRequest(c, errResults)
	}

	return _h.Helper.SendSuccess(c, result)

}

func (_h *CartHandler) CartFindById(id string) (*response.CartResponse, string) {

	entityResult, err := _h.CartRepository.FindById(id)
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

	result := &response.CartResponse{}
	result.ParsingEntityToResponse(entityResult)

	return result, ""
}
