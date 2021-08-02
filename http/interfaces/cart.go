package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
)

type CartInterface interface {
	FindById(id *bson.ObjectId) (*entity.Cart, error)
	Create(e *entity.Cart) (bool, error)
	Update(id *bson.ObjectId, e *entity.Cart) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingCartRequest) ([]*entity.Cart, error, int)
	Checkout(cartId []*bson.ObjectId, paymentCode string) error
}

