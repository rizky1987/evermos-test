package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
)

type CartInterface interface {
	FindByCartName(block string) (*entity.Cart, error)
	FindById(id string) (*entity.Cart, error)
	Create(e *entity.Cart) (bool, error)
	Update(id string, e *entity.Cart) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingCartRequest) ([]*entity.Cart, error, int)
	Checkout(cartId string) error
}

