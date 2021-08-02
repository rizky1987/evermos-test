package api

import (
	"evermos-test/config/env"
	"evermos-test/database/entity"
	"evermos-test/helper"
	"evermos-test/http/interfaces"
	"evermos-test/http/request"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"sync"
)

type PaymentHandler struct {
	Helper         					helper.HTTPHelper
	PaymentRepository 				interfaces.PaymentInterface
	CartRepository 					interfaces.CartInterface
	ProductRepository 				interfaces.ProductInterface
	InventoryAdjustmentRepository 	interfaces.InventoryAdjustmentInterface
	Config         					env.Config
}

// @Tags Payment
// @Description Callback Payment
// @ID Callback Payment
// @Accept  json
// @Produce  json
// @Param CreatePayment body request.CallbackPaymentRequest true "callback payment info"
// @Success 200 {object} response.PaymentSuccessResponse
// @Failure 400 {object} response.PaymentFailedResponse
// @Failure 404 {object} response.PaymentFailedResponse
// @Router /payment/callback [post]
func (_h *PaymentHandler) Callback(c echo.Context) error {

	var (
		errResults []string
		err        error
		input      request.CallbackPaymentRequest
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

	input.PaymentCode = helper.TrimWhiteSpace(input.PaymentCode)

	//get all cart records by payment Code
    cartEntities, _ := _h.CartRepository.FindByPaymentCode(input.PaymentCode)
	var wg sync.WaitGroup


	go func() {
		for i:=0; i<len(cartEntities); i++ {
			wg.Add(4)
			//Get product records
			productEntity, errProductEntity := _h.ProductRepository.FindById(cartEntities[i].ProductId)
			if errProductEntity != nil {
				errResults = append(errResults, "line 71" + errProductEntity.Error())
			}
			wg.Done()

			//Update Product Record
			if input.Status == helper.PaymentStatusSettlement {

				productEntity.OnHoldQuantity = productEntity.OnHoldQuantity - cartEntities[i].Quantity
				productEntity.SoldQuantity = productEntity.SoldQuantity + cartEntities[i].Quantity
			} else {

				productEntity.OnHoldQuantity = productEntity.OnHoldQuantity - cartEntities[i].Quantity
				productEntity.Quantity = productEntity.Quantity + cartEntities[i].Quantity
			}

			_h.ProductRepository.Update(productEntity.Id, productEntity)
			wg.Done()

			//Add inventory Adjustment History
			inventoryAdjustment := &entity.InventoryAdjustment{
				Quantity:  cartEntities[i].Quantity,
				ProductId: cartEntities[i].ProductId,
				Process:   helper.ProcessOutText,
				Note:      helper.ProcessSoldText,
			}

			inventoryAdjustment.ValidateBeforeCreate()
			_h.InventoryAdjustmentRepository.Create(inventoryAdjustment)
			wg.Done()

			_h.CartRepository.DeleteByPaymentCode(input.PaymentCode)
			wg.Done()
		}
	}()


	wg.Wait()
	err = _h.PaymentRepository.UpdateStatus(input.PaymentCode, helper.PaymentStatusSettlement)

	if err != nil {

		errResults = append(errResults, err.Error())
		return _h.Helper.SendBadRequest(c, errResults)

	}

	// End Save To DB

	return _h.Helper.SendSuccess(c, nil)
}