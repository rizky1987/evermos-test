package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
)

type CustomerInterface interface {
	FindByCustomerName(name string) (*entity.Customer, error)
	FindById(id *bson.ObjectId) (*entity.Customer, error)
	Create(e *entity.Customer) (bool, error)
	Update(id *bson.ObjectId, e *entity.Customer) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingCustomerRequest) ([]*entity.Customer, error, int)
}
