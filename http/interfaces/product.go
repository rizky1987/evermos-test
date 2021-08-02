package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
)

type ProductInterface interface {
	FindByProductName(block string) (*entity.Product, error)
	FindById(id *bson.ObjectId) (*entity.Product, error)
	Create(e *entity.Product) (bool, error)
	Update(id *bson.ObjectId, e *entity.Product) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingProductRequest) ([]*entity.Product, error, int)
}
