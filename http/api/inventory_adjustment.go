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

type InventoryAdjustmentHandler struct {
	Helper         					helper.HTTPHelper
	Config         					env.Config
	InventoryAdjustmentRepository 	interfaces.InventoryAdjustmentInterface
	ProductRepository 				interfaces.ProductInterface
}

// @Tags InventoryAdjustment
// @Description Create InventoryAdjustment
// @ID Create InventoryAdjustment
// @Accept  json
// @Produce  json
// @Param CreateInventoryAdjustment body request.CreateInventoryAdjustmentRequest true "create inventory adjustment info"
// @Success 200 {object} response.InventoryAdjustmentSuccessResponse
// @Failure 400 {object} response.InventoryAdjustmentFailedResponse
// @Failure 404 {object} response.InventoryAdjustmentFailedResponse
// @Router /inventory-adjustment [post]
func (_h *InventoryAdjustmentHandler) CreateInventoryAdjustment(c echo.Context) error {

	var (
		errResults []string
		err        error
		input      request.CreateInventoryAdjustmentRequest
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
	var newMainEntityToSave entity.InventoryAdjustment

	errResults = newMainEntityToSave.ValidateBeforeCreate()
	newMainEntityToSave.MappingCreateDataToEntityStruct(input)
	_, err = _h.InventoryAdjustmentRepository.Create(&newMainEntityToSave)

	if err != nil {
		errResults = append(errResults, err.Error())
	}

	productId, errProductId := helper.ChangeStringOfObjectIdToBsonObjectId(input.ProductId)
	if errProductId != nil {
		errResults = append(errResults, errProductId.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}

	productEntity, errProductEntity := _h.ProductRepository.FindById(productId)
	if errProductEntity != nil {
		errResults = append(errResults, errProductEntity.Error())
	}

	if helper.TrimWhiteSpace(input.Process) == helper.ProcessInText {
		productEntity.Quantity = productEntity.Quantity + input.Quantity
	} else {
		productEntity.Quantity = productEntity.Quantity - input.Quantity
	}

	_h.ProductRepository.Update(productId, productEntity)
	if len(errResults) > 0 {
		return _h.Helper.SendBadRequest(c, errResults)
	}
	// End Save To DB

	return _h.Helper.SendSuccess(c, nil)

}

// @Tags InventoryAdjustment
// @Description Find All InventoryAdjustment
// @Accept  json
// @Produce  json
// @Param SearchInventoryAdjustment body request.SearchParamInventoryAdjustmentRequest true "search product info"
// @Success 200 {object} response.InventoryAdjustmentSuccessWithPagingResponse
// @Failure 400 {object} response.InventoryAdjustmentFailedResponse
// @Router /inventory-adjustment/find-all [post]
func (_h *InventoryAdjustmentHandler) FindAll(c echo.Context) error {

	var (
		errResults []string
		err   error
		input request.SearchParamWithPagingInventoryAdjustmentRequest
	)

	err = c.Bind(&input)
	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)
	}


	if err = _h.Helper.Validate.Struct(input); err != nil {
		return _h.Helper.SendInputValidationError(c, err.(validator.ValidationErrors))
	}

	entities, err, totalRecords := _h.InventoryAdjustmentRepository.FindAll(input)
	errResults = helper.NotFoundValidationForSearching(err)
	if len(errResults) > 0 {

		return _h.Helper.SendBadRequest(c, errResults)
	}

	// begin parsing from entity to response
	result := &response.InventoryAdjustmentSearchResponse{}
	userResponseArray := []*response.InventoryAdjustmentResponse{}
	for _, entity := range entities {
		userResponse := &response.InventoryAdjustmentResponse{}

		userResponse.ParsingEntityToResponse(entity)
		userResponseArray = append(userResponseArray, userResponse)
	}

	result.GeneratePagingResponse(userResponseArray, input.Page, input.Limit, totalRecords)

	return _h.Helper.SendSuccess(c, result)
}

// @Tags InventoryAdjustment
// @Description Find a InventoryAdjustment
// @ID Find a InventoryAdjustment
// @Accept  json
// @Produce  json
// @Success 200 {object} response.InventoryAdjustmentSuccessResponse
// @Failure 400 {object} response.InventoryAdjustmentFailedResponse
// @Failure 404 {object} response.InventoryAdjustmentFailedResponse
// @Router /inventory-adjustment/{id} [get]
func (_h *InventoryAdjustmentHandler) FindById(c echo.Context) error {

	var errResults []string
	id := c.Param("id")

	result, err := _h.InventoryAdjustmentFindById(id)

	if result == nil || err != "" {

		errResults = append(errResults, err)
		return _h.Helper.SendBadRequest(c, errResults)
	}

	return _h.Helper.SendSuccess(c, result)

}

func (_h *InventoryAdjustmentHandler) InventoryAdjustmentFindById(id string) (*response.InventoryAdjustmentResponse, string) {

	productId, errInventoryAdjustmentId := helper.ChangeStringOfObjectIdToBsonObjectId(id)
	if  errInventoryAdjustmentId != nil {
		return nil, errInventoryAdjustmentId.Error()
	}

	entityResult, err := _h.InventoryAdjustmentRepository.FindById(productId)
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

	result := &response.InventoryAdjustmentResponse{}
	result.ParsingEntityToResponse(entityResult)

	return result, ""
}
