package interfaces

import (
	"evermos-test/database/entity"
	"evermos-test/http/request"
)

type InventoryAdjustmentInterface interface {
	FindByInventoryAdjustmentName(block string) (*entity.InventoryAdjustment, error)
	FindById(id string) (*entity.InventoryAdjustment, error)
	Create(e *entity.InventoryAdjustment) (bool, error)
	FindAll(searchParam request.SearchParamWithPagingInventoryAdjustmentRequest) ([]*entity.InventoryAdjustment, error, int)
}
