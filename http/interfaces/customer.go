package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
)

type CustomerInterface interface {
	FindByCustomerName(block string) (*entity.Customer, error)
	FindById(id string) (*entity.Customer, error)
	Create(e *entity.Customer) (bool, error)
	Update(id string, e *entity.Customer) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingCustomerRequest) ([]*entity.Customer, error, int)
}
