package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
)

type ProductInterface interface {
	FindByProductName(block string) (*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Create(e *entity.Product) (bool, error)
	Update(id string, e *entity.Product) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingProductRequest) ([]*entity.Product, error, int)
}
