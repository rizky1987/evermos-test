package interfaces

import (
	"evermos-test/database/entity"
)

type PaymentRequestInterface interface {
	FindByPaymentRequestName(block string) (*entity.PaymentRequest, error)
	FindById(id string) (*entity.PaymentRequest, error)
	Create(e *entity.PaymentRequest) (bool, error)
	Update(id string, e *entity.PaymentRequest) (bool, error)
	//FindAll(searchParam request.SearchParamWithPagingPaymentRequestRequest) ([]*entity.PaymentRequest, error, int)
}
