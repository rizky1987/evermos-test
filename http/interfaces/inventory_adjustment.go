package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
)

type InventoryAdjustmentInterface interface {
	FindByInventoryAdjustmentName(name string) (*entity.InventoryAdjustment, error)
	FindById(id *bson.ObjectId) (*entity.InventoryAdjustment, error)
	Create(e *entity.InventoryAdjustment) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingInventoryAdjustmentRequest) ([]*entity.InventoryAdjustment, error, int)
}
