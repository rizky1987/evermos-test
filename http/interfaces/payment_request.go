package interfaces

import (
	"evermos-test/database/entity"
	"gopkg.in/mgo.v2/bson"
)

type PaymentRequestInterface interface {
	FindByPaymentRequestName(block string) (*entity.PaymentRequest, error)
	FindById(id *bson.ObjectId) (*entity.PaymentRequest, error)
	Create(e *entity.PaymentRequest) (bool, error)
	Update(id *bson.ObjectId, e *entity.PaymentRequest) (bool, error)
	//FindAll(searchParam request.SearchParamWithPagingPaymentRequestRequest) ([]*entity.PaymentRequest, error, int)
}
